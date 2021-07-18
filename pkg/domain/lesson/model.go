package lesson

type Lesson struct {
	ID          uint
	UUID        string
	Name        string
	Module      *string
	SubjectID   string
	CreatedAt   int64
	UpdatedAt   int64
	DeletedAt   *int64
	PublishedAt *int64
}
