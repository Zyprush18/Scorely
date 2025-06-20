package userservice

import (
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/models/response"
	"github.com/Zyprush18/Scorely/repository/repouser"
)

type ServiceUser interface {
	AllUser()([]response.Users, error)
	CreateUser(data *request.User) error
	ShowUser(id int) (*response.Users, error)
	UpdateUser(id int, data *request.User) error
}

type UserRepo struct {
	repo repouser.UserRepo
}

func NewUserService(r repouser.UserRepo) ServiceUser  {
	return &UserRepo{repo: r}
}

func (u *UserRepo) AllUser()([]response.Users, error) {
	data, err := u.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (u *UserRepo) CreateUser(data *request.User) error {
	if err:= u.repo.Create(data);err != nil {
		return err
	}

	return nil
}

func (u *UserRepo) ShowUser(id int) (*response.Users, error) {
	data, err:= u.repo.Show(id);
	if err != nil {
		return nil, err
	}

	return data,nil
}

func (u *UserRepo) UpdateUser(id int, data *request.User) error  {
	if err:= u.repo.Update(id, data);err != nil {
		return err
	}

	return nil
}