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
	Id			int
	Response    *response.Roles
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
		t.Run(v.Name, func(t *testing.T) {
			mock.On("CreateRole", v.Request).Return(v.ExpectedErr)
			err := service.Create(v.Request)
			if v.ExpectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			mock.AssertExpectations(t)
		})
	}
}

func TestShowRoleById(t *testing.T) {
	mock := new(RepoRoleMock)
	servicerole := NewRoleService(mock)

	data := []StructTest{
		{
			Name: "Success Show Role",
			Id: 1,
			Response: &response.Roles{
				IdRole:   1,
				NameRole: "Admin",
			},
			ExpectedErr: nil,
		},
		{
			Name: "Failed Show Role",
			Id: 2,
			Response: nil,
			ExpectedErr: errors.New("Not Found Id: 2"),
		},
	}

	for _, v := range data {
		t.Run(v.Name, func(t *testing.T) {
			mock.On("ShowById", v.Id).Return(v.Response, v.ExpectedErr)
			resp, err := servicerole.ShowRoleById(v.Id)
			if v.ExpectedErr != nil {
				assert.Error(t, err)
				assert.Nil(t, resp)
			}else{
				assert.NoError(t, err)
				assert.Equal(t, uint(1), resp.IdRole)
				assert.Equal(t, "Admin", resp.NameRole)
			}

			mock.AssertExpectations(t)
		})
	}
}

func TestUpdateRole(t *testing.T) {
	mock := new(RepoRoleMock)
	servicerole := NewRoleService(mock)
	data := []StructTest{
		{
			Name: "Success Update Role",
			Id: 1,
			Request: &request.Roles{
				NameRole: "Admin",
			},
			ExpectedErr: nil,
		},
		{
			Name: "Failed Update Role",
			Id: 2,
			Request: nil,
			ExpectedErr: errors.New("Not Found Id: 2"),
		},
	}

	for _, v := range data {
		t.Run(v.Name, func(t *testing.T) {
			mock.On("UpdateRole", v.Id, v.Request).Return(v.ExpectedErr)
			err := servicerole.UpdateRole(v.Id,v.Request)
			if v.ExpectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			mock.AssertExpectations(t)
		})
	}
}

func TestDeleteRole(t *testing.T) {
	mock := new(RepoRoleMock)
	servicerepo := NewRoleService(mock)
	data := []StructTest{
		{
			Name: "Success Delete Role",
			Id: 1,
			ExpectedErr: nil,
		},
		{
			Name: "Failed Delete Role",
			Id: 2,
			ExpectedErr: errors.New("Not Found Id: 2"),
		},
	}

	for _, v := range data {
		t.Run(v.Name,func(t *testing.T) {
			mock.On("DeleteRole",v.Id).Return(v.ExpectedErr)
			err := servicerepo.DeleteRole(v.Id)
			if v.ExpectedErr != nil {
				assert.Error(t, err)
			}else{
				assert.NoError(t, err)
			}
			mock.AssertExpectations(t)
		})
	}
}
