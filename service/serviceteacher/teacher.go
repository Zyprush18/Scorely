package serviceteacher

import (
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/models/response"
	"github.com/Zyprush18/Scorely/repository/repoteacher"
)

type ServiceTeacher interface {
	GetAllTeacher(Search, Sort string, Page,Perpage int) ([]response.Teachers, int64, error)
	CreateTeacher(data *request.Teachers) error
}

type RepoTeacherStruct struct {
	repo repoteacher.RepoTeacher
}

func ConnectRepo(r repoteacher.RepoTeacher) ServiceTeacher  {
	return &RepoTeacherStruct{repo: r}
}

func (r *RepoTeacherStruct) GetAllTeacher(Search, Sort string, Page,Perpage int) ([]response.Teachers, int64, error) {
	return r.repo.GetAll(Search,Sort,Page,Perpage)
}

func (r *RepoTeacherStruct) CreateTeacher(data *request.Teachers) error {
	return r.repo.Create(data)
}