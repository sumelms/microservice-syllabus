package domain

type Repository interface {
	Create(*Activity) (Activity, error)
	Find(string) (Activity, error)
	Update(*Activity) (Activity, error)
	Delete(string) error
	List() ([]Activity, error)
}
