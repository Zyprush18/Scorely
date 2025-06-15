package user

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

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

		mockUser.On("CreateUser", data).Return(errors.New("Cannot Add or Update child row"))
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