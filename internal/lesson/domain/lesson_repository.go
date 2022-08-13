package domain

import "github.com/google/uuid"

type LessonRepository interface {
	Lesson(id uuid.UUID) (Lesson, error)
	Lessons() ([]Lesson, error)
	CreateLesson(lesson *Lesson) error
	UpdateLesson(lesson *Lesson) error
	DeleteLesson(id uuid.UUID) error
	AddActivity(lessonActivity *LessonActivity) error
	RemoveActivity(lessonID, activityID uuid.UUID) error
}
