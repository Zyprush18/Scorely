package subjectservice

import (
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/models/response"
	"github.com/Zyprush18/Scorely/repository/reposubject"
)

type ServiceSubject interface {
	GetAllSubject(Search,Sort string, Page,Perpage int) ([]response.Subjects, int64, error)
	CreateSubject(data *request.Subjects) error
	ShowSubject(id int) (*response.Subjects, error)
	UpdateSubject(id int,data *request.Subjects) error
	DeleteSubject(id int) error
}

type RepoStruct struct {
	repo reposubject.RepoSubject
}

func ConnectRepo(r reposubject.RepoSubject) ServiceSubject  {
	return &RepoStruct{repo: r}
}

func (r *RepoStruct) GetAllSubject(Search,Sort string, Page,Perpage int) ([]response.Subjects, int64, error)  {
	return r.repo.GetAll(Search,Sort,Page,Perpage)
}

func (r *RepoStruct) CreateSubject(data *request.Subjects) error {
	return r.repo.Create(data)
}

func (r *RepoStruct) ShowSubject(id int) (*response.Subjects, error) {
	return r.repo.Show(id)
}

func (r *RepoStruct) UpdateSubject(id int,data *request.Subjects) error {
	return r.repo.Update(id, data)
}

func (r *RepoStruct) DeleteSubject(id int) error {
	return r.repo.Delete(id)
}