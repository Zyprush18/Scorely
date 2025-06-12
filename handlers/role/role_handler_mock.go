package role

import (
	"fmt"

	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/models/response"
	"github.com/stretchr/testify/mock"
)

type ServiceRole struct {
	mock.Mock
}

func (s *ServiceRole) Create(data *request.Roles) error  {
	args := s.Called(data)
	return args.Error(0)
}

func (s *ServiceRole) ShowRoleById(id int) (*response.Roles, error)  {
	args := s.Called(id)
	return args.Get(0).(*response.Roles), args.Error(1)
}

type LoggerMock struct {}

func (l LoggerMock) Logfile(msg string)  {
	fmt.Println(msg)
}

func (s *ServiceRole) GetAllData() ([]response.Roles, error)  {
	args := s.Called()
	return args.Get(0).([]response.Roles), args.Error(1)
}

func (s *ServiceRole) UpdateRole(id int, data *request.Roles) error  {
	args := s.Called(id, data)
	return args.Error(0)
}