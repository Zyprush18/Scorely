package servicerole

import (
	"errors"
	"testing"
	"time"

	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/models/response"
	"github.com/stretchr/testify/assert"
)


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

		mock.AssertExpectations(t)
	})

	t.Run("Failed Get All Data Role", func(t *testing.T) {
		mock := new(RepoRoleMock)
		service := NewRoleService(mock)

		mock.On("GetAllDataRole").Return([]response.Roles(nil), errors.New("Database is refused"))

		resperr, err := service.GetAllData()
		assert.Error(t, err)
		assert.Nil(t, resperr)

		mock.AssertExpectations(t)
	})

}

func TestCreateServiceRole(t *testing.T) {
	mock := new(RepoRoleMock)
	service := NewRoleService(mock)
	t.Run("Service Success Create a New Role", func(t *testing.T) {
		rolePass := &request.Roles{
			NameRole: "Admin",
		}
		mock.On("CreateRole", rolePass).Return(nil)

		err := service.Create(rolePass)
		assert.NoError(t, err)
		mock.AssertExpectations(t)

	})

	t.Run("Service Failed Create a New Role", func(t *testing.T) {
		roleFails := &request.Roles{
			NameRole: "",
		}
		mock.On("CreateRole", roleFails).Return(errors.New("failed"))
		errs := service.Create(roleFails)

		assert.Error(t, errs)
		mock.AssertExpectations(t)
	})

}

func TestShowRoleById(t *testing.T) {
	mock := new(RepoRoleMock)
	servicerole := NewRoleService(mock)
	data := &response.Roles{
		IdRole:   1,
		NameRole: "Admin",
	}
	t.Run("Success Show Role by id", func(t *testing.T) {

		mock.On("ShowById", 1).Return(data, nil)

		resp, err := servicerole.ShowRoleById(1)
		assert.NoError(t, err)
		assert.Equal(t, uint(1), resp.IdRole)
		assert.Equal(t, "Admin", resp.NameRole)

		mock.AssertExpectations(t)
	})

	t.Run("Failed Show Role by id", func(t *testing.T) {

		mock.On("ShowById", 2).Return(data, errors.New("Not Found role id: 2"))

		resp, err := servicerole.ShowRoleById(2)

		assert.Error(t, err)
		assert.Nil(t, resp)

		mock.AssertExpectations(t)
	})
}

func TestUpdateRole(t *testing.T)  {
	
}
