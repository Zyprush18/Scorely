package servicerole

import (
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/repository/reporole"
)

type ServiceRole interface {
	Create(data *request.Roles) error
}

type RoleRepo struct {
	// dependency
	Repo reporole.RoleMysql
}

func RoleService(r reporole.RoleMysql) ServiceRole  {
	// injcetion
	return &RoleRepo{Repo: r}
}

func (r *RoleRepo) Create(data *request.Roles) error  {
	err := r.Repo.CreateRole(data)
	if err != nil {
		return err
	}

	return nil
}