package database

import (
	"github.com/google/uuid"
	"github.com/sumelms/microservice-activity/internal/activitylog/domain"
)

func toDBModel(entity *domain.ActivityLog) ActivityLog {
	al := ActivityLog{
		Title:       entity.Title,
		Description: entity.Description,
		ContentID:   uuid.MustParse(entity.ContentID),
		ContentType: entity.ContentType,
	}

	if len(entity.UUID) > 0 {
		al.UUID = uuid.MustParse(entity.UUID)
	}

	if entity.ID > 0 {
		// gorm.Model fields
		al.ID = entity.ID
		al.CreatedAt = entity.CreatedAt
		al.UpdatedAt = entity.UpdatedAt

		if !entity.DeletedAt.IsZero() {
			al.DeletedAt = entity.DeletedAt
		}
	}
	return al
}

func toDomainModel(entity *ActivityLog) domain.ActivityLog {
	return domain.ActivityLog{
		ID:          entity.ID,
		UUID:        entity.UUID.String(),
		Title:       entity.Title,
		Description: entity.Description,
		ContentID:   entity.ContentID.String(),
		ContentType: entity.ContentType,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
		DeletedAt:   entity.DeletedAt,
	}
}
