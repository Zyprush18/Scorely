package servicerole

import (
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/models/response"
	"github.com/stretchr/testify/mock"
)

type RepoRoleMock struct {
	mock.Mock
}

// GetAllDataRole implements reporole.RoleService.
func (m *RepoRoleMock) GetAllDataRole(search string, sort string, page int, perpage int) ([]response.Roles, int64, error) {
	panic("unimplemented")
}

func (m *RepoRoleMock) CreateRole(data *request.Roles) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *RepoRoleMock) ShowById(id int) (*response.Roles, error) {
	args := m.Called(id)
	return args.Get(0).(*response.Roles), args.Error(1)
}

func (m *RepoRoleMock) GetAllData(search, sort string, page, perpage int) ([]response.Roles, int64, error) {
	args := m.Called(search, sort, page, perpage)

	return args.Get(0).([]response.Roles), int64(args.Int(1)), args.Error(2)
}

func (m *RepoRoleMock) UpdateRole(id int, data *request.Roles) error {
	args := m.Called(id, data)
	return args.Error(0)
}

func (m *RepoRoleMock) DeleteRole(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
