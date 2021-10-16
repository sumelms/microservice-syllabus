package database

import (
	"errors"

	"github.com/go-kit/kit/log"
	"github.com/jinzhu/gorm"
	"github.com/sumelms/microservice-activity/internal/activity/domain"
	merrors "github.com/sumelms/microservice-activity/pkg/errors"
)

const (
	whereActivityUUID = "uuid = ?"
)

// Repository struct
type Repository struct {
	db     *gorm.DB
	logger log.Logger
}

// NewRepository creates a new profile repository
func NewRepository(db *gorm.DB, logger log.Logger) *Repository {
	db.AutoMigrate(&Activity{})

	return &Repository{
		db:     db,
		logger: logger,
	}
}

// List activities
func (r *Repository) List() ([]domain.Activity, error) {
	var activities []Activity

	query := r.db.Find(&activities)
	if query.RecordNotFound() {
		return []domain.Activity{}, nil
	}
	if err := query.Error; err != nil {
		return []domain.Activity{}, merrors.WrapErrorf(err, merrors.ErrCodeUnknown, "list activities")
	}

	var list []domain.Activity
	for i := range activities {
		c := activities[i]
		list = append(list, toDomainModel(&c))
	}
	return list, nil
}

// Create creates a activity
func (r *Repository) Create(activity *domain.Activity) (domain.Activity, error) {
	entity := toDBModel(activity)

	if err := r.db.Create(&entity).Error; err != nil {
		return domain.Activity{}, merrors.WrapErrorf(err, merrors.ErrCodeUnknown, "can't create activity")
	}
	return toDomainModel(&entity), nil
}

// Find get a activity by its ID
func (r *Repository) Find(id string) (domain.Activity, error) {
	var activity Activity

	query := r.db.Where(whereActivityUUID, id).First(&activity)
	if query.RecordNotFound() {
		return domain.Activity{}, merrors.NewErrorf(merrors.ErrCodeNotFound, "activity not found")
	}
	if err := query.Error; err != nil {
		return domain.Activity{}, merrors.WrapErrorf(err, merrors.ErrCodeUnknown, "find activity")
	}

	return toDomainModel(&activity), nil
}

// Update the given activity
func (r *Repository) Update(c *domain.Activity) (domain.Activity, error) {
	var activity Activity

	query := r.db.Where(whereActivityUUID, c.UUID).First(&activity)

	if query.RecordNotFound() {
		return domain.Activity{}, merrors.NewErrorf(merrors.ErrCodeNotFound, "activity not found")
	}

	query = r.db.Model(&activity).Update(&c)

	if err := query.Error; err != nil {
		return domain.Activity{}, merrors.WrapErrorf(err, merrors.ErrCodeUnknown, "can't update activity")
	}

	return *c, nil
}

// Delete a activity by its ID
func (r *Repository) Delete(id string) error {
	query := r.db.Where(whereActivityUUID, id).Delete(&Activity{})

	if err := query.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return merrors.WrapErrorf(err, merrors.ErrCodeNotFound, "activity not found")
		}
		return merrors.WrapErrorf(err, merrors.ErrCodeUnknown, "delete activity")
	}

	return nil
}
