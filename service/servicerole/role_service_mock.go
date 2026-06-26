package servicerole

import (
	"context"

	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/models/response"
	"github.com/stretchr/testify/mock"
)

type RepoRoleMock struct {
	mock.Mock
}

func (m *RepoRoleMock) GetAllDataRole(ctx context.Context, search, sort string, page, perpage int) ([]response.Roles, int64, error) {
	args := m.Called(ctx, search, sort, page, perpage)

	return args.Get(0).([]response.Roles), int64(args.Int(1)), args.Error(2)
}

func (m *RepoRoleMock) CreateRole(ctx context.Context, data *request.Roles) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}

func (m *RepoRoleMock) ShowById(ctx context.Context, id int) (*response.Roles, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*response.Roles), args.Error(1)
}

func (m *RepoRoleMock) UpdateRole(ctx context.Context, id int, data *request.Roles) error {
	args := m.Called(ctx, id, data)
	return args.Error(0)
}

func (m *RepoRoleMock) DeleteRole(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
