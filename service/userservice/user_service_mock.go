package userservice

import (
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/models/response"
	"github.com/stretchr/testify/mock"
)

type UserRepository struct {
	mock.Mock
}
func (u *UserRepository) GetAll() ([]response.Users, error) {
	args := u.Called()
	return args.Get(0).([]response.Users), args.Error(1)	
}

func (u *UserRepository) Create(data *request.User) error {
	args := u.Called(data)
	return args.Error(0)
}

func (u *UserRepository) Show(id int) (*response.Users, error)  {
	args := u.Called(id)
	return args.Get(0).(*response.Users), args.Error(1)
}

