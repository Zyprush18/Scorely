package userservice

import (
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/stretchr/testify/mock"
)

type UserRepository struct {
	mock.Mock
}

func (u *UserRepository) Create(data *request.User) error {
	args := u.Called(data)
	return args.Error(0)
}