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

func (s ClientService) Delete(id ID) (cl ClientEntity, err error) {
	cl, err = s.Repo.Get(id)
	if err != nil {
		return
	}

	err = s.Repo.Delete(id)
	if err != nil {
		return
	}

	return
}

func (s ClientService) Get(id ID) (cl ClientEntity, err error) {
	cl, err = s.Repo.Get(id)
	if err != nil {
		return
	}

	return
}
