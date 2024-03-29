package domain

import (
	"context"

	"github.com/go-kit/log"
	"github.com/google/uuid"
)

// ServiceInterface defines the domains Service interface
type ServiceInterface interface {
	Activity(ctx context.Context, id uuid.UUID) (Activity, error)
	Activities(ctx context.Context) ([]Activity, error)
	CreateActivity(ctx context.Context, activity *Activity) error
	UpdateActivity(ctx context.Context, activity *Activity) error
	DeleteActivity(ctx context.Context, id uuid.UUID) error
}

type serviceConfiguration func(svc *Service) error

type Service struct {
	activities ActivityRepository
	logger     log.Logger
}

// NewService creates a new domain Service instance
func NewService(cfgs ...serviceConfiguration) (*Service, error) {
	svc := &Service{}
	for _, cfg := range cfgs {
		err := cfg(svc)
		if err != nil {
			return nil, err
		}
	}
	return svc, nil
}

// WithActivityRepository injects the course repository to the domain Service
func WithActivityRepository(ar ActivityRepository) serviceConfiguration {
	return func(svc *Service) error {
		svc.activities = ar
		return nil
	}
}

// WithLogger injects the logger to the domain Service
func WithLogger(l log.Logger) serviceConfiguration {
	return func(svc *Service) error {
		svc.logger = l
		return nil
	}
}
