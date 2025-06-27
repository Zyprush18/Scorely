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

type StructTest struct {
	Name        string
	Data        []response.Roles
	Request     *request.Roles
	Counts      int
	Page        int
	Perpage     int
	Search      string
	Sort        string
	ExpectedErr error
}

func TestGetAlldatasRole(t *testing.T) {
	mock := new(RepoRoleMock)
	service := NewRoleService(mock)
	data := []StructTest{
		{
			Name: "Success Get Data Role",
			Data: []response.Roles{
				{
					IdRole:   1,
					NameRole: "Admin",
					Models: helper.Models{
						CreatedAt: time.Now(),
					},
				},
				{
					IdRole:   2,
					NameRole: "User",
					Models: helper.Models{
						CreatedAt: time.Now(),
					},
				},
			},
			Counts:      1,
			Page:        1,
			Perpage:     10,
			Search:      "",
			Sort:        "",
			ExpectedErr: nil,
		},
		{
			Name:        "Failed Get Data Role",
			Data:        []response.Roles(nil),
			Counts:      0,
			Page:        1,
			Perpage:     10,
			Search:      "",
			Sort:        "",
			ExpectedErr: errors.New("Database is refused"),
		},
	}

	for _, v := range data {
		val := v
		t.Run(val.Name, func(t *testing.T) {

			mock.ExpectedCalls = nil
			mock.Calls = nil
			
			mock.On("GetAllDataRole", v.Search, v.Sort, v.Page, v.Perpage).Return(val.Data, v.Counts, val.ExpectedErr)
			resp, count, err := service.GetAllData(v.Search, v.Sort, v.Page, v.Perpage)
			if val.ExpectedErr != nil {
				assert.Error(t, err)
				assert.Nil(t, resp)
				assert.Equal(t, int64(v.Counts), count)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, val.Data, resp)
				assert.Equal(t, int64(v.Counts), count)
			}
			mock.AssertExpectations(t)
		})
	}

}

func TestCreateServiceRole(t *testing.T) {
	mock := new(RepoRoleMock)
	service := NewRoleService(mock)

	data := []StructTest{
		{
			Name: "Service Success Create a New Role",
			Request: &request.Roles{
				NameRole: "Admin",
			},
			ExpectedErr: nil,
		},
		{
			Name: "Service Failed Create a New Role",
			Request: &request.Roles{
				NameRole: "User",
			},
			ExpectedErr: errors.New("Failed Create Role"),
		},
	}

	for _, v := range data {
		t.Run(v.Name,func(t *testing.T) {
			mock.On("CreateRole", v.Request).Return(v.ExpectedErr)
			err := service.Create(v.Request)
			if v.ExpectedErr != nil {
				assert.Error(t, err)
			}else{
				assert.NoError(t, err)
			}

			mock.AssertExpectations(t)
		})
	}
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

		mock.On("ShowById", 2).Return((*response.Roles)(nil), errors.New("Not Found role id: 2"))

		resp, err := servicerole.ShowRoleById(2)

		assert.Error(t, err)
		assert.Nil(t, resp)

		mock.AssertExpectations(t)
	})
}

func TestUpdateRole(t *testing.T) {
	mock := new(RepoRoleMock)
	servicerole := NewRoleService(mock)

	data := &request.Roles{
		NameRole: "Admin",
	}
	t.Run("Success Update Role", func(t *testing.T) {
		mock.On("UpdateRole", 1, data).Return(nil)

		err := servicerole.UpdateRole(1, data)
		assert.NoError(t, err)

		mock.AssertExpectations(t)
	})

	t.Run("Failed Delete Role", func(t *testing.T) {
		mock.On("UpdateRole", 90, data).Return(errors.New("Not Found Role Id: 90"))

		err := servicerole.UpdateRole(90, data)

		assert.Error(t, err)

		mock.AssertExpectations(t)
	})
}

func TestDeleteRole(t *testing.T) {
	mock := new(RepoRoleMock)
	servicerepo := NewRoleService(mock)

	t.Run("Success Delete Role", func(t *testing.T) {
		mock.On("DeleteRole", 1).Return(nil)

		err := servicerepo.DeleteRole(1)
		assert.NoError(t, err)

		mock.AssertExpectations(t)
	})

	t.Run("Failed Delete Role", func(t *testing.T) {
		mock.On("DeleteRole", 3).Return(errors.New("Not Found Id_Role: 3"))

		err := servicerepo.DeleteRole(3)
		assert.Error(t, err)

		mock.AssertExpectations(t)
	})
}
