package user

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

func TestAllUser(t *testing.T) {
	t.Run("Method Not Allowed", func(t *testing.T) {
		mockUser := MockUserServices{Mock: mock.Mock{}}
		logger := LoggerMock{}
		userHandler := NewHandlerUser(&mockUser, logger)
		req := httptest.NewRequest(helper.Post, "/user", nil)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		userHandler.GetAllUser(w, req)

		assert.Equal(t, helper.MethodNotAllowed, w.Code)
	})

	t.Run("Failed Get All user", func(t *testing.T) {
		mockUser := MockUserServices{Mock: mock.Mock{}}
		logger := LoggerMock{}
		userHandler := NewHandlerUser(&mockUser, logger)
		req := httptest.NewRequest(helper.Gets, "/user", nil)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		mockUser.On("AllUser").Return([]response.Users(nil), errors.New("something went wrong"))

		userHandler.GetAllUser(w, req)

		assert.Equal(t, helper.InternalServError, w.Code)
		mockUser.AssertExpectations(t)
	})

	t.Run("Success Get All user", func(t *testing.T) {
		mockUser := MockUserServices{Mock: mock.Mock{}}
		logger := LoggerMock{}
		userHandler := NewHandlerUser(&mockUser, logger)
		data := []response.Users{
			{
				IdUser:   3,
				Email:    "user@gmail.com",
				Password: "user123",
				RoleId:   2,
				Models: helper.Models{
					CreatedAt: time.Now(),
				},
			},
			{
				IdUser:   4,
				Email:    "user2@gmail.com",
				Password: "user123",
				RoleId:   2,
				Models: helper.Models{
					CreatedAt: time.Now(),
				},
			},
		}

		req := httptest.NewRequest(helper.Gets, "/user", nil)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		mockUser.On("AllUser").Return(data, nil)

		userHandler.GetAllUser(w, req)

		assert.Equal(t, helper.Success, w.Code)
		mockUser.AssertExpectations(t)

	})
}

func TestCreateUser(t *testing.T) {
	mockUser := MockUserServices{Mock: mock.Mock{}}
	logger := LoggerMock{}
	userHandler := NewHandlerUser(&mockUser, logger)

	t.Run("Method Not Allowed", func(t *testing.T) {
		data := &request.User{
			Email:    "Admin@gmail.com",
			Password: "admin123",
			RoleId:   1,
		}
		jmarshal, _ := json.Marshal(data)

		req := httptest.NewRequest(helper.Gets, "/add/user", bytes.NewReader(jmarshal))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		userHandler.Create(w, req)
		assert.Equal(t, helper.MethodNotAllowed, w.Code)
	})

	t.Run("Body request is missing", func(t *testing.T) {
		req := httptest.NewRequest(helper.Post, "/add/user", nil)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		userHandler.Create(w, req)
		assert.Equal(t, helper.BadRequest, w.Code)
	})

	t.Run("Validation Error", func(t *testing.T) {
		data := &request.User{
			Email:    "",
			Password: "admin123",
			RoleId:   1,
		}
		jmarshal, _ := json.Marshal(data)

		req := httptest.NewRequest(helper.Post, "/add/user", bytes.NewReader(jmarshal))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		userHandler.Create(w, req)
		assert.Equal(t, helper.UnprocessbleEntity, w.Code)
	})

	t.Run("Failed Create a New User (database refused)", func(t *testing.T) {
		data := &request.User{
			Email:    "Admin@gmail.com",
			Password: "admin123",
			RoleId:   1,
		}
		jmarshal, _ := json.Marshal(data)

		req := httptest.NewRequest(helper.Post, "/add/user", bytes.NewReader(jmarshal))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		mockUser.On("CreateUser", data).Return(errors.New("Cannot Add child row"))
		userHandler.Create(w, req)
		assert.Equal(t, helper.InternalServError, w.Code)
		mockUser.AssertExpectations(t)
	})

	t.Run("Failed Create a New User (Duplicate Email)", func(t *testing.T) {
		data := &request.User{
			Email:    "Admin33@gmail.com",
			Password: "admin123",
			RoleId:   1,
		}

		dupErr := &mysql.MySQLError{
			Number:  1062,
			Message: "Duplicate entry",
		}
		jmarshal, _ := json.Marshal(data)

		req := httptest.NewRequest(helper.Post, "/add/user", bytes.NewReader(jmarshal))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		mockUser.On("CreateUser", data).Return(dupErr)
		userHandler.Create(w, req)
		assert.Equal(t, helper.Conflict, w.Code)
		mockUser.AssertExpectations(t)
	})

	t.Run("Success Create a New User", func(t *testing.T) {
		data := &request.User{
			Email:    "Admin@gmail.com",
			Password: "admin123456",
			RoleId:   1,
		}
		jmarshal, _ := json.Marshal(data)

		req := httptest.NewRequest(helper.Post, "/add/user", bytes.NewReader(jmarshal))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		mockUser.On("CreateUser", data).Return(nil)
		userHandler.Create(w, req)
		assert.Equal(t, helper.Created, w.Code)
		mockUser.AssertExpectations(t)
	})
}

func TestHandlerShow(t *testing.T) {
	mockuser := MockUserServices{}
	loggeruser := LoggerMock{}
	handler := NewHandlerUser(&mockuser, loggeruser)

	t.Run("Method Not Allowed", func(t *testing.T) {
		req := httptest.NewRequest(helper.Post, "/user/1", nil)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		mux := http.NewServeMux()
		mux.HandleFunc("/user/{id}", handler.Show)
		mux.ServeHTTP(w, req)

		assert.Equal(t, helper.MethodNotAllowed, w.Code)
	})

	t.Run("Invalid User id format", func(t *testing.T) {
		req := httptest.NewRequest(helper.Gets, "/user/abc", nil)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		mux := http.NewServeMux()
		mux.HandleFunc("/user/{id}", handler.Show)
		mux.ServeHTTP(w, req)

		assert.Equal(t, helper.BadRequest, w.Code)
	})

	t.Run("Failed: Not Found Id user", func(t *testing.T) {
		dataUser := &response.Users{
			IdUser:   1,
			Email:    "Admin@gmail.com",
			Password: "admin123",
			RoleId:   1,
		}
		req := httptest.NewRequest(helper.Gets, "/user/67", nil)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		mockuser.On("ShowUser", 67).Return(dataUser, gorm.ErrRecordNotFound)
		mux := http.NewServeMux()
		mux.HandleFunc("/user/{id}", handler.Show)
		mux.ServeHTTP(w, req)

		assert.Equal(t, helper.Notfound, w.Code)
		mockuser.AssertExpectations(t)
	})

	t.Run("Failed: Database Refused", func(t *testing.T) {
		dataUser := &response.Users{
			IdUser:   5,
			Email:    "Adminfail@gmail.com",
			Password: "admin123",
			RoleId:   1,
		}
		req := httptest.NewRequest(helper.Gets, "/user/5", nil)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		mockuser.On("ShowUser", 5).Return(dataUser, errors.New("Database Is Refused"))
		mux := http.NewServeMux()
		mux.HandleFunc("/user/{id}", handler.Show)
		mux.ServeHTTP(w, req)

		assert.Equal(t, helper.InternalServError, w.Code)
		mockuser.AssertExpectations(t)
	})

	t.Run("Success Show User By id", func(t *testing.T) {
		dataUser := &response.Users{
			IdUser:   1,
			Email:    "Admin@gmail.com",
			Password: "admin123",
			RoleId:   1,
		}
		req := httptest.NewRequest(helper.Gets, "/user/1", nil)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		mockuser.On("ShowUser", 1).Return(dataUser, nil)
		mux := http.NewServeMux()
		mux.HandleFunc("/user/{id}", handler.Show)
		mux.ServeHTTP(w, req)

		assert.Equal(t, helper.Success, w.Code)
		mockuser.AssertExpectations(t)
	})
}

func TestHandlerUpdate(t *testing.T)  {
	mockUser := MockUserServices{Mock: mock.Mock{}}
	logger := LoggerMock{}
	userHandler := NewHandlerUser(&mockUser, logger)
	mux := http.NewServeMux()
	mux.HandleFunc("/user/{id}/update", userHandler.Update)

	t.Run("Method Not Allowed", func(t *testing.T) {
		datareq := &request.User{
			Email: "admin@gmail.com",
		}

		jsom , err := json.Marshal(datareq)
		assert.NoError(t, err)
		req := httptest.NewRequest(helper.Gets, "/user/4/update", bytes.NewReader(jsom))
		req.Header.Set("Content-Type","application/json")
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, helper.MethodNotAllowed, w.Code)
	})

	t.Run("Body Is Missing", func(t *testing.T) {
		req := httptest.NewRequest(helper.Put, "/user/4/update", nil)
		req.Header.Set("Content-Type","application/json")
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, helper.BadRequest, w.Code)
	})
	
	t.Run("Invalid Format Id", func(t *testing.T) {
		datareq := &request.User{
			Email: "admin11@gmail.com",
		}

		jsom , err := json.Marshal(datareq)
		assert.NoError(t, err)
		req := httptest.NewRequest(helper.Put, "/user/abc/update", bytes.NewReader(jsom))
		req.Header.Set("Content-Type","application/json")
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, helper.BadRequest, w.Code)
	})

	t.Run("Failed Update: Not Found Id",func(t *testing.T) {
		datareq := &request.User{
			Email: "admin12@gmail.com",
		}

		jsom , err := json.Marshal(datareq)
		assert.NoError(t, err)
		req := httptest.NewRequest(helper.Put, "/user/90/update", bytes.NewReader(jsom))
		req.Header.Set("Content-Type","application/json")
		w := httptest.NewRecorder()
		mockUser.On("UpdateUser", 90, datareq).Return(gorm.ErrRecordNotFound)

		mux.ServeHTTP(w, req)

		assert.Equal(t, helper.Notfound, w.Code)
		mockUser.AssertExpectations(t)
	})

	t.Run("Failed Update: Email is Exist", func(t *testing.T) {
		datareq := &request.User{
			Email: "admin125@gmail.com",
		}

		dupErr := &mysql.MySQLError{
			Number:  1062,
			Message: "Duplicate entry",
		}

		jsom , err := json.Marshal(datareq)
		assert.NoError(t, err)
		req := httptest.NewRequest(helper.Put, "/user/1/update", bytes.NewReader(jsom))
		req.Header.Set("Content-Type","application/json")
		w := httptest.NewRecorder()
		mockUser.On("UpdateUser", 1, datareq).Return(dupErr)

		mux.ServeHTTP(w, req)

		assert.Equal(t, helper.Conflict, w.Code)
		mockUser.AssertExpectations(t)
	})

	t.Run("Failed Update: Internal Server Error", func(t *testing.T) {
		datareq := &request.User{
			Email: "admin13@gmail.com",
		}

		jsom , err := json.Marshal(datareq)
		assert.NoError(t, err)
		req := httptest.NewRequest(helper.Put, "/user/9/update", bytes.NewReader(jsom))
		req.Header.Set("Content-Type","application/json")
		w := httptest.NewRecorder()
		mockUser.On("UpdateUser", 9, datareq).Return(errors.New("Cannot Update Child Row"))

		mux.ServeHTTP(w, req)

		assert.Equal(t, helper.InternalServError, w.Code)
		mockUser.AssertExpectations(t)
	})

	t.Run("Success Update User",func(t *testing.T) {
		datareq := &request.User{
			Email: "admin123@gmail.com",
		}

		jsom , err := json.Marshal(datareq)
		assert.NoError(t, err)
		req := httptest.NewRequest(helper.Put, "/user/1/update", bytes.NewReader(jsom))
		req.Header.Set("Content-Type","application/json")
		w := httptest.NewRecorder()
		mockUser.On("UpdateUser", 1, datareq).Return(nil)

		mux.ServeHTTP(w, req)

		assert.Equal(t, helper.Success, w.Code)
		mockUser.AssertExpectations(t)
	})
}