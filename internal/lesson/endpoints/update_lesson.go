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

	"github.com/sumelms/microservice-syllabus/internal/lesson/domain"

	"github.com/sumelms/microservice-syllabus/pkg/validator"
)

type updateLessonRequest struct {
	UUID        uuid.UUID `json:"uuid" validate:"required"`
	Name        string    `json:"name" validate:"required,max=100"`
	Description string    `json:"description" validate:"max=255"`
	Objective   string    `json:"objective" validate:"max=100"`
	Type        string    `json:"type" validate:"required,max=40"`
	Module      string    `json:"module" validate:"max=40"`
}

type updateLessonResponse struct {
	UUID        uuid.UUID `json:"uuid"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	Objective   string    `json:"objective,omitempty"`
	Type        string    `json:"type"`
	Module      string    `json:"module,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewUpdateLessonHandler(s domain.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeUpdateLessonEndpoint(s),
		decodeUpdateLessonRequest,
		encodeUpdateLessonResponse,
		opts...,
	)
}

func makeUpdateLessonEndpoint(s domain.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(updateLessonRequest)
		if !ok {
			return nil, fmt.Errorf("invalid argument")
		}

		v := validator.NewValidator()
		if err := v.Validate(req); err != nil {
			return nil, err
		}

		c := domain.Lesson{}
		data, _ := json.Marshal(req)
		if err := json.Unmarshal(data, &c); err != nil {
			return nil, err
		}

		if err := s.UpdateLesson(ctx, &c); err != nil {
			return nil, err
		}

		return updateLessonResponse{
			UUID:        c.UUID,
			Name:        c.Name,
			Description: c.Description,
			Objective:   c.Objective,
			Type:        c.Type,
			Module:      c.Module,
		}, nil
	}
}

func decodeUpdateLessonRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["uuid"]
	if !ok {
		return nil, fmt.Errorf("invalid argument")
	}

	var req updateLessonRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	req.UUID = uuid.MustParse(id)

	return req, nil
}

func encodeUpdateLessonResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
