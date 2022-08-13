package endpoints

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/google/uuid"

	"github.com/sumelms/microservice-syllabus/internal/activity/domain"
	"github.com/sumelms/microservice-syllabus/pkg/validator"
)

type createActivityRequest struct {
	Name        string    `json:"name" validate:"required,max=100"`
	Description string    `json:"description" validate:"required,max=255"`
	ContentID   uuid.UUID `json:"content_id" validate:"required"`
	ContentType string    `json:"content_type" validate:"required,max=140"`
}

type createActivityResponse struct {
	UUID        uuid.UUID `json:"uuid"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ContentID   uuid.UUID `json:"content_id"`
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

		if err := s.CreateActivity(ctx, &a); err != nil {
			return nil, err
		}

		return createActivityResponse{
			UUID:        a.UUID,
			Name:        a.Name,
			Description: a.Description,
			ContentID:   a.ContentID,
			ContentType: a.ContentType,
			CreatedAt:   a.CreatedAt,
			UpdatedAt:   a.UpdatedAt,
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
