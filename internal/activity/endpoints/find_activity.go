package endpoints

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/sumelms/microservice-activity/internal/activity/domain"
)

type findActivityRequest struct {
	UUID string `json:"uuid"`
}

type findActivityResponse struct {
	UUID        string    `json:"uuid"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ContentID   string    `json:"content_id"`
	ContentType string    `json:"content_type"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewFindActivityHandler(s domain.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeFindActivityEndpoint(s),
		decodeFindActivityRequest,
		encodeFindActivityResponse,
		opts...,
	)
}

func makeFindActivityEndpoint(s domain.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(findActivityRequest)
		if !ok {
			return nil, fmt.Errorf("invalid argument")
		}

		a, err := s.FindActivity(ctx, req.UUID)
		if err != nil {
			return nil, err
		}

		return &findActivityResponse{
			UUID:        a.UUID,
			Title:       a.Title,
			Description: a.Description,
			ContentID:   a.ContentID,
			ContentType: a.ContentType,
		}, nil
	}
}

func decodeFindActivityRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["uuid"]
	if !ok {
		return nil, fmt.Errorf("invalid argument")
	}

	return findActivityRequest{UUID: id}, nil
}

func encodeFindActivityResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
