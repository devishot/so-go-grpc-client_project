package domain

//go:generate moq -out client_repository__gen_mock.go . ClientRepository
type ClientRepository interface {
	Get(id ID) (ClientEntity, error)
	Delete(id ID) error
	Create(entity ClientEntity) error
}
