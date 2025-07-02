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
	DataShow       *response.Roles
	Id            any
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
	mockservice := &ServiceRole{}
	mocklogger := &LoggerMock{}
	handler := RoleHandler(mockservice, mocklogger)

	data := []roleHandlerTest{
		{
			Name:       "Method Not Allowed",
			DataShow:   nil,
			Id:         5,
			Method:     helper.Post,
			Target:     "/role/1",
			Code:       helper.MethodNotAllowed,
			Muxserver:  http.NewServeMux(),
			ErrorMocks: nil,
			UseMock:    false,
		},
		{
			Name:       "Invalid Format Param",
			DataShow:   nil,
			Id:         5,
			Method:     helper.Gets,
			Target:     "/role/abc",
			Code:       helper.BadRequest,
			Muxserver:  http.NewServeMux(),
			ErrorMocks: nil,
			UseMock:    false,
		},
		{
			Name:       "Not Found Data By Id",
			DataShow:   nil,
			Id:         11,
			Method:     helper.Gets,
			Target:     "/role/11",
			Code:       helper.Notfound,
			Muxserver:  http.NewServeMux(),
			ErrorMocks: gorm.ErrRecordNotFound,
			UseMock:    true,
		},
		{
			Name:       "Internal Server Error",
			DataShow:   nil,
			Id:         12,
			Method:     helper.Gets,
			Target:     "/role/12",
			Code:       helper.InternalServError,
			Muxserver:  http.NewServeMux(),
			ErrorMocks: errors.New("Database Is Refused"),
			UseMock:    true,
		},
		{
			Name: "Success Show role by id",
			DataShow: &response.Roles{
				IdRole:   1,
				NameRole: "Admin",
			},
			Id:         1,
			Method:     helper.Gets,
			Target:     "/role/1",
			Code:       helper.Success,
			Muxserver:  http.NewServeMux(),
			ErrorMocks: nil,
			UseMock:    true,
		},
	}

	for _, v := range data {
		t.Run(v.Name, func(t *testing.T) {
			req := httptest.NewRequest(v.Method, v.Target, nil)
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			if v.UseMock {
				mockservice.On("ShowRoleById", v.Id).Return(v.DataShow, v.ErrorMocks)
			}
			v.Muxserver.HandleFunc("/role/{id}", handler.Show)
			v.Muxserver.ServeHTTP(w, req)

			assert.Equal(t, v.Code, w.Code)
			mockservice.AssertExpectations(t)

		})
	}

}

func TestHandlerUpdate(t *testing.T) {
	mockservice := &ServiceRole{}
	mocklogger := &LoggerMock{}
	handler := RoleHandler(mockservice, mocklogger)

	data := []roleHandlerTest{
		{
			Name:           "Method Not Allowed",
			Id:             2,
			Request:        nil,
			Method:         helper.Gets,
			Target:         "/role/2/update",
			Code:           helper.MethodNotAllowed,
			RequestConvert: nil,
			Muxserver:      http.NewServeMux(),
			ErrorMocks:     nil,
			UseMock:        false,
		},
		{
			Name: "Body Request Is missing",
			Id:   2,
			Request: nil,
			Method:         helper.Put,
			Target:         "/role/2/update",
			Code:           helper.BadRequest,
			RequestConvert: nil,
			Muxserver:      http.NewServeMux(),
			ErrorMocks:     nil,
			UseMock:        false,
		},
		{
			Name: "Invalid Format Param",
			Id:   "abc",
			Request: nil,
			Method:         helper.Put,
			Target:         "/role/abc/update",
			Code:           helper.BadRequest,
			RequestConvert: nil,
			Muxserver:      http.NewServeMux(),
			ErrorMocks:     nil,
			UseMock:        false,
		},
		{
			Name: "Failed Update Role Not Found Data By Id",
			Id:   5,
			Request: &request.Roles{
				NameRole: "Admin5",
			},
			Method: helper.Put,
			Target: "/role/5/update",
			Code:   helper.Notfound,
			RequestConvert: func(data *request.Roles) []byte {
				jsonm, _ := json.Marshal(data)
				return jsonm
			},
			Muxserver:  http.NewServeMux(),
			ErrorMocks: gorm.ErrRecordNotFound,
			UseMock:    true,
		},
		{
			Name: "Server Error",
			Id:   5,
			Request: &request.Roles{
				NameRole: "Admin5",
			},
			Method: helper.Put,
			Target: "/role/5/update",
			Code:   helper.Notfound,
			RequestConvert: func(data *request.Roles) []byte {
				jsonm, _ := json.Marshal(data)
				return jsonm
			},
			Muxserver:  http.NewServeMux(),
			ErrorMocks: errors.New("Database Refused"),
			UseMock:    true,
		},
		{
			Name: "Success Update Role",
			Id:   1,
			Request: &request.Roles{
				NameRole: "Admin",
			},
			Method: helper.Put,
			Target: "/role/1/update",
			Code:   helper.Success,
			RequestConvert: func(data *request.Roles) []byte {
				jsonm, _ := json.Marshal(data)
				return jsonm
			},
			Muxserver:  http.NewServeMux(),
			ErrorMocks: nil,
			UseMock:    true,
		},
	}

	for _, v := range data {
		t.Run(v.Name, func(t *testing.T) {
			bodyreq := bytes.NewReader([]byte{})
			if v.Request != nil {
				bodyreq = bytes.NewReader(v.RequestConvert(v.Request))
			}
			req := httptest.NewRequest(v.Method, v.Target, bodyreq)
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			if v.UseMock {
				mockservice.On("UpdateRole", v.Id, v.Request).Return(v.ErrorMocks)
			}

			v.Muxserver.HandleFunc("/role/{id}/update", handler.Update)
			v.Muxserver.ServeHTTP(w, req)

			assert.Equal(t, v.Code, w.Code)
			mockservice.AssertExpectations(t)
		})
	}
}

func TestHandlerDelete(t *testing.T)  {
	mockservice := &ServiceRole{}
	mocklogger := &LoggerMock{}
	handler := RoleHandler(mockservice, mocklogger)

	data := []roleHandlerTest{
		{
			Name: "Method Not Found",
			Id: 10,
			Method: helper.Gets,
			Target: "/role/10/delete",
			Code: helper.MethodNotAllowed,
			Muxserver: http.NewServeMux(),
			ErrorMocks: nil,
			UseMock: false,
		},
		{
			Name: "Invalid Format Param",
			Id: "abc",
			Method: helper.Delete,
			Target: "/role/abc/delete",
			Code: helper.BadRequest,
			Muxserver: http.NewServeMux(),
			ErrorMocks: nil,
			UseMock: false,
		},
		{
			Name: "Invalid Format Param",
			Id: "abc",
			Method: helper.Delete,
			Target: "/role/abc/delete",
			Code: helper.BadRequest,
			Muxserver: http.NewServeMux(),
			ErrorMocks: nil,
			UseMock: false,
		},
		{
			Name: "Not Found Data",
			Id: 100,
			Method: helper.Delete,
			Target: "/role/100/delete",
			Code: helper.Notfound,
			Muxserver: http.NewServeMux(),
			ErrorMocks: gorm.ErrRecordNotFound,
			UseMock: true,
		},
		{
			Name: "Server Error",
			Id: 11,
			Method: helper.Delete,
			Target: "/role/11/delete",
			Code: helper.InternalServError,
			Muxserver: http.NewServeMux(),
			ErrorMocks: errors.New("Database Is Refused"),
			UseMock: true,
		},
		{
			Name: "Success Delete",
			Id: 1,
			Method: helper.Delete,
			Target: "/role/1/delete",
			Code: helper.Success,
			Muxserver: http.NewServeMux(),
			ErrorMocks: nil,
			UseMock: true,
		},
	}

	for _, v := range data {
		t.Run(v.Name, func(t *testing.T) {
			req := httptest.NewRequest(v.Method, v.Target, nil)
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			if v.UseMock {
				mockservice.On("DeleteRole", v.Id).Return(v.ErrorMocks)
			}

			v.Muxserver.HandleFunc("/role/{id}/delete", handler.Delete)
			v.Muxserver.ServeHTTP(w, req)

			assert.Equal(t, v.Code, w.Code)
			mockservice.AssertExpectations(t)
		})
	}
}
