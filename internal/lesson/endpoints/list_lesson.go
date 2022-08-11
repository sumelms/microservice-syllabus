package endpoints

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"

	"github.com/sumelms/microservice-syllabus/internal/lesson/domain"
)

type listLessonResponse struct {
	Lessons []findLessonResponse `json:"lessons"`
}

func NewListLessonHandler(s domain.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeListLessonEndpoint(s),
		decodeListLessonRequest,
		encodeListLessonResponse,
		opts...,
	)
}

func makeListLessonEndpoint(s domain.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		cc, err := s.Lessons(ctx)
		if err != nil {
			return nil, err
		}

		var list []findLessonResponse
		for i := range cc {
			c := cc[i]
			list = append(list, findLessonResponse{
				UUID:        c.UUID,
				Name:        c.Name,
				Underline:   c.Underline,
				Image:       c.Image,
				ImageCover:  c.ImageCover,
				Excerpt:     c.Excerpt,
				Description: c.Description,
				CreatedAt:   c.CreatedAt,
				UpdatedAt:   c.UpdatedAt,
			})
		}

		return &listLessonResponse{Lessons: list}, nil
	}
}

func decodeListLessonRequest(_ context.Context, _ *http.Request) (interface{}, error) {
	return nil, nil
}

func encodeListLessonResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
