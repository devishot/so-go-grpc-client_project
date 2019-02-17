package domain

type ProjectRepository interface {
	Get(id ID) (ProjectEntity, error)
	Delete(id ID) error
	Create(entity ProjectEntity) error
	GetByClient(clientID ID) ([]ProjectEntity, error)
}
