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
	Underline   string    `json:"underline" validate:"required,max=100"`
	Image       string    `json:"image"`
	ImageCover  string    `json:"image_cover"`
	Excerpt     string    `json:"excerpt" validate:"required,max=140"`
	Description string    `json:"description" validate:"required,max=255"`
}

type updateLessonResponse struct {
	UUID        uuid.UUID `json:"uuid"`
	Name        string    `json:"name"`
	Underline   string    `json:"underline"`
	Image       string    `json:"image"`
	ImageCover  string    `json:"image_cover"`
	Excerpt     string    `json:"excerpt"`
	Description string    `json:"description"`
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
			Underline:   c.Underline,
			Image:       c.Image,
			ImageCover:  c.ImageCover,
			Excerpt:     c.Excerpt,
			Description: c.Description,
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
