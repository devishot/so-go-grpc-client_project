package domain

type ClientService struct {
	Repo ClientRepository
}

func (s ClientService) Create(data ClientEntity) (cl ClientEntity, err error) {
	cl = NewClient(data.FirstName, data.LastName, data.CompanyName)

	err = s.Repo.Create(cl)
	if err != nil {
		return
	}

	return
}

func (s ClientService) Delete(clientID ID) (cl ClientEntity, err error) {
	cl, err = s.Repo.Get(clientID)
	if err != nil {
		return
	}

	err = s.Repo.Delete(clientID)
	if err != nil {
		return
	}

	return
}

func (s ClientService) Get(clientID ID) (cl ClientEntity, err error) {
	cl, err = s.Repo.Get(clientID)
	if err != nil {
		return
	}

	return
}
