package user

import (
	"fmt"

	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/models/response"
	"github.com/stretchr/testify/mock"
)

type MockUserServices struct {
	mock.Mock
}

type LoggerMock struct {}

func (l LoggerMock) Logfile(msg string)  {
	fmt.Println(msg)
}

func (m *MockUserServices) AllUser(search, sort string, page,perpage int)([]response.Users, int64,error) {
	args := m.Called(search,sort,page,perpage)
	return args.Get(0).([]response.Users), int64(args.Int(1)),args.Error(2)
}

func (m *MockUserServices) CreateUser(data *request.User) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *MockUserServices) ShowUser(id int) (*response.Users, error)  {
	args := m.Called(id)
	return args.Get(0).(*response.Users), args.Error(1)
}

func (m *MockUserServices) UpdateUser(id int, data *request.User) error  {
	args := m.Called(id, data)
	return args.Error(0)
}
func (m *MockUserServices) DeleteUser(id int) error  {
	args := m.Called(id)
	return  args.Error(0)
}