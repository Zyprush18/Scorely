package userservice

import (
	"errors"
	"testing"
	"time"

	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/models/response"
	"github.com/stretchr/testify/assert"
)

type TestUser struct {
	Name          string
	DataRespAll   []response.Users
	RequestUser   *request.User
	Page, Perpage int
	Search, Sort  string
	Counts        int
	Response      *response.Users
	Id            int
	ExpectedErr   error
	IsError       bool
}

func TestGetAll(t *testing.T) {
	mockuser := UserRepository{}
	serviceuser := NewUserService(&mockuser)
	data := []TestUser{
		{
			Name: "Success Get All User",
			DataRespAll: []response.Users{
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
			},
			Counts:      2,
			Search:      "",
			Sort:        "ASC",
			Page:        1,
			Perpage:     10,
			ExpectedErr: nil,
			IsError:     false,
		},
		{
			Name:        "Failed Get All User",
			DataRespAll: nil,
			Counts:      0,
			Search:      "",
			Sort:        "ASC",
			Page:        1,
			Perpage:     10,
			ExpectedErr: errors.New("Database Is Refused"),
			IsError:     true,
		},
	}

	for _, v := range data {
		t.Run(v.Name, func(t *testing.T) {
			// hapus cache testing sebelumnya
			mockuser.ExpectedCalls = nil
			mockuser.Calls = nil

			mockuser.On("GetAll", v.Search, v.Sort, v.Page, v.Perpage).Return(v.DataRespAll, v.Counts, v.ExpectedErr)
			resp, count, err := serviceuser.AllUser(v.Search, v.Sort, v.Page, v.Perpage)
			if v.IsError {
				assert.Error(t, err)
				assert.Equal(t, int64(v.Counts), count)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, int64(v.Counts), count)
				assert.NotNil(t, resp)
			}
			mockuser.AssertExpectations(t)
		})
	}
}

func TestCreateUser(t *testing.T) {
	mockuser := UserRepository{}
	serviceuser := NewUserService(&mockuser)
	data := []TestUser{
		{
			Name: "Success Create User",
			RequestUser: &request.User{
				Email:    "Admin@gmail.com",
				Password: "admin123",
				RoleId:   1,
			},
			ExpectedErr: nil,
			IsError:     false,
		},
		{
			Name: "Success Create User",
			RequestUser: &request.User{
				Email:    "Admin@gmail.com",
				Password: "admin123",
				RoleId:   2,
			},
			ExpectedErr: errors.New("Database Is Refused"),
			IsError:     true,
		},
	}

	for _, v := range data {
		t.Run(v.Name, func(t *testing.T) {
			mockuser.On("Create", v.RequestUser).Return(v.ExpectedErr)
			err := serviceuser.CreateUser(v.RequestUser)
			if v.IsError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			mockuser.AssertExpectations(t)
		})
	}
}

func TestShowUser(t *testing.T) {
	mockuser := UserRepository{}
	serviceuser := NewUserService(&mockuser)
	data := []TestUser{
		{
			Name: "Success Show User",
			Response: &response.Users{
				IdUser:   1,
				Email:    "Admin@gmail.com",
				Password: "admin123",
				RoleId:   1,
				Models: helper.Models{
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
			},
			Id:          1,
			ExpectedErr: nil,
			IsError:     false,
		},
		{
			Name: "Success Show User",
			Response: &response.Users{
				IdUser:   2,
				Email:    "User@gmail.com",
				Password: "user123",
				RoleId:   2,
				Models: helper.Models{
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
			},
			Id:          4,
			ExpectedErr: errors.New("Not Found Data"),
			IsError:     true,
		},
	}

	for _, v := range data {
		t.Run(v.Name, func(t *testing.T) {
			mockuser.On("Show", v.Id).Return(v.Response, v.ExpectedErr)
			resp, err := serviceuser.ShowUser(v.Id)
			if v.IsError {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, uint(v.Id), resp.IdUser)
				assert.Equal(t,  "Admin@gmail.com", resp.Email)
			}
			mockuser.AssertExpectations(t)
		})
	}
}

func TestUpdateUser(t *testing.T)  {
	mockuser := UserRepository{}
	seriveuser := NewUserService(&mockuser)
	data := []TestUser{
		{
			Name: "Success Update User",
			RequestUser: &request.User{
				Email: "Admin@gmail.com",
			},
			Id: 1,
			ExpectedErr: nil,
			IsError: false,
		},
		{
			Name: "Failed Update User",
			RequestUser: &request.User{
				Email: "AdminUpdate@gmail.com",
			},
			Id: 2,
			ExpectedErr: errors.New("Not Found Data"),
			IsError: true,
		},
	}

	for _, v := range data {
		t.Run(v.Name, func(t *testing.T) {
			mockuser.On("Update", v.Id,v.RequestUser).Return(v.ExpectedErr)
			err := seriveuser.UpdateUser(v.Id, v.RequestUser)
			if v.IsError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			mockuser.AssertExpectations(t)
		})
	}
}

func TestDeleteUser(t *testing.T)  {
	mockuser := UserRepository{}
	serviceuser := NewUserService(&mockuser)
	data := []TestUser{
		{
			Name: "Success Delete User",
			Id: 1,
			ExpectedErr: nil,
			IsError: false,
		},
		{
			Name: "Failed Delete User",
			Id: 2,
			ExpectedErr: errors.New("Not Found Data"),
			IsError: true,
		},
	}

	for _, v := range data {
		t.Run(v.Name, func(t *testing.T) {
			mockuser.On("Delete", v.Id).Return(v.ExpectedErr)
			err := serviceuser.DeleteUser(v.Id)
			if v.IsError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			mockuser.AssertExpectations(t)
		})
	}
}