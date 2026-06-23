package role

import (
	"context"
	"fmt"

	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/models/response"
	"github.com/stretchr/testify/mock"
)

type ServiceRole struct {
	mock.Mock
}

type LoggerMock struct {}

func (l LoggerMock) Logfile(msg string)  {
	fmt.Println(msg)
}

func (s *ServiceRole) GetAllData(ctx context.Context, search,sort string,page,perpage int) ([]response.Roles, int64,error)  {
	args := s.Called(ctx, search,sort,page,perpage)
	return args.Get(0).([]response.Roles), int64(args.Int(1)),args.Error(2)
}

func (s *ServiceRole) Create(ctx context.Context, data *request.Roles) error  {
	args := s.Called(ctx, data)
	return args.Error(0)
}

func (s *ServiceRole) ShowRoleById(ctx context.Context, id int) (*response.Roles, error)  {
	args := s.Called(ctx, id)
	return args.Get(0).(*response.Roles), args.Error(1)
}

func (s *ServiceRole) UpdateRole(ctx context.Context, id int, data *request.Roles) error  {
	args := s.Called(ctx, id, data)
	return args.Error(0)
}

func (s *ServiceRole) DeleteRole(ctx context.Context, id int) error  {
	args := s.Called(ctx, id)
	return args.Error(0)
}