package domain

type Repository interface {
	Create(*ActivityLog) (ActivityLog, error)
	Find(string) (ActivityLog, error)
	Update(*ActivityLog) (ActivityLog, error)
	Delete(string) error
	List(map[string]interface{}) ([]ActivityLog, error)
}
