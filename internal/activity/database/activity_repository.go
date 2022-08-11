package database

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/sumelms/microservice-syllabus/internal/activity/domain"
	"github.com/sumelms/microservice-syllabus/pkg/errors"
)

func NewActivityRepository(db *sqlx.DB) (activityRepository, error) { //nolint: revive
	sqlStatements := make(map[string]*sqlx.Stmt)

	for queryName, query := range queriesActivity() {
		stmt, err := db.Preparex(query)
		if err != nil {
			return activityRepository{}, errors.WrapErrorf(err, errors.ErrCodeUnknown,
				"error preparing statement %s", queryName)
		}
		sqlStatements[queryName] = stmt
	}

	return activityRepository{
		statements: sqlStatements,
	}, nil
}

type activityRepository struct {
	statements map[string]*sqlx.Stmt
}

// Activity get the Activity by given id
func (r activityRepository) Activity(id uuid.UUID) (domain.Activity, error) {
	stmt, ok := r.statements[getActivity]
	if !ok {
		return domain.Activity{}, errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", getActivity)
	}

	var c domain.Activity
	if err := stmt.Get(&c, id); err != nil {
		return domain.Activity{}, errors.WrapErrorf(err, errors.ErrCodeUnknown, "error getting activity")
	}
	return c, nil
}

// Activities list all activities
func (r activityRepository) Activities() ([]domain.Activity, error) {
	stmt, ok := r.statements[listActivity]
	if !ok {
		return []domain.Activity{}, errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", listActivity)
	}

	var cc []domain.Activity
	if err := stmt.Select(&cc); err != nil {
		return []domain.Activity{}, errors.WrapErrorf(err, errors.ErrCodeUnknown, "error getting activities")
	}
	return cc, nil
}

// CreateActivity creates a new activity
func (r activityRepository) CreateActivity(c *domain.Activity) error {
	stmt, ok := r.statements[createActivity]
	if !ok {
		return errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", createActivity)
	}

	if err := stmt.Get(c, c.Code, c.Name, c.Underline, c.Image, c.ImageCover, c.Excerpt, c.Description); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error creating activity")
	}
	return nil
}

// UpdateActivity update the given activity
func (r activityRepository) UpdateActivity(c *domain.Activity) error {
	stmt, ok := r.statements[updateActivity]
	if !ok {
		return errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", updateActivity)
	}

	if err := stmt.Get(c, c.Code, c.Name, c.Underline, c.Image, c.ImageCover, c.Excerpt, c.Description, c.UUID); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error updating activity")
	}
	return nil
}

// DeleteActivity soft delete the activity by given id
func (r activityRepository) DeleteActivity(id uuid.UUID) error {
	stmt, ok := r.statements[deleteActivity]
	if !ok {
		return errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", deleteActivity)
	}

	if _, err := stmt.Exec(id); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error deleting activity")
	}
	return nil
}
