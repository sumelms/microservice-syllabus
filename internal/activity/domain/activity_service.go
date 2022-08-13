package domain

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (s *Service) Activity(_ context.Context, id uuid.UUID) (Activity, error) {
	a, err := s.activities.Activity(id)
	if err != nil {
		return Activity{}, fmt.Errorf("service can't find activity: %w", err)
	}
	return a, nil
}

func (s *Service) Activities(_ context.Context) ([]Activity, error) {
	aa, err := s.activities.Activities()
	if err != nil {
		return []Activity{}, fmt.Errorf("service didn't found any activity: %w", err)
	}
	return aa, nil
}

func (s *Service) CreateActivity(_ context.Context, activity *Activity) error {
	if err := s.activities.CreateActivity(activity); err != nil {
		return fmt.Errorf("service can't create activity: %w", err)
	}
	return nil
}

func (s *Service) UpdateActivity(_ context.Context, activity *Activity) error {
	if err := s.activities.UpdateActivity(activity); err != nil {
		return fmt.Errorf("service can't update activity: %w", err)
	}
	return nil
}

func (s *Service) DeleteActivity(_ context.Context, id uuid.UUID) error {
	err := s.activities.DeleteActivity(id)
	if err != nil {
		return fmt.Errorf("service can't delete activity: %w", err)
	}
	return nil
}
