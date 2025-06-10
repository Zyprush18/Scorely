package role

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/models/response"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var Mockservice = &ServiceRole{Mock: mock.Mock{}}
var Mocklogger = &LoggerMock{}

func TestHandlerCreate(t *testing.T) {

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