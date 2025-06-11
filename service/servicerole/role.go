package servicerole

import (
	"fmt"

	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/models/response"
	"github.com/Zyprush18/Scorely/repository/reporole"
)

type ServiceRole interface {
	GetAllData() ([]response.Roles, error)
	Create(data *request.Roles) error
	ShowRoleById(id int) (*response.Roles, error)
}

type RoleRepo struct {
	Repo reporole.RoleService
}

func NewRoleService(r reporole.RoleService) ServiceRole  {
	return &RoleRepo{Repo: r}
}

func (r *RoleRepo) GetAllData() ([]response.Roles, error){
	fmt.Println("cuyyyy gua di panggil")
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