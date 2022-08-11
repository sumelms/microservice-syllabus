package endpoints

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/sumelms/microservice-syllabus/internal/activity/domain"
)

type deleteActivityRequest struct {
	UUID string `json:"uuid" validate:"required"`
}

func NewDeleteActivityHandler(s domain.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeDeleteActivityEndpoint(s),
		decodeDeleteActivityRequest,
		encodeDeleteActivityResponse,
		opts...,
	)
}

func makeDeleteActivityEndpoint(s domain.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(deleteActivityRequest)
		if !ok {
			return nil, fmt.Errorf("invalid argument")
		}

		err := s.DeleteActivity(ctx, req.UUID)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}
}

func decodeDeleteActivityRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["uuid"]
	if !ok {
		return nil, fmt.Errorf("invalid argument")
	}

	return deleteActivityRequest{UUID: id}, nil
}

func encodeDeleteActivityResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
