package transport

import (
	"net/http"

	"github.com/sumelms/microservice-syllabus/internal/lesson/endpoints"
	"github.com/sumelms/microservice-syllabus/pkg/errors"

	kittransport "github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/gorilla/mux"

	"github.com/sumelms/microservice-syllabus/internal/lesson/domain"
)

func NewHTTPHandler(r *mux.Router, s domain.ServiceInterface, logger log.Logger) {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(kittransport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(errors.EncodeError),
	}

	// Lesson handlers

	listLessonHandler := endpoints.NewListLessonHandler(s, opts...)
	createLessonHandler := endpoints.NewCreateLessonHandler(s, opts...)
	findLessonHandler := endpoints.NewFindLessonHandler(s, opts...)
	updateLessonHandler := endpoints.NewUpdateLessonHandler(s, opts...)
	deleteLessonHandler := endpoints.NewDeleteLessonHandler(s, opts...)

	r.Handle("/lessons", createLessonHandler).Methods(http.MethodPost)
	r.Handle("/lessons", listLessonHandler).Methods(http.MethodGet)
	r.Handle("/lessons/{uuid}", findLessonHandler).Methods(http.MethodGet)
	r.Handle("/lessons/{uuid}", updateLessonHandler).Methods(http.MethodPut)
	r.Handle("/lessons/{uuid}", deleteLessonHandler).Methods(http.MethodDelete)
}
