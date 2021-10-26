package domain

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/log"
)

type ServiceInterface interface {
	ListActivityLog(context.Context, map[string]interface{}) ([]ActivityLog, error)
	CreateActivityLog(context.Context, *ActivityLog) (ActivityLog, error)
	FindActivityLog(context.Context, string) (ActivityLog, error)
	UpdateActivityLog(context.Context, *ActivityLog) (ActivityLog, error)
	DeleteActivityLog(context.Context, string) error
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

func (s *Service) ListActivityLog(_ context.Context, filters map[string]interface{}) ([]ActivityLog, error) {
	cs, err := s.repo.List(filters)
	if err != nil {
		return []ActivityLog{}, fmt.Errorf("Service didn't found any activity log: %w", err)
	}
	return cs, nil
}

func (s *Service) CreateActivityLog(_ context.Context, activity *ActivityLog) (ActivityLog, error) {
	c, err := s.repo.Create(activity)
	if err != nil {
		return ActivityLog{}, fmt.Errorf("Service can't create activity log: %w", err)
	}
	return c, nil
}

func (s *Service) FindActivityLog(_ context.Context, id string) (ActivityLog, error) {
	c, err := s.repo.Find(id)
	if err != nil {
		return ActivityLog{}, fmt.Errorf("Service can't find activity log: %w", err)
	}
	return c, nil
}

func (s *Service) UpdateActivityLog(_ context.Context, activity *ActivityLog) (ActivityLog, error) {
	c, err := s.repo.Update(activity)
	if err != nil {
		return ActivityLog{}, fmt.Errorf("Service can't update activity log: %w", err)
	}
	return c, nil
}

func (s *Service) DeleteActivityLog(_ context.Context, id string) error {
	err := s.repo.Delete(id)
	if err != nil {
		return fmt.Errorf("Service can't delete activity log: %w", err)
	}
	return nil
}
