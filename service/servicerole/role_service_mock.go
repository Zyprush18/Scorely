package servicerole

import (
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/stretchr/testify/mock"
)

type RepoRoleMock struct {
	mock.Mock
}

func (m *RepoRoleMock) CreateRole(data *request.Roles) error {
	args := m.Called(data)
	return args.Error(0)
}

