package servicerole

import (
	"context"

	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/models/response"
	"github.com/Zyprush18/Scorely/repository/reporole"
)

type ServiceRole interface {
	GetAllData(ctx context.Context, search,sort string,page,perpage int) ([]response.Roles, int64,error)
	Create(ctx context.Context, data *request.Roles) error
	ShowRoleById(ctx context.Context, id int) (*response.Roles, error)
	UpdateRole(ctx context.Context, id int, data *request.Roles) error
	DeleteRole(ctx context.Context, id int) error
}

type RoleRepo struct {
	Repo reporole.RoleService
}

func NewRoleService(r reporole.RoleService) ServiceRole  {
	return &RoleRepo{Repo: r}
}

func (r *RoleRepo) GetAllData(ctx context.Context, search,sort string,page,perpage int) ([]response.Roles, int64,error){
	return r.Repo.GetAllDataRole(ctx, search,sort,page,perpage)
}

func (r *RoleRepo) Create(ctx context.Context, data *request.Roles) error  {
	return r.Repo.CreateRole(ctx, data)
}

	
func (r *RoleRepo) ShowRoleById(ctx context.Context, id int) (*response.Roles, error) {
	return r.Repo.ShowById(ctx, id)
}

func (r *RoleRepo) UpdateRole(ctx context.Context, id int, data *request.Roles) error  {
	return r.Repo.UpdateRole(ctx, id,data)
}

func (r *RoleRepo) DeleteRole(ctx context.Context, id int) error {
	return r.Repo.DeleteRole(ctx, id)
}