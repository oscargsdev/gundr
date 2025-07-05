package models

type Project struct {
	ID          int
	ProjectName string
	Genre       string
	Location    string
	Status      string
}

type ProjectModel struct {
}

func (m *ProjectModel) Get(id int) (Project, error) {
	return Project{
		ID:          id,
		ProjectName: "Test Project",
	}, nil
}
