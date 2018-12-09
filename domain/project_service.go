package domain

type ProjectService struct {
	Repo ProjectRepository
}

func (s *ProjectService) Create(in ProjectInputValue) error {
	p := NewProject(in.ClientID, in.Title, in.Description)

	return s.Repo.Create(p)
}

func (s *ProjectService) Delete(id ID) error {
	return s.Repo.Delete(id)
}
