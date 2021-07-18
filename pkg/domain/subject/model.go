package subject

type Subject struct {
	ID          uint
	UUID        string
	Name        string
	Description *string
	CreatedAt   int64
	UpdatedAt   int64
	DeletedAt   *int64
	PublishedAt *int64
}
