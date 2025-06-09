package servicerole

import (
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/repository/reporole"
)

type ServiceRole interface {
	Create(data *request.Roles) error
}

type RoleRepo struct {
	Repo reporole.RoleService
}

func NewRoleService(r reporole.RoleService) ServiceRole  {
	return &RoleRepo{Repo: r}
}

func (r *RoleRepo) Create(data *request.Roles) error  {
	return r.Repo.CreateRole(data)
}