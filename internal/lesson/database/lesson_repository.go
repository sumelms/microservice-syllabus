package database

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/sumelms/microservice-syllabus/internal/lesson/domain"
	"github.com/sumelms/microservice-syllabus/pkg/errors"
)

func NewLessonRepository(db *sqlx.DB) (lessonRepository, error) { //nolint: revive
	sqlStatements := make(map[string]*sqlx.Stmt)

	for queryName, query := range queriesLesson() {
		stmt, err := db.Preparex(query)
		if err != nil {
			return lessonRepository{}, errors.WrapErrorf(err, errors.ErrCodeUnknown,
				"error preparing statement %s", queryName)
		}
		sqlStatements[queryName] = stmt
	}

	return lessonRepository{
		statements: sqlStatements,
	}, nil
}

type lessonRepository struct {
	statements map[string]*sqlx.Stmt
}

// Lesson get the Lesson by given id
func (r lessonRepository) Lesson(id uuid.UUID) (domain.Lesson, error) {
	stmt, ok := r.statements[getLesson]
	if !ok {
		return domain.Lesson{}, errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", getLesson)
	}

	var c domain.Lesson
	if err := stmt.Get(&c, id); err != nil {
		return domain.Lesson{}, errors.WrapErrorf(err, errors.ErrCodeUnknown, "error getting lesson")
	}
	return c, nil
}

// Lessons list all lessons
func (r lessonRepository) Lessons() ([]domain.Lesson, error) {
	stmt, ok := r.statements[listLesson]
	if !ok {
		return []domain.Lesson{}, errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", listLesson)
	}

	var cc []domain.Lesson
	if err := stmt.Select(&cc); err != nil {
		return []domain.Lesson{}, errors.WrapErrorf(err, errors.ErrCodeUnknown, "error getting lesson")
	}
	return cc, nil
}

// CreateLesson creates a new lesson
func (r lessonRepository) CreateLesson(c *domain.Lesson) error {
	stmt, ok := r.statements[createLesson]
	if !ok {
		return errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", createLesson)
	}

	if err := stmt.Get(c, c.Code, c.Name, c.Underline, c.Image, c.ImageCover, c.Excerpt, c.Description); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error creating lesson")
	}
	return nil
}

// UpdateLesson update the given lesson
func (r lessonRepository) UpdateLesson(c *domain.Lesson) error {
	stmt, ok := r.statements[updateLesson]
	if !ok {
		return errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", updateLesson)
	}

	if err := stmt.Get(c, c.Code, c.Name, c.Underline, c.Image, c.ImageCover, c.Excerpt, c.Description, c.UUID); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error updating lesson")
	}
	return nil
}

// DeleteLesson soft delete the lesson by given id
func (r lessonRepository) DeleteLesson(id uuid.UUID) error {
	stmt, ok := r.statements[deleteLesson]
	if !ok {
		return errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", deleteLesson)
	}

	if _, err := stmt.Exec(id); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error deleting lesson")
	}
	return nil
}
