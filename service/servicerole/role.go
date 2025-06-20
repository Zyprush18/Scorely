package servicerole

import (

	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/models/response"
	"github.com/Zyprush18/Scorely/repository/reporole"
)

type ServiceRole interface {
	GetAllData() ([]response.Roles, error)
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

func (r *RoleRepo) GetAllData() ([]response.Roles, error){
	return r.Repo.GetAllDataRole()
}

func (r *RoleRepo) Create(data *request.Roles) error  {
	return r.Repo.CreateRole(data)
}

	
func (r *RoleRepo) ShowRoleById(id int) (*response.Roles, error) {
	resp, err := r.Repo.ShowById(id)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *RoleRepo) UpdateRole(id int, data *request.Roles) error  {
	err := r.Repo.UpdateRole(id,data)
	if err != nil {
		return err
	}

	return nil
}

func (r *RoleRepo) DeleteRole(id int) error {
	err:= r.Repo.DeleteRole(id)
	if err != nil {
		return err
	}

	return nil
}