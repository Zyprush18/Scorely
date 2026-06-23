package majorservice

import (
	"github.com/Zyprush18/Scorely/models/entity"
	"github.com/Zyprush18/Scorely/models/response"
	"github.com/stretchr/testify/mock"
)

type RepoMajorMock struct {
	mock.Mock
}

func (m *RepoMajorMock) GetAllData(search, sort string, page, perpage int) ([]response.Majors, int64, error)  {
	args := m.Called(search,sort,page,perpage)
	return args.Get(0).([]response.Majors), int64(args.Int(1)), args.Error(2)
}

func (m *RepoMajorMock) Create(data *entity.Majors) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *RepoMajorMock) ShowById(id int) (*response.Majors, error) {
	args := m.Called(id)
	return args.Get(0).(*response.Majors),args.Error(1)
}

func (m *RepoMajorMock) Updates(id int, data *entity.Majors) error {
	args := m.Called(id,data)
	return args.Error(0)
}

func (m *RepoMajorMock) Deletes(id int) error {
	args := m.Called(id)
	return args.Error(0)
}