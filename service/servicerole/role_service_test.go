package servicerole

import (
	"errors"
	"testing"

	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/models/response"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

	var Mockservice = RepoRoleMock{Mock: mock.Mock{}}

func TestServiceRole(t *testing.T) {
	service := NewRoleService(&Mockservice)
	t.Run("Service Success Create a New Role", func(t *testing.T) {
		rolePass := &request.Roles{
		NameRole: "Admin",
	}
		Mockservice.On("CreateRole", rolePass).Return(nil)

		err := service.Create(rolePass)
		assert.NoError(t, err)
		Mockservice.Mock.AssertExpectations(t)

	})

	t.Run("Service Failed Create a New Role", func(t *testing.T) {
		roleFails := &request.Roles{
		NameRole: "",
	}
		Mockservice.On("CreateRole", roleFails).Return(errors.New("failed"))
		errs := service.Create(roleFails)

		assert.Error(t, errs)
		Mockservice.Mock.AssertExpectations(t)
	})

}

func TestShowRoleById(t *testing.T)  {
	servicerole := NewRoleService(&Mockservice)
	data:= &response.Roles{
			IdRole: 1,
			NameRole: "Admin",
		}
	t.Run("Success Show Role by id", func(t *testing.T) {
		
		Mockservice.On("ShowById", 1).Return(data, nil) 

		resp, err := servicerole.ShowRoleById(1)
		assert.NoError(t, err)
		assert.Equal(t, uint(1), resp.IdRole)
		assert.Equal(t, "Admin", resp.NameRole)

		Mockservice.AssertExpectations(t)
	})

	t.Run("Failed Show Role by id", func(t *testing.T) {
		
		Mockservice.On("ShowById", 2).Return(data, errors.New("Not Found role id: 2")) 

		resp, err := servicerole.ShowRoleById(2)

		assert.Error(t, err)
		assert.Nil(t, resp)

		Mockservice.AssertExpectations(t)
	})
}
