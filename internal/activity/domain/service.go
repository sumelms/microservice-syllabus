package domain

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/log"
)

type ServiceInterface interface {
	ListActivity(context.Context, map[string]interface{}) ([]Activity, error)
	CreateActivity(context.Context, *Activity) (Activity, error)
	FindActivity(context.Context, string) (Activity, error)
	UpdateActivity(context.Context, *Activity) (Activity, error)
	DeleteActivity(context.Context, string) error
}

type Service struct {
	repo   Repository
	logger log.Logger
}

func NewService(repo Repository, logger log.Logger) *Service {
	return &Service{
		repo:   repo,
		logger: logger,
	}
}

func (s *Service) ListActivity(_ context.Context, filters map[string]interface{}) ([]Activity, error) {
	cs, err := s.repo.List(filters)
	if err != nil {
		return []Activity{}, fmt.Errorf("Service didn't found any activity: %w", err)
	}
	return cs, nil
}

func (s *Service) CreateActivity(_ context.Context, activity *Activity) (Activity, error) {
	c, err := s.repo.Create(activity)
	if err != nil {
		return Activity{}, fmt.Errorf("Service can't create activity: %w", err)
	}
	return c, nil
}

func (s *Service) FindActivity(_ context.Context, id string) (Activity, error) {
	c, err := s.repo.Find(id)
	if err != nil {
		return Activity{}, fmt.Errorf("Service can't find activity: %w", err)
	}
	return c, nil
}

func (s *Service) UpdateActivity(_ context.Context, activity *Activity) (Activity, error) {
	c, err := s.repo.Update(activity)
	if err != nil {
		return Activity{}, fmt.Errorf("Service can't update activity: %w", err)
	}
	return c, nil
}

func (s *Service) DeleteActivity(_ context.Context, id string) error {
	err := s.repo.Delete(id)
	if err != nil {
		return fmt.Errorf("Service can't delete activity: %w", err)
	}
	return nil
}
