package transport

import (
	"net/http"

	"github.com/sumelms/microservice-activity/internal/activitylog/endpoints"
	"github.com/sumelms/microservice-activity/pkg/errors"

	"github.com/go-kit/kit/log"
	kittransport "github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/sumelms/microservice-activity/internal/activitylog/domain"
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

	r.Handle("/activitylogs", createActivityHandler).Methods(http.MethodPost)
	r.Handle("/activitylogs", listActivityHandler).Methods(http.MethodGet)
	r.Handle("/activitylogs/{uuid}", findActivityHandler).Methods(http.MethodGet)
	r.Handle("/activitylogs/{uuid}", updateCourseHandler).Methods(http.MethodPut)
	r.Handle("/activitylogs/{uuid}", deleteCourseHandler).Methods(http.MethodDelete)
}
