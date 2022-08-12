package domain

import (
	"time"

	"github.com/google/uuid"
)

type Lesson struct {
	ID          uint       `json:"id"`
	UUID        uuid.UUID  `json:"uuid"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Objective   string     `json:"objective"`
	Type        string     `json:"type"`
	Module      string     `json:"module"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

type LessonActivity struct {
	ID         uint      `json:"id"`
	ActivityID uuid.UUID `json:"activity_id"`
	LessonID   uuid.UUID `json:"lesson_id"`
}
