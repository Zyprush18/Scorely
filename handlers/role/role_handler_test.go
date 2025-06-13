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

		// succes get all data
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

		// Failed get all data
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

	// Method Not Allowed get all data
	t.Run("Method Not Allowed", func(t *testing.T) {
		Mockservice := &ServiceRole{Mock: mock.Mock{}}
		Mocklogger := &LoggerMock{}
		handler := RoleHandler(Mockservice,Mocklogger)
		req:= httptest.NewRequest(helper.Post, "/role", nil)
		req.Header.Set("Content-Type","application/json")

		w := httptest.NewRecorder()

		handler.GetRole(w,req)

		assert.Equal(t, helper.MethodNotAllowed, w.Code)
	})
}

func TestHandlerCreate(t *testing.T) {
	Mockservice := &ServiceRole{Mock: mock.Mock{}}
	Mocklogger := &LoggerMock{}
	handler := RoleHandler(Mockservice, Mocklogger)

	// Success Create New Role
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

	// failed Create new role
	t.Run("Failed Create a New Role", func(t *testing.T) {
		bodyFail := &request.Roles{
			NameRole: "AdminUpdate",
		}
		jsonBody, _ := json.Marshal(bodyFail)
		req := httptest.NewRequest(helper.Post, "/add/role", bytes.NewReader(jsonBody))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		Mockservice.On("Create", bodyFail).Return(errors.New("Failed Create a New Role"))
		handler.AddRoles(w, req)

		assert.Equal(t, helper.BadRequest, w.Code)
		Mockservice.AssertExpectations(t)
	})

	// body nil
	t.Run("Request Body Nil", func(t *testing.T) {
		req := httptest.NewRequest(helper.Post, "/add/role", nil)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		handler.AddRoles(w, req)

		assert.Equal(t, helper.BadRequest, w.Code)
	})


	// validation error create new role
	t.Run("Validation Create a New Role (empety name_role)", func(t *testing.T) {
		bodyFail := &request.Roles{
			NameRole: "",
		}
		jsonBody, _ := json.Marshal(bodyFail)
		req := httptest.NewRequest(helper.Post, "/add/role", bytes.NewReader(jsonBody))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		handler.AddRoles(w, req)

		assert.Equal(t, helper.BadRequest, w.Code)
	})

	// method not Allowed Create a new role
	t.Run("Method Not Allowed", func(t *testing.T) {
		bodyFail := &request.Roles{
			NameRole: "AdminUpdate",
		}
		jsonBody, _ := json.Marshal(bodyFail)
		handler := RoleHandler(Mockservice,Mocklogger)
		req:= httptest.NewRequest(helper.Gets, "/add/role", bytes.NewReader(jsonBody))
		req.Header.Set("Content-Type","application/json")

		w := httptest.NewRecorder()

		handler.AddRoles(w,req)

		assert.Equal(t, helper.MethodNotAllowed, w.Code)
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
		// success show data by id
	t.Run("Success show role by id", func(t *testing.T) {
		jsonmars, _ := json.Marshal(data)

		req := httptest.NewRequest(helper.Gets, "/role/1", bytes.NewReader(jsonmars))

		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		Mockservice.On("ShowRoleById", 1).Return(data,nil)
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

		Mockservice.On("ShowRoleById", 2).Return(data,errors.New("Not Found Role id 2"))
		mux := http.NewServeMux()
		mux.HandleFunc("/role/{id}", handler.Show)
		mux.ServeHTTP(w, req)


		assert.Equal(t, helper.Notfound, w.Code)
		Mockservice.AssertExpectations(t)
	})

	// Method Not AllowedShow data
	t.Run("Method Not Allowed", func(t *testing.T) {
		handler := RoleHandler(Mockservice,Mocklogger)
		req:= httptest.NewRequest(helper.Post, "/role/{id}", nil)
		req.Header.Set("Content-Type","application/json")

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


		assert.Equal(t, helper.InternalServError, w.Code)
	})

}

func TestHandlerUpdate(t *testing.T)  {
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

	// failed update role
	t.Run("Failed Update Role", func(t *testing.T) {
		jsom, _ := json.Marshal(reqdata)
		req := httptest.NewRequest(helper.Put, "/role/90/update", bytes.NewReader(jsom))
		w := httptest.NewRecorder() 

		mockservice.On("UpdateRole", 90, reqdata).Return(errors.New("Not Found ID: 90"))
		mux := http.NewServeMux()
		mux.HandleFunc("/role/{id}/update", handler.Update)
		mux.ServeHTTP(w, req)

		assert.Equal(t, helper.BadRequest, w.Code)

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
		req.Header.Set("Content-Type","application/json")
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


		assert.Equal(t, helper.InternalServError, w.Code)
	})

}

func TestHandlerDelete(t *testing.T)  {
	mockservice := &ServiceRole{Mock: mock.Mock{}}
	logger := &LoggerMock{}

	handler := RoleHandler(mockservice, logger)

	t.Run("Success Delete Role", func(t *testing.T) {
		req := httptest.NewRequest(helper.Delete, "/role/1/delete", nil)

		mockservice.On("DeleteRole", 1).Return(nil)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		mux:= http.NewServeMux()
		mux.HandleFunc("/role/{id}/delete", handler.Delete)
		mux.ServeHTTP(w, req)

		assert.Equal(t, helper.Success, w.Code)
		mockservice.AssertExpectations(t)
	})

	t.Run("Failed Delete Role", func(t *testing.T) {
		req := httptest.NewRequest(helper.Delete, "/role/90/delete", nil)
		w := httptest.NewRecorder() 

		mockservice.On("DeleteRole", 90).Return(errors.New("Not Found ID: 90"))
		mux := http.NewServeMux()
		mux.HandleFunc("/role/{id}/delete", handler.Delete)
		mux.ServeHTTP(w, req)

		assert.Equal(t, helper.Notfound, w.Code)
		mockservice.AssertExpectations(t)

	})

	t.Run("Failed Path Value Delete Role", func(t *testing.T) {
		req := httptest.NewRequest(helper.Delete, "/role/abc/delete", nil)

		req.Header.Set("Content-Type", "application/json")
		req.SetPathValue("id", "abc")
		w := httptest.NewRecorder()
		mux:= http.NewServeMux()
		mux.HandleFunc("/role/{id}/delete", handler.Delete)
		mux.ServeHTTP(w, req)


		assert.Equal(t, helper.InternalServError, w.Code)
	})

	t.Run("Method Not Allowed", func(t *testing.T) {
		req := httptest.NewRequest(helper.Gets, "/role/90/delete", nil)
		req.Header.Set("Content-Type","application/json")
		w := httptest.NewRecorder() 

		mux := http.NewServeMux()
		mux.HandleFunc("/role/{id}/delete", handler.Delete)
		mux.ServeHTTP(w, req)

		assert.Equal(t, helper.MethodNotAllowed, w.Code)
	})
}