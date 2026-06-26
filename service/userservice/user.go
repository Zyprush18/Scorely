package userservice

import (
	"context"
	"time"

	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/models/entity"
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/models/response"
	"github.com/Zyprush18/Scorely/repository/repouser"
)

type ServiceUser interface {
	AllUser(ctx context.Context, search, sort string, page, perpage int) ([]response.Users, int64, error)
	CreateUser(ctx context.Context, data *request.User) error
	ShowUser(ctx context.Context, id int) (*response.Users, error)
	UpdateUser(ctx context.Context, id int, data *request.User) error
	DeleteUser(ctx context.Context, id int) error
}

type UserRepo struct {
	repo repouser.UserRepo
}

func NewUserService(r repouser.UserRepo) ServiceUser {
	return &UserRepo{repo: r}
}

func (u *UserRepo) AllUser(ctx context.Context, search, sort string, page, perpage int) ([]response.Users, int64, error) {
	entities, count, err := u.repo.GetAll(ctx, search, sort, page, perpage)
	if err != nil {
		return nil, 0, err
	}
	return response.ParseUserResponse(entities), count, nil
}

func (u *UserRepo) CreateUser(ctx context.Context, data *request.User) error {
	ent := &entity.Users{
		Email:    data.Email,
		Password: helper.HashingPassword(data.Password),
		RoleId:   data.RoleId,
		Models: helper.Models{
			CreatedAt: time.Now().Local(),
		},
	}
	return u.repo.Create(ctx, ent)
}

func (u *UserRepo) ShowUser(ctx context.Context, id int) (*response.Users, error) {
	ent, err := u.repo.Show(ctx, id)
	if err != nil {
		return nil, err
	}
	return response.MapUserResp(ent), nil
}

func (u *UserRepo) UpdateUser(ctx context.Context, id int, data *request.User) error {
	ent := &entity.Users{
		Email:    data.Email,
		RoleId:   data.RoleId,
		Models: helper.Models{
			UpdatedAt: time.Now().Local(),
		},
	}

	if data.Password != "" {
		ent.Password = helper.HashingPassword(data.Password)
	}

	return u.repo.Update(ctx, id, ent)
}

func (u *UserRepo) DeleteUser(ctx context.Context, id int) error {
	return u.repo.Delete(ctx, id)
}


