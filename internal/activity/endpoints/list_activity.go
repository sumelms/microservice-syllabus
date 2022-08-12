package endpoints

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"

	"github.com/sumelms/microservice-syllabus/internal/activity/domain"
)

type listActivityRequest struct {
	ContentID   string `json:"content_id"`
	ContentType string `json:"content_type"`
}

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
		req, ok := request.(listActivityRequest)
		if !ok {
			return nil, fmt.Errorf("invalid argument")
		}

		filters := make(map[string]interface{})
		if len(req.ContentID) > 0 {
			filters["content_id"] = req.ContentID
		}
		if len(req.ContentType) > 0 {
			filters["content_type"] = req.ContentType
		}

		activities, err := s.Activities(ctx)
		if err != nil {
			return nil, err
		}

		var list []findActivityResponse
		for i := range activities {
			a := activities[i]
			list = append(list, findActivityResponse{
				UUID:        a.UUID,
				Name:        a.Name,
				Description: a.Description,
				ContentID:   a.ContentID,
				ContentType: a.ContentType,
				CreatedAt:   a.CreatedAt,
				UpdatedAt:   a.UpdatedAt,
			})
		}

		return &listActivityResponse{Activities: list}, nil
	}
}

func decodeListActivityRequest(_ context.Context, r *http.Request) (interface{}, error) {
	contentID := r.FormValue("content_id")
	contentType := r.FormValue("content_type")
	return listActivityRequest{
		ContentID:   contentID,
		ContentType: contentType,
	}, nil
}

func encodeListActivityResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
