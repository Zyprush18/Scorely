package major

import (
	"fmt"

	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/models/response"
	"github.com/stretchr/testify/mock"
)

type ServiceMajorMock struct {
	mock.Mock
}

type LoggerMock struct{}
func (l LoggerMock) Logfile(msg string)  {
	fmt.Println(msg)
}

func (m *ServiceMajorMock) GetAllMajor(search, sort string, page, perpage int) ([]response.Majors, int64, error) {
	args := m.Called(search,sort,page,perpage)
	return  args.Get(0).([]response.Majors), int64(args.Int(1)), args.Error(2)
}

func (m *ServiceMajorMock) CreateMajor(data *request.Majors) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *ServiceMajorMock) ShowMajor(id int) (*response.Majors, error) {
	args := m.Called(id)
	return args.Get(0).(*response.Majors),args.Error(1)
}

func (m *ServiceMajorMock) UpdatedMajor(id int, data *request.Majors) error  {
	args := m.Called(id, data)
	return  args.Error(0)
}

func (m *ServiceMajorMock) DeleteMajor(id int) error {
	args := m.Called(id)
	return args.Error(0)
}