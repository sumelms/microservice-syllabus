package activity

import (
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	"github.com/go-kit/log"

	"github.com/sumelms/microservice-syllabus/internal/activity/database"
	"github.com/sumelms/microservice-syllabus/internal/activity/domain"
	"github.com/sumelms/microservice-syllabus/internal/activity/transport"
)

func NewService(db *sqlx.DB, logger log.Logger) (*domain.Service, error) {
	activity, err := database.NewActivityRepository(db)
	if err != nil {
		return nil, err
	}

	service, err := domain.NewService(
		domain.WithLogger(logger),
		domain.WithActivityRepository(activity))
	if err != nil {
		return nil, err
	}
	return service, nil
}

func NewHTTPService(router *mux.Router, service domain.ServiceInterface, logger log.Logger) error {
	transport.NewHTTPHandler(router, service, logger)
	return nil
}
