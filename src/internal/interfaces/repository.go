package interfaces

type Repository[T Entity] interface {
	Create(itemDAO DAO) (T, error)
	Persist(entity T) error
	FindById(id string) (T, error)
}
