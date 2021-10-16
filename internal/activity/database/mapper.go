package database

import (
	"github.com/google/uuid"
	"github.com/sumelms/microservice-activity/internal/activity/domain"
)

func toDBModel(entity *domain.Activity) Activity {
	activity := Activity{
		Title:       entity.Title,
		Subtitle:    entity.Subtitle,
		Excerpt:     entity.Excerpt,
		Description: entity.Description,
	}

	if len(entity.UUID) > 0 {
		activity.UUID = uuid.MustParse(entity.UUID)
	}

	if entity.ID > 0 {
		// gorm.Model fields
		activity.ID = entity.ID
		activity.CreatedAt = entity.CreatedAt
		activity.UpdatedAt = entity.UpdatedAt

		if !entity.DeletedAt.IsZero() {
			activity.DeletedAt = entity.DeletedAt
		}
	}
	return activity
}

func toDomainModel(entity *Activity) domain.Activity {
	return domain.Activity{
		ID:          entity.ID,
		UUID:        entity.UUID.String(),
		Title:       entity.Title,
		Subtitle:    entity.Subtitle,
		Excerpt:     entity.Excerpt,
		Description: entity.Description,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
		DeletedAt:   entity.DeletedAt,
	}
}
