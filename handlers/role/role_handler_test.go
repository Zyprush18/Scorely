package role

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/models/response"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// var Mockservice = &ServiceRole{Mock: mock.Mock{}}
// var Mocklogger = &LoggerMock{}


func TestHandlerGetAllData(t *testing.T)  {
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
	t.Run("Success Get All Data", func(t *testing.T) {
		// di buat di sini biar nggak jadi pararel/menimpa return value subtest di bwah
		Mockservice := &ServiceRole{Mock: mock.Mock{}}
		Mocklogger := &LoggerMock{}
		handler := RoleHandler(Mockservice,Mocklogger)
		req:= httptest.NewRequest(helper.Gets, "/role", nil)
		req.Header.Set("Content-Type","application/json")

		w := httptest.NewRecorder()

		Mockservice.Mock.On("GetAllData").Return(data, nil)

		handler.GetRole(w,req)

		assert.Equal(t, helper.Success, w.Code)

		Mockservice.Mock.AssertExpectations(t)
	})

	t.Run("Failed Get All Data", func(t *testing.T) {
		Mockservice := &ServiceRole{Mock: mock.Mock{}}
		Mocklogger := &LoggerMock{}
		handler := RoleHandler(Mockservice,Mocklogger)
		req:= httptest.NewRequest(helper.Gets, "/role", nil)
		req.Header.Set("Content-Type","application/json")

		w := httptest.NewRecorder()

		Mockservice.Mock.On("GetAllData").Return([]response.Roles(nil), errors.New("Database is refused"))

		handler.GetRole(w,req)

		assert.Equal(t, helper.BadRequest, w.Code)

		Mockservice.Mock.AssertExpectations(t)
	})
}

func TestHandlerCreate(t *testing.T) {
	Mockservice := &ServiceRole{Mock: mock.Mock{}}
	Mocklogger := &LoggerMock{}
	handler := RoleHandler(Mockservice, Mocklogger)

	t.Run("Succes Create a New Role", func(t *testing.T) {
		body := &request.Roles{
			NameRole: "Admin",
		}

		jsonBody, _ := json.Marshal(body)

		req := httptest.NewRequest(helper.Post, "/add/role", bytes.NewReader(jsonBody))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		Mockservice.On("Create", body).Return(nil)
		handler.AddRoles(w, req)

		assert.Equal(t, helper.Created, w.Code)
		Mockservice.AssertExpectations(t)
	})

	t.Run("Failed Create a New Role", func(t *testing.T) {
		bodyFail := &request.Roles{
			NameRole: "",
		}
		jsonBody, _ := json.Marshal(bodyFail)
		req := httptest.NewRequest(helper.Post, "/add/role", bytes.NewReader(jsonBody))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		Mockservice.On("Create", mock.Anything).Return(errors.New("Failed Create a New Role"))
		handler.AddRoles(w, req)

		assert.Equal(t, helper.BadRequest, w.Code)
		Mockservice.AssertExpectations(t)
	})
}


func TestHandlerShow(t *testing.T)  {
	Mockservice := &ServiceRole{Mock: mock.Mock{}}
	Mocklogger := &LoggerMock{}
	handler := RoleHandler(Mockservice, Mocklogger)

	data:= &response.Roles{
			IdRole: 1,
			NameRole: "Admin",
		}
	t.Run("Success show role by id", func(t *testing.T) {
		jsonmars, _ := json.Marshal(data)

		req := httptest.NewRequest(helper.Gets, "/role/1", bytes.NewReader(jsonmars))

		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		Mockservice.On("ShowRoleById", 1).Return(data,nil)
		handler.Show(w, req)


		assert.Equal(t, helper.Success, w.Code)
		Mockservice.AssertExpectations(t)
	})

	t.Run("Failed show role by id", func(t *testing.T) {
		jsonmars, _ := json.Marshal(data)

		req := httptest.NewRequest(helper.Gets, "/role/2", bytes.NewReader(jsonmars))

		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		Mockservice.On("ShowRoleById", 2).Return(data,errors.New("Not Found Role id 2"))
		handler.Show(w, req)


		assert.Equal(t, helper.Notfound, w.Code)
		Mockservice.AssertExpectations(t)
	})
}