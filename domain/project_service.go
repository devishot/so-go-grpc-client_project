package domain

type ProjectService struct {
	Repo ProjectRepository
}

func (s ProjectService) Create(in ProjectEntity) (p ProjectEntity, err error) {
	p = NewProject(in.ClientID, in.Title, in.Description)

	err = s.Repo.Create(p)
	if err != nil {
		return
	}

	return
}

func (s ProjectService) Delete(id ID) (p ProjectEntity, err error) {
	p, err = s.Repo.Get(id)
	if err != nil {
		return
	}

	err = s.Repo.Delete(id)
	if err != nil {
		return
	}

	return
}

func (s ProjectService) Get(ID ID) (p ProjectEntity, err error) {
	p, err = s.Repo.Get(ID)
	if err != nil {
		return
	}

	return
}
