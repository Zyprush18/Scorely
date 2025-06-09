package role

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandlerCreate(t *testing.T) {
	mockservice := new(ServiceRole)
	mocklogger := new(LoggerMock)

	handler := RoleHandler(mockservice, mocklogger)

	t.Run("Succes Create a New Role", func(t *testing.T) {
		body := &request.Roles{
			NameRole: "Admin",
		}

		jsonBody, _ := json.Marshal(body)

		req := httptest.NewRequest(helper.Post, "/add/role", bytes.NewReader(jsonBody))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		mockservice.On("Create", body).Return(nil)
		handler.AddRoles(w, req)

		assert.Equal(t, helper.Created, w.Code)
		mockservice.AssertExpectations(t)
	})

	t.Run("Failed Create a New Role", func(t *testing.T) {
		bodyFail := &request.Roles{
			NameRole: "",
		}
		jsonBody, _ := json.Marshal(bodyFail)
		req := httptest.NewRequest(helper.Post, "/add/role", bytes.NewReader(jsonBody))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		mockservice.On("Create", mock.Anything).Return(errors.New("Failed Create a New Role"))
		handler.AddRoles(w, req)

		assert.Equal(t, helper.BadRequest, w.Code)
		mockservice.AssertExpectations(t)
	})
}
