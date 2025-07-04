package userservice

import (
	"errors"
	"testing"
	"time"

	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/models/response"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAllUser(t *testing.T) {
	t.Run("Success All Data Users", func(t *testing.T) {
		mockuser := UserRepository{Mock: mock.Mock{}}
		userservices := NewUserService(&mockuser)
		data := []response.Users{
			{
				IdUser:   1,
				Email:    "Admin@gmail.com",
				Password: "Admin123",
				RoleId:   1,
				Models: helper.Models{
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
			},
			{
				IdUser:   2,
				Email:    "Admin2@gmail.com",
				Password: "Admin123",
				RoleId:   2,
				Models: helper.Models{
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
			},
		}

		mockuser.On("GetAll").Return(data, nil)
		resp, _,err := userservices.AllUser("","",0,0)
		assert.NoError(t, err)
		assert.NotNil(t, resp)

		mockuser.AssertExpectations(t)

	})

	t.Run("Failed All Data User", func(t *testing.T) {
		mockuser := UserRepository{Mock: mock.Mock{}}
		userservices := NewUserService(&mockuser)
		mockuser.On("GetAll").Return([]response.Users(nil), errors.New("something went wrong"))
		resp, _,err := userservices.AllUser("","",0,0)
		assert.Error(t, err)
		assert.Nil(t, resp)

		mockuser.AssertExpectations(t)
	})
}

func TestCreateUserService(t *testing.T) {
	mockuser := UserRepository{Mock: mock.Mock{}}
	userservice := NewUserService(&mockuser)

	t.Run("Success Create a New User", func(t *testing.T) {
		reqUser := &request.User{
			Email:    "Admin@gmail.com",
			Password: "admin123",
			RoleId:   1,
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

func TestShowuserserviceById(t *testing.T) {
	mockuser := UserRepository{}
	serviceuser := NewUserService(&mockuser)

	t.Run("Success Show By id", func(t *testing.T) {
		dataUser := &response.Users{
			IdUser:   1,
			Email:    "Admin@gmail.com",
			Password: "admin123",
			RoleId:   1,
		}

		mockuser.On("Show", 1).Return(dataUser, nil)

		data, err := serviceuser.ShowUser(1)
		assert.NoError(t, err)

		assert.NotNil(t, data)
		assert.Equal(t, "Admin@gmail.com", data.Email)

		mockuser.AssertExpectations(t)
	})

	t.Run("Failed Show User by Id", func(t *testing.T) {
		dataUser := &response.Users{
			IdUser:   2,
			Email:    "Admin1@gmail.com",
			Password: "admin123",
			RoleId:   1,
		}

		mockuser.On("Show", 4).Return(dataUser, errors.New("Not Found Id user: 4"))

		data, err := serviceuser.ShowUser(4)
		assert.Error(t, err)

		assert.Nil(t, data)

		mockuser.AssertExpectations(t)
	})
}

func TestUpdateuserService(t *testing.T)  {
	mockrepousers := UserRepository{mock.Mock{}}
	service := NewUserService(&mockrepousers)

	t.Run("Success Update User", func(t *testing.T) {
		id_success := 1
		datareq := &request.User{
			Email: "Admin@gmail.com",
		}

		mockrepousers.On("Update", id_success, datareq).Return(nil)

		err := service.UpdateUser(id_success, datareq)
		assert.NoError(t, err)
		mockrepousers.AssertExpectations(t)
	})

	t.Run("Failed Update User", func(t *testing.T) {
		id_failed := 2
		datareq := &request.User{
			Email: "Admin123@gmail.com",
		}

		mockrepousers.On("Update", id_failed, datareq).Return(errors.New("not found id: 2"))

		err := service.UpdateUser(id_failed, datareq)
		assert.Error(t, err)
		mockrepousers.AssertExpectations(t)
	})
}

func TestDeleteUser(t *testing.T)  {
	usermock := UserRepository{Mock: mock.Mock{}}
	servceuser := NewUserService(&usermock)

	t.Run("Success Delete User", func(t *testing.T) {
		id_sucess := 1
		usermock.On("Delete", id_sucess).Return(nil)
		err := servceuser.DeleteUser(id_sucess)
		assert.NoError(t, err)

		usermock.AssertExpectations(t)
	})

	t.Run("Failed Delete User", func(t *testing.T) {
		id_failed := 2
		usermock.On("Delete", id_failed).Return(errors.New("Not Found id: 2"))
		err := servceuser.DeleteUser(id_failed)
		assert.Error(t, err)

		usermock.AssertExpectations(t)
	})
}