package domain

import "github.com/google/uuid"

type ActivityRepository interface {
	Activity(id uuid.UUID) (Activity, error)
	Activities() ([]Activity, error)
	CreateActivity(activity *Activity) error
	UpdateActivity(activity *Activity) error
	DeleteActivity(id uuid.UUID) error
}
