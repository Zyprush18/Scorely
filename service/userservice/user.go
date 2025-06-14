package userservice

import (
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/repository/repouser"
)

type ServiceUser interface {
	CreateUser(data *request.User) error
}

type UserRepo struct {
	repo repouser.UserRepo
}

func NewUserService(r repouser.UserRepo) ServiceUser  {
	return &UserRepo{repo: r}
}

func (u *UserRepo) CreateUser(data *request.User) error {
	if err:= u.repo.Create(data);err != nil {
		return err
	}

	return nil
}