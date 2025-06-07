package services

import (
	"testing"

	"github.com/Zyprush18/Scorely/test/services/mocks"
	"github.com/Zyprush18/Scorely/test/services/mocks/repo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var Rolemock = &mocks.ServiceMock{Mock: mock.Mock{}}
var Reporole = &repo.RoleService{Repo: Rolemock}

func TestAddRole(t *testing.T)  {
	t.Run("Success_Add", func(t *testing.T) {
		role := mocks.Role{
			IdRole: 1,
			NameRole: "Admin",
		}

		Rolemock.Mock.On("AddRole", mock.AnythingOfType("Role")).Return("Success Create a New Role",nil)

		msg, err := Reporole.Create(role)

		assert.NoError(t, err)
		assert.Equal(t, "Success Create a New Role", msg)

	})
	Rolemock.Mock.AssertExpectations(t)
}