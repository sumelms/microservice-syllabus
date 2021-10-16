package transport

import (
	"net/http"

	"github.com/sumelms/microservice-activity/internal/activity/endpoints"
	"github.com/sumelms/microservice-activity/pkg/errors"

	"github.com/go-kit/kit/log"
	kittransport "github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/sumelms/microservice-activity/internal/activity/domain"
)

func NewHTTPHandler(r *mux.Router, s domain.ServiceInterface, logger log.Logger) {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(kittransport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(errors.EncodeError),
	}

	listActivityHandler := endpoints.NewListActivityHandler(s, opts...)
	createActivityHandler := endpoints.NewCreateActivityHandler(s, opts...)
	findActivityHandler := endpoints.NewFindActivityHandler(s, opts...)
	updateCourseHandler := endpoints.NewUpdateActivityHandler(s, opts...)
	deleteCourseHandler := endpoints.NewDeleteActivityHandler(s, opts...)

	r.Handle("/activities", createActivityHandler).Methods(http.MethodPost)
	r.Handle("/activities", listActivityHandler).Methods(http.MethodGet)
	r.Handle("/activities/{uuid}", findActivityHandler).Methods(http.MethodGet)
	r.Handle("/activities/{uuid}", updateCourseHandler).Methods(http.MethodPut)
	r.Handle("/activities/{uuid}", deleteCourseHandler).Methods(http.MethodDelete)
}
