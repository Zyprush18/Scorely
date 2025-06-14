package userservice

import (
	"errors"
	"testing"

	"github.com/Zyprush18/Scorely/models/request"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateUserService(t *testing.T)  {
	mockuser := UserRepository{Mock: mock.Mock{}}
	userservice := NewUserService(&mockuser)

	t.Run("Success Create a New User", func(t *testing.T) {
		reqUser := &request.User{
			Email: "Admin@gmail.com",
			Password: "admin123",
			RoleId: 1,
		}

		mockuser.On("Create", reqUser).Return(nil)
		err := userservice.CreateUser(reqUser)
		assert.NoError(t, err)

		mockuser.AssertExpectations(t)
	})

	t.Run("Failed Create a New User", func(t *testing.T) {
		reqUser := &request.User{
			Email: "",
		}
		mockuser.On("Create", reqUser).Return(errors.New("Cannot Add or Update child row"))
		err := userservice.CreateUser(reqUser)
		assert.Error(t, err)

		mockuser.AssertExpectations(t)
	})
}