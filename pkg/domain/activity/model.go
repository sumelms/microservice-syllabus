package activity

type Activity struct {
	ID          uint
	UUID        string
	Name        string
	Description string
	LessonID    string
	ContentID   string
	ContentType string
	CreatedAt   int64
	UpdatedAt   int64
	DeletedAt   *int64
	PublishedAt *int64
}
