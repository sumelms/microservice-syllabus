package endpoints

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/sumelms/microservice-activity/internal/activity/domain"
)

type listActivityResponse struct {
	Activities []findActivityResponse `json:"activities"`
}

func NewListActivityHandler(s domain.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeListActivityEndpoint(s),
		decodeListActivityRequest,
		encodeListActivityResponse,
		opts...,
	)
}

func makeListActivityEndpoint(s domain.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		activities, err := s.ListActivity(ctx)
		if err != nil {
			return nil, err
		}

		var list []findActivityResponse
		for i := range activities {
			a := activities[i]
			list = append(list, findActivityResponse{
				UUID:        a.UUID,
				Title:       a.Title,
				Subtitle:    a.Subtitle,
				Excerpt:     a.Excerpt,
				Description: a.Description,
				CreatedAt:   a.CreatedAt,
				UpdatedAt:   a.UpdatedAt,
			})
		}

		return &listActivityResponse{Activities: list}, nil
	}
}

func decodeListActivityRequest(_ context.Context, _ *http.Request) (interface{}, error) {
	return nil, nil
}

func encodeListActivityResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
