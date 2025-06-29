package userservice

import (
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/models/response"
	"github.com/stretchr/testify/mock"
)

type UserRepository struct {
	mock.Mock
}
func (u *UserRepository) GetAll(search, sort string, page,perpage int)([]response.Users, int64,error) {
	args := u.Called(search,sort,page,perpage)
	return args.Get(0).([]response.Users), int64(args.Int(1)),args.Error(2)	
}

func (u *UserRepository) Create(data *request.User) error {
	args := u.Called(data)
	return args.Error(0)
}

func (u *UserRepository) Show(id int) (*response.Users, error)  {
	args := u.Called(id)
	return args.Get(0).(*response.Users), args.Error(1)
}

func (u *UserRepository) Update(id int, data *request.User) error  {
	args := u.Called(id, data)
	return args.Error(0)
}

	
func (u *UserRepository) Delete(id int) error {
	args := u.Called(id)
	return args.Error(0)
}