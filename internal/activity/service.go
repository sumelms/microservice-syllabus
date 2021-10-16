package activity

import (
	"github.com/gorilla/mux"

	"github.com/go-kit/kit/log"

	"github.com/jinzhu/gorm"
	"github.com/sumelms/microservice-activity/internal/activity/database"
	"github.com/sumelms/microservice-activity/internal/activity/domain"
	"github.com/sumelms/microservice-activity/internal/activity/transport"
)

func NewHTTPService(router *mux.Router, db *gorm.DB, logger log.Logger) {
	repository := database.NewRepository(db, logger)
	service := domain.NewService(repository, logger)

	transport.NewHTTPHandler(router, service, logger)
}