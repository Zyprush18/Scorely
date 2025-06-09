package servicerole

import (
	"errors"
	"testing"

	"github.com/Zyprush18/Scorely/models/request"
	"github.com/stretchr/testify/assert"
)

func TestServiceRole(t *testing.T) {
	mockservice := new(RepoRoleMock)
	service := NewRoleService(mockservice)
	t.Run("Service Success Create a New Role", func(t *testing.T) {
		rolePass := &request.Roles{
		NameRole: "Admin",
	}
		mockservice.On("CreateRole", rolePass).Return(nil)

		err := service.Create(rolePass)
		assert.NoError(t, err)
		mockservice.Mock.AssertExpectations(t)

	})

	t.Run("Service Failed Create a New Role", func(t *testing.T) {
		roleFails := &request.Roles{
		NameRole: "",
	}
		mockservice.On("CreateRole", roleFails).Return(errors.New("failed"))
		errs := service.Create(roleFails)

		assert.Error(t, errs)
		mockservice.Mock.AssertExpectations(t)
	})

}
