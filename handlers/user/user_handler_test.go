package user

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/models/response"
	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

type TestHandlerUser struct {
	Name 			string
	Method 			string
	Target 			string
	ResponseAll 	[]response.Users
	RequestUser 	*request.User
	Search,Sort		string
	Page,Perpage,Code,Count	int
	Mux 			*http.ServeMux
	Id 				any
	Response 		*response.Users
	ExpectedErr		error
	UseMock 		bool	
}


func TestGetAll(t *testing.T)  {
	mockuser := MockUserServices{}
	mocklogger := LoggerMock{}
	handleruser := NewHandlerUser(&mockuser,mocklogger)
	data:=[]TestHandlerUser{
		{
			Name: "Method Not Allowed",
			Method: helper.Post,
			Target: "/user",
			Code: helper.MethodNotAllowed,
			ResponseAll: nil,
			Mux: http.NewServeMux(),
			ExpectedErr: nil,
			UseMock: false,
		},
		{
			Name: "Invalid Query Format Params",
			Method: helper.Gets,
			Target: "/user?page=abc",
			Code: helper.BadRequest,
			ResponseAll: nil,
			Mux: http.NewServeMux(),
			ExpectedErr: nil,
			UseMock: false,
		},
		{
			Name: "Internal Server Error",
			Method: helper.Gets,
			Target: "/user",
			Code: helper.InternalServError,
			ResponseAll: nil,
			Search: "",
			Sort: "asc",
			Page: 1,
			Perpage: 10,
			Count: 0,
			Mux: http.NewServeMux(),
			ExpectedErr: errors.New("Database Is Refused"),
			UseMock: true,
		},
		{
			Name: "Success",
			Method: helper.Gets,
			Target: "/user",
			Code: helper.InternalServError,
			ResponseAll: nil,
			Search: "",
			Sort: "asc",
			Page: 1,
			Perpage: 10,
			Count: 0,
			Mux: http.NewServeMux(),
			ExpectedErr: nil,
			UseMock: true,
		},
	}

	for _, v := range data {
		t.Run(v.Name, func(t *testing.T) {
			req := httptest.NewRequest(v.Method,v.Target, nil)
			req.Header.Set("Content-Type","application/json")
			w := httptest.NewRecorder()
			if v.UseMock {
				mockuser.On("AllUser",v.Search,v.Sort,v.Page,v.Perpage).Return(v.ResponseAll,v.Count,v.ExpectedErr)
			}

			v.Mux.HandleFunc("/user", handleruser.GetAllUser)
			v.Mux.ServeHTTP(w, req)

			assert.Equal(t, v.Code, w.Code)
			mockuser.AssertExpectations(t)

		})
	}
}

func TestCreate(t *testing.T)  {
	mockuser := MockUserServices{}
	mocklogger := LoggerMock{}
	handleruser := NewHandlerUser(&mockuser, mocklogger)
	
	data := []TestHandlerUser{
		{
			Name: "Method Not Allowed",
			RequestUser: nil,
			Method: helper.Gets,
			Target: "/user/add",
			Code: helper.MethodNotAllowed,
			Mux: http.NewServeMux(),
			ExpectedErr: nil,
			UseMock: false,
		},
		{
			Name: "Body Request Is Missing",
			RequestUser: nil,
			Method: helper.Post,
			Target: "/user/add",
			Code: helper.BadRequest,
			Mux: http.NewServeMux(),
			ExpectedErr: nil,
			UseMock: false,
		},
		{
			Name: "Invalid Validation",
			RequestUser: &request.User{},
			Method: helper.Post,
			Target: "/user/add",
			Code: helper.UnprocessbleEntity,
			Mux: http.NewServeMux(),
			ExpectedErr: nil,
			UseMock: false,
		},
		{
			Name: "Failed: Email Duplicate",
			RequestUser: &request.User{
				Email: "admindup@gmail.com",
				Password: "admindup123",
				RoleId: 1,
			},
			Method: helper.Post,
			Target: "/user/add",
			Code: helper.Conflict,
			Mux: http.NewServeMux(),
			ExpectedErr: &mysql.MySQLError{
				Message: "Duplicate Entry",
				Number: 1062,
			},
			UseMock: true,
		},
		{
			Name: "Failed: Internal Server Error",
			RequestUser: &request.User{
				Email: "adminerr@gmail.com",
				Password: "adminerr123",
				RoleId: 1,
			},
			Method: helper.Post,
			Target: "/user/add",
			Code: helper.InternalServError,
			Mux: http.NewServeMux(),
			ExpectedErr: errors.New("Database Is Refused"),
			UseMock: true,
		},
		{
			Name: "Success Create User",
			RequestUser: &request.User{
				Email: "admin@gmail.com",
				Password: "admin123",
				RoleId: 1,
			},
			Method: helper.Post,
			Target: "/user/add",
			Code: helper.Created,
			Mux: http.NewServeMux(),
			ExpectedErr: nil,
			UseMock: true,
		},
	}

	for _, v := range data {
		t.Run(v.Name,func(t *testing.T) {
			var userreq []byte
			if v.RequestUser != nil {
				var err error
				userreq, err = json.Marshal(v.RequestUser)
				assert.NoError(t, err)
			}

			req := httptest.NewRequest(v.Method,v.Target, bytes.NewReader(userreq))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			if v.UseMock {
				mockuser.On("CreateUser", v.RequestUser).Return(v.ExpectedErr)
			}

			v.Mux.HandleFunc("/user/add", handleruser.Create)
			v.Mux.ServeHTTP(w, req)

			assert.Equal(t, v.Code, w.Code)
			mockuser.AssertExpectations(t)
		})
	}
}

func TestShow(t *testing.T)  {
	mockuser := MockUserServices{}
	mocklogger := LoggerMock{}
	handleruser := NewHandlerUser(&mockuser,mocklogger)

	data := []TestHandlerUser{
		{
			Name: "Method Not Allowed",
			Method: helper.Post,
			Target: "/user/18",
			Code: helper.MethodNotAllowed,
			Mux: http.NewServeMux(),
			ExpectedErr: nil,
			UseMock: false,
		},
		{
			Name: "Invalid Id Format",
			Method: helper.Gets,
			Target: "/user/abc",
			Code: helper.BadRequest,
			Mux: http.NewServeMux(),
			ExpectedErr: nil,
			UseMock: false,
		},
		{
			Name: "Failed: Not Found Data",
			Method: helper.Gets,
			Id: 13,
			Target: "/user/13",
			Code: helper.Notfound,
			Mux: http.NewServeMux(),
			ExpectedErr: gorm.ErrRecordNotFound,
			UseMock: true,
		},
		{
			Name: "Failed: Internal Server Error",
			Method: helper.Gets,
			Id: 14,
			Target: "/user/14",
			Code: helper.InternalServError,
			Mux: http.NewServeMux(),
			ExpectedErr: errors.New("Database Is Refused"),
			UseMock: true,
		},
		{
			Name: "success",
			Method: helper.Gets,
			Id: 1,
			Target: "/user/1",
			Code: helper.Success,
			Mux: http.NewServeMux(),
			ExpectedErr: nil,
			UseMock: true,
		},
	}
	for _, v := range data {
		t.Run(v.Name,func(t *testing.T) {
			req := httptest.NewRequest(v.Method,v.Target,nil)
			req.Header.Set("Content-Type","application/json")
			w := httptest.NewRecorder()

			if v.UseMock {
				mockuser.On("ShowUser",v.Id).Return(v.Response,v.ExpectedErr)
			}
			v.Mux.HandleFunc("/user/{id}",handleruser.Show)
			v.Mux.ServeHTTP(w, req)

			assert.Equal(t, v.Code, w.Code)
			mockuser.AssertExpectations(t)
		})
	}
}

func TestUpdate(t *testing.T)  {
	mockuser := MockUserServices{}
	mocklogger := LoggerMock{}
	handleruser := NewHandlerUser(&mockuser,mocklogger)
	data := []TestHandlerUser{
		{
			Name: "Method Not Allowed",
			Method: helper.Gets,
			Target: "/user/13/update",
			Code: helper.MethodNotAllowed,
			Mux: http.NewServeMux(),
			ExpectedErr: nil,
			UseMock: false,
		},
		{
			Name: "Body Request Is Missing",
			RequestUser: nil,
			Method: helper.Put,
			Target: "/user/14/update",
			Code: helper.BadRequest,
			Mux: http.NewServeMux(),
			ExpectedErr: nil,
			UseMock: false,
		},
		{
			Name: "Invalid Id Format",
			RequestUser: nil,
			Method: helper.Put,
			Target: "/user/abc/update",
			Code: helper.BadRequest,
			Mux: http.NewServeMux(),
			ExpectedErr: nil,
			UseMock: false,
		},
		{
			Name: "Failed: Not Found Data",
			RequestUser: &request.User{
				Email: "Adminnfound@gmail.com",
			},
			Id: 2,
			Method: helper.Put,
			Target: "/user/2/update",
			Code: helper.Notfound,
			Mux: http.NewServeMux(),
			ExpectedErr: gorm.ErrRecordNotFound,
			UseMock: true,
		},
		{
			Name: "Failed: Duplicate Data",
			RequestUser: &request.User{
				Email: "Admindup@gmail.com",
			},
			Id: 3,
			Method: helper.Put,
			Target: "/user/3/update",
			Code: helper.Conflict,
			Mux: http.NewServeMux(),
			ExpectedErr: &mysql.MySQLError{
				Message: "Duplicate Entry",
				Number: 1062, //error untuk menyatakn bahwa data nya duplicate
			},
			UseMock: true,
		},
		{
			Name: "Failed: Internal Server Error",
			RequestUser: &request.User{
				Email: "Adminerr@gmail.com",
			},
			Id: 9,
			Method: helper.Put,
			Target: "/user/9/update",
			Code: helper.InternalServError,
			Mux: http.NewServeMux(),
			ExpectedErr: errors.New("Database Is Refused"),
			UseMock: true,
		},
		{
			Name: "Success",
			RequestUser: &request.User{
				Email: "Admin@gmail.com",
			},
			Id: 1,
			Method: helper.Put,
			Target: "/user/1/update",
			Code: helper.Success,
			Mux: http.NewServeMux(),
			ExpectedErr: nil,
			UseMock: true,
		},
	}

	for _, v := range data {
		t.Run(v.Name,func(t *testing.T) {
			var usereq []byte
			if v.RequestUser != nil {
				var err error
				usereq, err = json.Marshal(v.RequestUser)
				assert.NoError(t, err)
			}

			req := httptest.NewRequest(v.Method,v.Target, bytes.NewReader(usereq))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			if v.UseMock {
				mockuser.On("UpdateUser", v.Id,v.RequestUser).Return(v.ExpectedErr)
			}
			v.Mux.HandleFunc("/user/{id}/update", handleruser.Update)
			v.Mux.ServeHTTP(w, req)

			assert.Equal(t, v.Code, w.Code)
			mockuser.AssertExpectations(t)
		})
	}
}

func TestDelete(t *testing.T)  {
	mockuser := MockUserServices{}
	mocklogger:= LoggerMock{}
	handleruser := NewHandlerUser(&mockuser, mocklogger)
	data := []TestHandlerUser{
		{
			Name: "Method Not Allowed",
			Method: helper.Gets,
			Target: "/user/90/delete",
			Code: helper.MethodNotAllowed,
			Mux: http.NewServeMux(),
			ExpectedErr: nil,
			UseMock: false,
		},
		{
			Name: "Invalid Id Format",
			Method: helper.Delete,
			Target: "/user/abc/delete",
			Code: helper.BadRequest,
			Mux: http.NewServeMux(),
			ExpectedErr: nil,
			UseMock: false,
		},
		{
			Name: "Failed: Not Found Data",
			Method: helper.Delete,
			Target: "/user/10/delete",
			Id: 10,
			Code: helper.Notfound,
			Mux: http.NewServeMux(),
			ExpectedErr: gorm.ErrRecordNotFound,
			UseMock: true,
		},
		{
			Name: "Failed: Internal Server Erro",
			Method: helper.Delete,
			Target: "/user/5/delete",
			Id: 5,
			Code: helper.InternalServError,
			Mux: http.NewServeMux(),
			ExpectedErr: errors.New("Database Is Refused"),
			UseMock: true,
		},
		{
			Name: "Success",
			Method: helper.Delete,
			Target: "/user/1/delete",
			Id: 1,
			Code: helper.Success,
			Mux: http.NewServeMux(),
			ExpectedErr: nil,
			UseMock: true,
		},
	}

	for _, v := range data {
		t.Run(v.Name,func(t *testing.T) {
			req := httptest.NewRequest(v.Method, v.Target, nil)
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			if v.UseMock {
				mockuser.On("DeleteUser", v.Id).Return(v.ExpectedErr)
			}
			v.Mux.HandleFunc("/user/{id}/delete", handleruser.Delete)
			v.Mux.ServeHTTP(w, req)

			assert.Equal(t, v.Code, w.Code)
			mockuser.AssertExpectations(t)
		})
	}
}