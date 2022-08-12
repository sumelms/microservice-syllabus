package domain

import (
	"context"

	"github.com/go-kit/log"
	"github.com/google/uuid"
)

// ServiceInterface defines the domains Service interface
type ServiceInterface interface {
	Lesson(ctx context.Context, id uuid.UUID) (Lesson, error)
	Lessons(ctx context.Context) ([]Lesson, error)
	CreateLesson(ctx context.Context, c *Lesson) error
	UpdateLesson(ctx context.Context, c *Lesson) error
	DeleteLesson(ctx context.Context, id uuid.UUID) error
}

type serviceConfiguration func(svc *Service) error

type Service struct {
	lessons LessonRepository
	logger  log.Logger
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

// WithLessonRepository injects the course repository to the domain Service
func WithLessonRepository(lr LessonRepository) serviceConfiguration {
	return func(svc *Service) error {
		svc.lessons = lr
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
