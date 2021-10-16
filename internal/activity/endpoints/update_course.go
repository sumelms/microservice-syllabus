package endpoints

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/sumelms/microservice-activity/internal/activity/domain"
	"github.com/sumelms/microservice-activity/pkg/validator"
)

type updateActivityRequest struct {
	UUID        string `json:"uuid" validate:"required"`
	Title       string `json:"title" validate:"required,max=100"`
	Subtitle    string `json:"subtitle" validate:"required,max=100"`
	Excerpt     string `json:"excerpt" validate:"required,max=140"`
	Description string `json:"description" validate:"required,max=255"`
}

type updateActivityResponse struct {
	UUID        string    `json:"uuid"`
	Title       string    `json:"title"`
	Subtitle    string    `json:"subtitle"`
	Excerpt     string    `json:"excerpt"`
	Description string    `json:"description"`
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
		data, _ := json.Marshal(req)
		err := json.Unmarshal(data, &a)
		if err != nil {
			return nil, err
		}

		updated, err := s.UpdateActivity(ctx, &a)
		if err != nil {
			return nil, err
		}

		return updateActivityResponse{
			UUID:        updated.UUID,
			Title:       updated.Title,
			Subtitle:    updated.Subtitle,
			Excerpt:     updated.Excerpt,
			Description: updated.Description,
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

	req.UUID = id

	return req, nil
}

func encodeUpdateActivityResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
