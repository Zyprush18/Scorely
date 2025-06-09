package role

import (
	"fmt"

	"github.com/Zyprush18/Scorely/models/request"
	"github.com/stretchr/testify/mock"
)

type ServiceRole struct {
	mock.Mock
}

func (s *ServiceRole) Create(data *request.Roles) error  {
	args := s.Called(data)
	return args.Error(0)
}

type LoggerMock struct {}

func (l LoggerMock) Logfile(msg string)  {
	fmt.Println(msg)
}
