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

	"github.com/sumelms/microservice-syllabus/internal/lesson/domain"

	"github.com/sumelms/microservice-syllabus/pkg/validator"
)

type createLessonRequest struct {
	Name        string `json:"name" validate:"required,max=100"`
	Description string `json:"description" validate:"max=255"`
	Objective   string `json:"objective" validate:"max=100"`
	Type        string `json:"type" validate:"required,max=40"`
	Module      string `json:"module" validate:"max=40"`
}

type createLessonResponse struct {
	UUID        uuid.UUID `json:"uuid"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	Objective   string    `json:"objective,omitempty"`
	Type        string    `json:"type"`
	Module      string    `json:"module,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// NewCreateLessonHandler creates new lesson handler
// @Summary      Create course
// @Description  Create a new course
// @Tags         course
// @Accept       json
// @Produce      json
// @Param        course	  body		createLessonRequest		true	"Add Course"
// @Success      200      {object}  createLessonResponse
// @Failure      400      {object}  error
// @Failure      404      {object}  error
// @Failure      500      {object}  error
// @Router       /courses [post]
func NewCreateLessonHandler(s domain.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeCreateLessonEndpoint(s),
		decodeCreateLessonRequest,
		encodeCreateLessonResponse,
		opts...,
	)
}

func makeCreateLessonEndpoint(s domain.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(createLessonRequest)
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

		if err := s.CreateLesson(ctx, &c); err != nil {
			return nil, err
		}

		return createLessonResponse{
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

func decodeCreateLessonRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req createLessonRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func encodeCreateLessonResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
