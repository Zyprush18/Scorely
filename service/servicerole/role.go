package servicerole

import (

	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/models/response"
	"github.com/Zyprush18/Scorely/repository/reporole"
)

type ServiceRole interface {
	GetAllData(search,sort string) ([]response.Roles, error)
	Create(data *request.Roles) error
	ShowRoleById(id int) (*response.Roles, error)
	UpdateRole(id int, data *request.Roles) error
	DeleteRole(id int) error
}

type RoleRepo struct {
	Repo reporole.RoleService
}

func NewRoleService(r reporole.RoleService) ServiceRole  {
	return &RoleRepo{Repo: r}
}

func (r *RoleRepo) GetAllData(search,sort string) ([]response.Roles, error){
	return r.Repo.GetAllDataRole(search,sort)
}

func (r *RoleRepo) Create(data *request.Roles) error  {
	return r.Repo.CreateRole(data)
}

	
func (r *RoleRepo) ShowRoleById(id int) (*response.Roles, error) {
	return r.Repo.ShowById(id)
}

func (r *RoleRepo) UpdateRole(id int, data *request.Roles) error  {
	return r.Repo.UpdateRole(id,data)
}

func (r *RoleRepo) DeleteRole(id int) error {
	return r.Repo.DeleteRole(id)
}