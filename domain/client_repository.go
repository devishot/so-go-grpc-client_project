package domain

type ClientRepository interface {
	Get(id ID) (ClientEntity, error)
	Delete(id ID) error
	Create(entity ClientEntity) error
}
