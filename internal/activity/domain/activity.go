package domain

import "time"

type Activity struct {
	ID          uint       `json:"id"`
	UUID        string     `json:"uuid"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	ContentID   string     `json:"content_id"`
	ContentType string     `json:"content_type"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}
