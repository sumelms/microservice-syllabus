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
	"github.com/gorilla/mux"

	"github.com/sumelms/microservice-syllabus/internal/activity/domain"
	"github.com/sumelms/microservice-syllabus/pkg/validator"
)

type updateActivityRequest struct {
	UUID        uuid.UUID `json:"uuid" validate:"required"`
	Name        string    `json:"name" validate:"required,max=100"`
	Description string    `json:"description" validate:"required,max=255"`
	ContentID   uuid.UUID `json:"content_id" validate:"required"`
	ContentType string    `json:"content_type" validate:"required,max=140"`
}

type updateActivityResponse struct {
	UUID        uuid.UUID `json:"uuid"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ContentID   uuid.UUID `json:"content_id"`
	ContentType string    `json:"content_type"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewUpdateActivityHandler(s domain.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeUpdateActivityEndpoint(s),
		decodeUpdateActivityRequest,
		encodeUpdateActivityResponse,
		opts...,
	)
}

func makeUpdateActivityEndpoint(s domain.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(updateActivityRequest)
		if !ok {
			return nil, fmt.Errorf("invalid argument")
		}

		v := validator.NewValidator()
		if err := v.Validate(req); err != nil {
			return nil, err
		}

		a := domain.Activity{}
		data, err := json.Marshal(req)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(data, &a)
		if err != nil {
			return nil, err
		}

		if err := s.UpdateActivity(ctx, &a); err != nil {
			return nil, err
		}

		return updateActivityResponse{
			UUID:        a.UUID,
			Name:        a.Name,
			Description: a.Description,
			ContentID:   a.ContentID,
			ContentType: a.ContentType,
		}, nil
	}
}

func decodeUpdateActivityRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["uuid"]
	if !ok {
		return nil, fmt.Errorf("invalid argument")
	}

	var req updateActivityRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	req.UUID = uuid.MustParse(id)

	return req, nil
}

func encodeUpdateActivityResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
