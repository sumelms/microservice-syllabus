package domain

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (s *Service) Lesson(_ context.Context, id uuid.UUID) (Lesson, error) {
	c, err := s.lessons.Lesson(id)
	if err != nil {
		return Lesson{}, fmt.Errorf("service can't find lesson: %w", err)
	}
	return c, nil
}

func (s *Service) Lessons(_ context.Context) ([]Lesson, error) {
	cc, err := s.lessons.Lessons()
	if err != nil {
		return []Lesson{}, fmt.Errorf("service didn't found any lesson: %w", err)
	}
	return cc, nil
}

func (s *Service) CreateLesson(_ context.Context, l *Lesson) error {
	if err := s.lessons.CreateLesson(l); err != nil {
		return fmt.Errorf("service can't create lesson: %w", err)
	}
	return nil
}

func (s *Service) UpdateLesson(_ context.Context, l *Lesson) error {
	if err := s.lessons.UpdateLesson(l); err != nil {
		return fmt.Errorf("service can't update lesson: %w", err)
	}
	return nil
}

func (s *Service) DeleteLesson(_ context.Context, id uuid.UUID) error {
	if err := s.lessons.DeleteLesson(id); err != nil {
		return fmt.Errorf("service can't delete lesson: %w", err)
	}
	return nil
}

func (s *Service) AddActivity(ctx context.Context, la *LessonActivity) error {
	if err := s.lessons.AddActivity(la); err != nil {
		return fmt.Errorf("service can't add activity to lesson: %w", err)
	}
	return nil
}

func (s *Service) RemoveActivity(_ context.Context, lessonID, activityID uuid.UUID) error {
	if err := s.lessons.RemoveActivity(lessonID, activityID); err != nil {
		return fmt.Errorf("service can't remove activity from lesson: %w", err)
	}
	return nil
}
