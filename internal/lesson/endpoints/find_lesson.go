package endpoints

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"github.com/sumelms/microservice-syllabus/internal/lesson/domain"
)

type findLessonRequest struct {
	UUID uuid.UUID `json:"uuid"`
}

type findLessonResponse struct {
	UUID        uuid.UUID `json:"uuid"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	Objective   string    `json:"objective,omitempty"`
	Type        string    `json:"type"`
	Module      string    `json:"module,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewFindLessonHandler(s domain.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeFindLessonEndpoint(s),
		decodeFindLessonRequest,
		encodeFindLessonResponse,
		opts...,
	)
}

func makeFindLessonEndpoint(s domain.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(findLessonRequest)
		if !ok {
			return nil, fmt.Errorf("invalid argument")
		}

		c, err := s.Lesson(ctx, req.UUID)
		if err != nil {
			return nil, err
		}

		return &findLessonResponse{
			UUID:        c.UUID,
			Name:        c.Name,
			Description: c.Description,
			Objective:   c.Objective,
			Type:        c.Type,
			Module:      c.Module,
			CreatedAt:   c.CreatedAt,
			UpdatedAt:   c.UpdatedAt,
		}, nil
	}
}

func decodeFindLessonRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["uuid"]
	if !ok {
		return nil, fmt.Errorf("invalid argument")
	}

	uid := uuid.MustParse(id)

	return findLessonRequest{UUID: uid}, nil
}

func encodeFindLessonResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
