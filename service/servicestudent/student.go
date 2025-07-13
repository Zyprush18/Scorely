package servicestudent

import (
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/models/response"
	"github.com/Zyprush18/Scorely/repository/repostudent"
)

type ServiceStudent interface {
	GetAllStudent(Search,Sort string, Page,Perpage int) ([]response.Students, int64, error)
	CreateStudent(data *request.Students) error
	ShowStudent(id int) (*response.Students, error)
	UpdateStudent(id int, data *request.Students) error
	DeleteStudent(id int) error
}

type RepoStudent struct {
	repo repostudent.StudentRepo
}

func NewServiceStudent(r repostudent.StudentRepo) ServiceStudent  {
	return &RepoStudent{repo: r}
}

func (r *RepoStudent) GetAllStudent(Search,Sort string, Page,Perpage int) ([]response.Students, int64, error) {
	return r.repo.GetAll(Search,Sort,Page,Perpage)
}

func (r *RepoStudent) CreateStudent(data *request.Students) error  {
	return  r.repo.Create(data)
}

func (r *RepoStudent) ShowStudent(id int)(*response.Students, error) {
	return r.repo.Show(id)
}

func (r *RepoStudent) UpdateStudent(id int, data *request.Students) error {
	return r.repo.Update(id, data)
}

func (r *RepoStudent) DeleteStudent(id int) error  {
	return r.repo.Delete(id)
}