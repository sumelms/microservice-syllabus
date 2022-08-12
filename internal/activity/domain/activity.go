package domain

import (
	"time"

	"github.com/google/uuid"
)

type Activity struct {
	ID          uint       `json:"id"`
	UUID        uuid.UUID  `json:"uuid"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	ContentID   uuid.UUID  `json:"content_id"`
	ContentType string     `json:"content_type"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}
