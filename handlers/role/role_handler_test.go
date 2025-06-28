package role

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/models/response"
	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type roleHandlerTest struct {
	Name           string
	Data           []response.Roles
	Method         string
	Target         string
	Code           int
	Muxserver      *http.ServeMux
	Mocks          func(data []response.Roles, count int, err error) *mock.Call
	Request        *request.Roles
	RequestConvert func(data *request.Roles) []byte
	ErrorMocks     error
	CountData      int
	UseMock        bool
}

func TestHandlerGetAllData(t *testing.T) {
	Mockservice := &ServiceRole{Mock: mock.Mock{}}
	Mocklogger := &LoggerMock{}
	handler := RoleHandler(Mockservice, Mocklogger)

	data := []roleHandlerTest{
		{
			Name:       "Method Not Allowed",
			Data:       nil,
			Method:     helper.Post,
			Target:     "/role",
			Code:       helper.MethodNotAllowed,
			Muxserver:  http.NewServeMux(),
			Mocks:      nil,
			ErrorMocks: nil,
			CountData:  0,
			UseMock:    false,
		},
		{
			Name:      "Invalid Page Format",
			Data:      nil,
			Method:    helper.Gets,
			Target:    "/role?page=ababa",
			Code:      helper.BadRequest,
			Muxserver: http.NewServeMux(),
			Mocks:     nil,
			CountData: 0,
			UseMock:   false,
		},
		{
			Name:      "Internal Server Error",
			Data:      nil,
			Method:    helper.Gets,
			Target:    "/role",
			Code:      helper.InternalServError,
			Muxserver: http.NewServeMux(),
			Mocks: func(data []response.Roles, count int, err error) *mock.Call {
				return Mockservice.On("GetAllData", "", "asc", 1, 10).Return(data, count, err)
			},
			ErrorMocks: errors.New("Database is refused"),
			CountData:  0,
			UseMock:    true,
		},
		{
			Name: "Success Get All Data",
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
			Method:    helper.Gets,
			Target:    "/role",
			Code:      helper.Success,
			Muxserver: http.NewServeMux(),
			Mocks: func(data []response.Roles, count int, err error) *mock.Call {
				return Mockservice.On("GetAllData", "", "asc", 1, 10).Return(data, count, err)
			},
			ErrorMocks: nil,
			CountData:  2,
			UseMock:    true,
		},
	}

	for _, v := range data {
		t.Run(v.Name, func(t *testing.T) {

			Mockservice.ExpectedCalls = nil
			Mockservice.Calls = nil

			req := httptest.NewRequest(v.Method, v.Target, nil)
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			if v.UseMock {
				v.Mocks(v.Data, v.CountData, v.ErrorMocks)
			}

			v.Muxserver.HandleFunc("/role", handler.GetRole)
			v.Muxserver.ServeHTTP(w, req)

			assert.Equal(t, v.Code, w.Code)
			Mockservice.AssertExpectations(t)
		})
	}
}

func TestHandlerCreate(t *testing.T) {
	Mockservice := &ServiceRole{Mock: mock.Mock{}}
	Mocklogger := &LoggerMock{}
	handler := RoleHandler(Mockservice, Mocklogger)

	duplicate := &mysql.MySQLError{
		Number:  1062,
		Message: "Duplicate entry",
	}

	data := []roleHandlerTest{
		{
			Name: "Method Not Allowed",
			Request: &request.Roles{
				NameRole: "Admin",
			},
			Method: helper.Gets,
			Target: "/add/role",
			Code:   helper.MethodNotAllowed,
			RequestConvert: func(data *request.Roles) []byte {
				jsonmars, _ := json.Marshal(data)
				return jsonmars
			},
			Muxserver: http.NewServeMux(),
			UseMock:   false,
		},
		{
			Name:    "Request Body Missing",
			Request: nil,
			Method:  helper.Post,
			Target:  "/add/role",
			Code:    helper.BadRequest,
			RequestConvert: func(data *request.Roles) []byte {
				if data != nil {

					jsonmars, _ := json.Marshal(data)
					return jsonmars
				}
				return nil
			},
			Muxserver:  http.NewServeMux(),
			ErrorMocks: nil,
			UseMock:    false,
		},
		{
			Name:    "Validation Failed",
			Request: &request.Roles{},
			Method:  helper.Post,
			Target:  "/add/role",
			Code:    helper.UnprocessbleEntity,
			RequestConvert: func(data *request.Roles) []byte {
				if data != nil {

					jsonmars, _ := json.Marshal(data)
					return jsonmars
				}
				return nil
			},
			Muxserver:  http.NewServeMux(),
			ErrorMocks: nil,
			UseMock:    false,
		},
		{
			Name: "Failed Create Role (Duplicate Name Role)",
			Request: &request.Roles{
				NameRole: "Admin11",
			},
			Method: helper.Post,
			Target: "/add/role",
			Code:   helper.Conflict,
			RequestConvert: func(data *request.Roles) []byte {
				if data != nil {
					jsonmars, _ := json.Marshal(data)
					return jsonmars
				}
				return nil
			},
			Muxserver:  http.NewServeMux(),
			ErrorMocks: duplicate,
			UseMock:    true,
		},
		{
			Name: "Failed Create Role (Server Error)",
			Request: &request.Roles{
				NameRole: "Admin2",
			},
			Method: helper.Post,
			Target: "/add/role",
			Code:   helper.InternalServError,
			RequestConvert: func(data *request.Roles) []byte {
				if data != nil {
					jsonmars, _ := json.Marshal(data)
					return jsonmars
				}
				return nil
			},
			Muxserver:  http.NewServeMux(),
			ErrorMocks: errors.New("Database Is Refused"),
			UseMock:    true,
		},
		{
			Name: "Success Create Role",
			Request: &request.Roles{
				NameRole: "Admin",
			},
			Method: helper.Post,
			Target: "/add/role",
			Code:   helper.Created,
			RequestConvert: func(data *request.Roles) []byte {
				if data != nil {
					jsonmars, _ := json.Marshal(data)
					return jsonmars
				}
				return nil
			},
			Muxserver:  http.NewServeMux(),
			ErrorMocks: nil,
			UseMock:    true,
		},
	}

	for _, v := range data {
		t.Run(v.Name, func(t *testing.T) {
			body := bytes.NewReader([]byte{})
			if v.Request != nil {
				body = bytes.NewReader(v.RequestConvert(v.Request))
			}
			req := httptest.NewRequest(v.Method, v.Target, body)
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			v.Muxserver.HandleFunc("/add/role", handler.AddRoles)
			if v.UseMock {
				Mockservice.On("Create", v.Request).Return(v.ErrorMocks)
			}
			v.Muxserver.ServeHTTP(w, req)

			assert.Equal(t, v.Code, w.Code)
			Mockservice.AssertExpectations(t)
		})
	}
}

func TestHandlerShow(t *testing.T) {
	Mockservice := &ServiceRole{Mock: mock.Mock{}}
	Mocklogger := &LoggerMock{}
	handler := RoleHandler(Mockservice, Mocklogger)

	data := &response.Roles{
		IdRole:   1,
		NameRole: "Admin",
	}
	// success show data by id
	t.Run("Success show role by id", func(t *testing.T) {
		jsonmars, _ := json.Marshal(data)

		req := httptest.NewRequest(helper.Gets, "/role/1", bytes.NewReader(jsonmars))

		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		Mockservice.On("ShowRoleById", 1).Return(data, nil)
		mux := http.NewServeMux()
		mux.HandleFunc("/role/{id}", handler.Show)
		mux.ServeHTTP(w, req)

		assert.Equal(t, helper.Success, w.Code)
		Mockservice.AssertExpectations(t)
	})

	// failed show data by id
	t.Run("Failed show role by id", func(t *testing.T) {
		jsonmars, _ := json.Marshal(data)

		req := httptest.NewRequest(helper.Gets, "/role/2", bytes.NewReader(jsonmars))

		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		Mockservice.On("ShowRoleById", 2).Return(data, gorm.ErrRecordNotFound)
		mux := http.NewServeMux()
		mux.HandleFunc("/role/{id}", handler.Show)
		mux.ServeHTTP(w, req)

		assert.Equal(t, helper.Notfound, w.Code)
		Mockservice.AssertExpectations(t)
	})

	// failed show data (database refused)
	t.Run("Failed show role (database refused)", func(t *testing.T) {
		jsonmars, _ := json.Marshal(data)

		req := httptest.NewRequest(helper.Gets, "/role/2", bytes.NewReader(jsonmars))

		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		Mockservice.On("ShowRoleById", 2).Return(data, errors.New("Some Thong Wrong"))
		mux := http.NewServeMux()
		mux.HandleFunc("/role/{id}", handler.Show)
		mux.ServeHTTP(w, req)

		assert.Equal(t, helper.Notfound, w.Code)
		Mockservice.AssertExpectations(t)
	})

	// Method Not AllowedShow data
	t.Run("Method Not Allowed", func(t *testing.T) {
		handler := RoleHandler(Mockservice, Mocklogger)
		req := httptest.NewRequest(helper.Post, "/role/{id}", nil)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		mux := http.NewServeMux()
		mux.HandleFunc("/role/{id}", handler.Show)
		mux.ServeHTTP(w, req)

		assert.Equal(t, helper.MethodNotAllowed, w.Code)
	})

	// path value is not int
	t.Run("failed Path Id", func(t *testing.T) {
		jsonmars, _ := json.Marshal(data)
		req := httptest.NewRequest(helper.Gets, "/role/abc", bytes.NewReader(jsonmars))
		req.SetPathValue("id", "abc")
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		mux := http.NewServeMux()
		mux.HandleFunc("/role/{id}", handler.Show)
		mux.ServeHTTP(w, req)

		assert.Equal(t, helper.BadRequest, w.Code)
	})

}

func TestHandlerUpdate(t *testing.T) {
	mockservice := &ServiceRole{Mock: mock.Mock{}}
	Mocklogger := &LoggerMock{}
	handler := RoleHandler(mockservice, Mocklogger)

	reqdata := &request.Roles{
		NameRole: "AdminUpdate",
	}

	// Success Update Role
	t.Run("Success Update Role", func(t *testing.T) {
		jsonm, _ := json.Marshal(reqdata)
		req := httptest.NewRequest(helper.Put, "/role/1/update", bytes.NewReader(jsonm))
		w := httptest.NewRecorder()

		mockservice.On("UpdateRole", 1, reqdata).Return(nil)
		mux := http.NewServeMux()
		mux.HandleFunc("/role/{id}/update", handler.Update)
		mux.ServeHTTP(w, req)

		assert.Equal(t, helper.Success, w.Code)

		mockservice.AssertExpectations(t)
	})

	// failed update role id not found
	t.Run("Failed Update Role id not found", func(t *testing.T) {
		jsom, _ := json.Marshal(reqdata)
		req := httptest.NewRequest(helper.Put, "/role/90/update", bytes.NewReader(jsom))
		w := httptest.NewRecorder()

		mockservice.On("UpdateRole", 90, reqdata).Return(gorm.ErrRecordNotFound)
		mux := http.NewServeMux()
		mux.HandleFunc("/role/{id}/update", handler.Update)
		mux.ServeHTTP(w, req)

		assert.Equal(t, helper.Notfound, w.Code)

		mockservice.AssertExpectations(t)
	})

	// failed update name role exist
	t.Run("Failed Update Role name role is exist", func(t *testing.T) {
		jsom, _ := json.Marshal(reqdata)
		req := httptest.NewRequest(helper.Put, "/role/4/update", bytes.NewReader(jsom))
		w := httptest.NewRecorder()

		errDub := &mysql.MySQLError{
			Number:  1062,
			Message: "Duplicate entry",
		}

		mockservice.On("UpdateRole", 4, reqdata).Return(errDub)
		mux := http.NewServeMux()
		mux.HandleFunc("/role/{id}/update", handler.Update)
		mux.ServeHTTP(w, req)

		assert.Equal(t, helper.Conflict, w.Code)

		mockservice.AssertExpectations(t)
	})

	// failed update database refused
	t.Run("Failed Update Role database refused", func(t *testing.T) {
		jsom, _ := json.Marshal(reqdata)
		req := httptest.NewRequest(helper.Put, "/role/6/update", bytes.NewReader(jsom))
		w := httptest.NewRecorder()

		mockservice.On("UpdateRole", 6, reqdata).Return(errors.New("Something went wrong"))
		mux := http.NewServeMux()
		mux.HandleFunc("/role/{id}/update", handler.Update)
		mux.ServeHTTP(w, req)

		assert.Equal(t, helper.InternalServError, w.Code)

		mockservice.AssertExpectations(t)
	})

	// body nil
	t.Run("Request Body Nil", func(t *testing.T) {
		req := httptest.NewRequest(helper.Put, "/role/abc/update", nil)
		req.SetPathValue("id", "abc")
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		mux := http.NewServeMux()
		mux.HandleFunc("/role/{id}/update", handler.Update)
		mux.ServeHTTP(w, req)

		assert.Equal(t, helper.BadRequest, w.Code)
	})

	// Method Not AllowedShow data
	t.Run("Method Not Allowed", func(t *testing.T) {
		jsom, _ := json.Marshal(reqdata)
		req := httptest.NewRequest(helper.Gets, "/role/90/update", bytes.NewReader(jsom))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		mux := http.NewServeMux()
		mux.HandleFunc("/role/{id}/update", handler.Update)
		mux.ServeHTTP(w, req)

		assert.Equal(t, helper.MethodNotAllowed, w.Code)
	})

	// path value is not int
	t.Run("failed Path Id", func(t *testing.T) {
		jsonmars, _ := json.Marshal(reqdata)
		req := httptest.NewRequest(helper.Put, "/role/abc/update", bytes.NewReader(jsonmars))
		req.SetPathValue("id", "abc")
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		mux := http.NewServeMux()
		mux.HandleFunc("/role/{id}/update", handler.Update)
		mux.ServeHTTP(w, req)

		assert.Equal(t, helper.BadRequest, w.Code)
	})

}

func TestHandlerDelete(t *testing.T) {
	mockservice := &ServiceRole{Mock: mock.Mock{}}
	logger := &LoggerMock{}

	handler := RoleHandler(mockservice, logger)

	t.Run("Success Delete Role", func(t *testing.T) {
		req := httptest.NewRequest(helper.Delete, "/role/1/delete", nil)

		mockservice.On("DeleteRole", 1).Return(nil)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		mux := http.NewServeMux()
		mux.HandleFunc("/role/{id}/delete", handler.Delete)
		mux.ServeHTTP(w, req)

		assert.Equal(t, helper.Success, w.Code)
		mockservice.AssertExpectations(t)
	})

	t.Run("Failed Delete Role (not found id)", func(t *testing.T) {
		req := httptest.NewRequest(helper.Delete, "/role/9/delete", nil)
		w := httptest.NewRecorder()

		mockservice.On("DeleteRole", 9).Return(gorm.ErrRecordNotFound)
		mux := http.NewServeMux()
		mux.HandleFunc("/role/{id}/delete", handler.Delete)
		mux.ServeHTTP(w, req)

		assert.Equal(t, helper.Notfound, w.Code)
		mockservice.AssertExpectations(t)

	})

	t.Run("Failed Delete Role (database refused)", func(t *testing.T) {
		req := httptest.NewRequest(helper.Delete, "/role/90/delete", nil)
		w := httptest.NewRecorder()

		mockservice.On("DeleteRole", 90).Return(errors.New("something went wrong"))
		mux := http.NewServeMux()
		mux.HandleFunc("/role/{id}/delete", handler.Delete)
		mux.ServeHTTP(w, req)

		assert.Equal(t, helper.InternalServError, w.Code)
		mockservice.AssertExpectations(t)

	})

	t.Run("Failed Path Value Delete Role", func(t *testing.T) {
		req := httptest.NewRequest(helper.Delete, "/role/abc/delete", nil)

		req.Header.Set("Content-Type", "application/json")
		req.SetPathValue("id", "abc")
		w := httptest.NewRecorder()
		mux := http.NewServeMux()
		mux.HandleFunc("/role/{id}/delete", handler.Delete)
		mux.ServeHTTP(w, req)

		assert.Equal(t, helper.BadRequest, w.Code)
	})

	t.Run("Method Not Allowed", func(t *testing.T) {
		req := httptest.NewRequest(helper.Gets, "/role/90/delete", nil)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		mux := http.NewServeMux()
		mux.HandleFunc("/role/{id}/delete", handler.Delete)
		mux.ServeHTTP(w, req)

		assert.Equal(t, helper.MethodNotAllowed, w.Code)
	})
}
