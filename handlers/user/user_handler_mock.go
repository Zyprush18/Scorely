package user

import (
	"fmt"

	"github.com/Zyprush18/Scorely/models/request"
	"github.com/stretchr/testify/mock"
)

type MockUserServices struct {
	mock.Mock
}

type LoggerMock struct {}

func (l LoggerMock) Logfile(msg string)  {
	fmt.Println(msg)
}

func (m *MockUserServices) CreateUser(data *request.User) error {
	args := m.Called(data)
	return args.Error(0)
}