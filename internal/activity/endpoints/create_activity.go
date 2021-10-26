package endpoints

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/sumelms/microservice-activity/internal/activity/domain"
	"github.com/sumelms/microservice-activity/pkg/validator"
)

type createActivityRequest struct {
	Title       string `json:"title" validate:"required,max=100"`
	Description string `json:"description" validate:"required,max=255"`
	ContentID   string `json:"content_id" validate:"required"`
	ContentType string `json:"content_type" validate:"required,max=140"`
}

type createActivityResponse struct {
	UUID        string    `json:"uuid"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ContentID   string    `json:"content_id"`
	ContentType string    `json:"content_type"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewCreateActivityHandler(s domain.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeCreateActivityEndpoint(s),
		decodeCreateActivityRequest,
		encodeCreateActivityResponse,
		opts...,
	)
}

func makeCreateActivityEndpoint(s domain.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(createActivityRequest)
		if !ok {
			return nil, fmt.Errorf("invalid argument")
		}

		v := validator.NewValidator()
		if err := v.Validate(req); err != nil {
			return nil, err
		}

		a := domain.Activity{}
		data, _ := json.Marshal(req)
		err := json.Unmarshal(data, &a)
		if err != nil {
			return nil, err
		}

		created, err := s.CreateActivity(ctx, &a)
		if err != nil {
			return nil, err
		}

		return createActivityResponse{
			UUID:        created.UUID,
			Title:       created.Title,
			Description: created.Description,
			ContentID:   created.ContentID,
			ContentType: created.ContentType,
			CreatedAt:   created.CreatedAt,
			UpdatedAt:   created.UpdatedAt,
		}, err
	}
}

func decodeCreateActivityRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req createActivityRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func encodeCreateActivityResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}