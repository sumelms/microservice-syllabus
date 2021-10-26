package database

import (
	"errors"

	"github.com/go-kit/kit/log"
	"github.com/jinzhu/gorm"
	"github.com/sumelms/microservice-activity/internal/activitylog/domain"
	merrors "github.com/sumelms/microservice-activity/pkg/errors"
)

const (
	whereActivityUUID = "UUID = ?"
)

// Repository struct
type Repository struct {
	db     *gorm.DB
	logger log.Logger
}

// NewRepository creates a new activity log repository
func NewRepository(db *gorm.DB, logger log.Logger) *Repository {
	db.AutoMigrate(&ActivityLog{})

	return &Repository{
		db:     db,
		logger: logger,
	}
}

// List activity logs
func (r *Repository) List(filters map[string]interface{}) ([]domain.ActivityLog, error) {
	var activityLogs []ActivityLog

	query := r.db.Find(&activityLogs, filters)
	if query.RecordNotFound() {
		return []domain.ActivityLog{}, nil
	}
	if err := query.Error; err != nil {
		return []domain.ActivityLog{}, merrors.WrapErrorf(err, merrors.ErrCodeUnknown, "list activityLogs")
	}

	var list []domain.ActivityLog
	for i := range activityLogs {
		al := activityLogs[i]
		list = append(list, toDomainModel(&al))
	}
	return list, nil
}

// Create creates an activity log
func (r *Repository) Create(activity *domain.ActivityLog) (domain.ActivityLog, error) {
	entity := toDBModel(activity)

	if err := r.db.Create(&entity).Error; err != nil {
		return domain.ActivityLog{}, merrors.WrapErrorf(err, merrors.ErrCodeUnknown, "can't create activity log")
	}
	return toDomainModel(&entity), nil
}

// Find get an activity log by its ID
func (r *Repository) Find(id string) (domain.ActivityLog, error) {
	var activity ActivityLog

	query := r.db.Where(whereActivityUUID, id).First(&activity)
	if query.RecordNotFound() {
		return domain.ActivityLog{}, merrors.NewErrorf(merrors.ErrCodeNotFound, "activity log not found")
	}
	if err := query.Error; err != nil {
		return domain.ActivityLog{}, merrors.WrapErrorf(err, merrors.ErrCodeUnknown, "find activity log")
	}

	return toDomainModel(&activity), nil
}

// Update the given activity log
func (r *Repository) Update(al *domain.ActivityLog) (domain.ActivityLog, error) {
	var activity ActivityLog

	query := r.db.Where(whereActivityUUID, al.UUID).First(&activity)

	if query.RecordNotFound() {
		return domain.ActivityLog{}, merrors.NewErrorf(merrors.ErrCodeNotFound, "activity log not found")
	}

	query = r.db.Model(&activity).Update(&al)

	if err := query.Error; err != nil {
		return domain.ActivityLog{}, merrors.WrapErrorf(err, merrors.ErrCodeUnknown, "can't update activity log")
	}

	return *al, nil
}

// Delete an activity log by its ID
func (r *Repository) Delete(id string) error {
	query := r.db.Where(whereActivityUUID, id).Delete(&ActivityLog{})

	if err := query.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return merrors.WrapErrorf(err, merrors.ErrCodeNotFound, "activity log not found")
		}
		return merrors.WrapErrorf(err, merrors.ErrCodeUnknown, "delete activity log")
	}

	return nil
}
