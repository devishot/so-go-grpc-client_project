package domain

type ClientRepository interface {
	Get(id ID) (ClientEntity, error)
	Delete(id ID) error
	Create(entity ClientEntity) error
}

type ProjectRepository interface {
	Get(id ID) (ProjectEntity, error)
	Delete(id ID) error
	Create(entity ProjectEntity) error
	GetByClient(clientID ID) ([]ProjectEntity, error)
}
