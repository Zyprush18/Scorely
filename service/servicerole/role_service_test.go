package servicerole

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/models/response"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var Mockservice = RepoRoleMock{Mock: mock.Mock{}}

func TestGetAllData(t *testing.T) {
	t.Run("Success Get All Data Role", func(t *testing.T) {
		mock := new(RepoRoleMock)
		service := NewRoleService(mock)
		data := []response.Roles{
			{
				IdRole:   1,
				NameRole: "Admin",
				Users: []response.Users{
					{
						IdUser:   1,
						Email:    "admin@gmail.com",
						Password: "admin123",
						RoleId:   1,
					},
					{
						IdUser:   2,
						Email:    "admin2@gmail.com",
						Password: "admin123",
						RoleId:   1,
					},
				},
				Models: helper.Models{
					CreatedAt: time.Now(),
				},
			},
			{
				IdRole:   2,
				NameRole: "User",
				Users: []response.Users{
					{
						IdUser:   3,
						Email:    "user@gmail.com",
						Password: "user123",
						RoleId:   2,
					},
					{
						IdUser:   4,
						Email:    "user2@gmail.com",
						Password: "user123",
						RoleId:   2,
					},
				},
				Models: helper.Models{
					CreatedAt: time.Now(),
				},
			},
		}

		mock.On("GetAllDataRole").Return(data, nil)

		resp, err := service.GetAllData()
		assert.NoError(t, err)
		assert.NotNil(t, resp)

		Mockservice.AssertExpectations(t)
	})

	t.Run("Failed Get All Data Role", func(t *testing.T) {
		service := NewRoleService(&Mockservice)

		Mockservice.On("GetAllDataRole").Return(nil, errors.New("Database is refused"))

		resperr, err := service.GetAllData()
		fmt.Println(resperr)
		fmt.Println(err)

		assert.Error(t, err)
		assert.Nil(t, resperr)

		Mockservice.AssertExpectations(t)
	})

}

func TestCreateServiceRole(t *testing.T) {
	service := NewRoleService(&Mockservice)
	t.Run("Service Success Create a New Role", func(t *testing.T) {
		rolePass := &request.Roles{
			NameRole: "Admin",
		}
		Mockservice.On("CreateRole", rolePass).Return(nil)

		err := service.Create(rolePass)
		assert.NoError(t, err)
		Mockservice.AssertExpectations(t)

	})

	t.Run("Service Failed Create a New Role", func(t *testing.T) {
		roleFails := &request.Roles{
			NameRole: "",
		}
		Mockservice.On("CreateRole", roleFails).Return(errors.New("failed"))
		errs := service.Create(roleFails)

		assert.Error(t, errs)
		Mockservice.AssertExpectations(t)
	})

}

func TestShowRoleById(t *testing.T) {
	servicerole := NewRoleService(&Mockservice)
	data := &response.Roles{
		IdRole:   1,
		NameRole: "Admin",
	}
	t.Run("Success Show Role by id", func(t *testing.T) {

		Mockservice.On("ShowById", 1).Return(data, nil)

		resp, err := servicerole.ShowRoleById(1)
		assert.NoError(t, err)
		assert.Equal(t, uint(1), resp.IdRole)
		assert.Equal(t, "Admin", resp.NameRole)

		Mockservice.AssertExpectations(t)
	})

	t.Run("Failed Show Role by id", func(t *testing.T) {

		Mockservice.On("ShowById", 2).Return(data, errors.New("Not Found role id: 2"))

		resp, err := servicerole.ShowRoleById(2)

		assert.Error(t, err)
		assert.Nil(t, resp)

		Mockservice.AssertExpectations(t)
	})
}
