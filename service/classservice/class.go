package classservice

import (
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/models/response"
	"github.com/Zyprush18/Scorely/repository/repoclass"
)

type ServiceClass interface {
	AllData(Search,Sort string, Page,Perpage int)([]response.Class, int64, error)
	CreateClass(data *request.Class) error
	ShowClass(id int)(*response.Class, error)
	UpdateClass(id int, data *request.Class) error
	DeleteClass(id int) error
}

type ClassService struct {
	repo repoclass.RepoClass
}

func NewClassService(r repoclass.RepoClass) ServiceClass  {
	return &ClassService{repo: r}
}

func (c *ClassService) AllData(Search,Sort string, Page,Perpage int) ([]response.Class, int64, error) {
	return c.repo.GetAll(Search,Sort,Page,Perpage)
}

func (c *ClassService) CreateClass(data *request.Class) error {
	return c.repo.Create(data)
}

func (c *ClassService) ShowClass(id int)(*response.Class, error) {
	return c.repo.Show(id)
}

func (c *ClassService) UpdateClass(id int, data *request.Class) error {
	return c.repo.Update(id, data)
}

func (c *ClassService) DeleteClass(id int) error {
	return c.repo.Delete(id)
}