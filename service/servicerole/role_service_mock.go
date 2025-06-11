package servicerole

import (
	"fmt"

	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/models/response"
	"github.com/stretchr/testify/mock"
)

type RepoRoleMock struct {
	mock.Mock
}


func (m *RepoRoleMock) CreateRole(data *request.Roles) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *RepoRoleMock) ShowById(id int) (*response.Roles, error)  {
	args := m.Called(id)
	return args.Get(0).(*response.Roles), args.Error(1)
}

func (m *RepoRoleMock) GetAllDataRole() ([]response.Roles, error)  {
	args := m.Called()

	fmt.Println("cuyyyy gua di panggil mock")


	// melakukan pengecekan 
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]response.Roles), args.Error(1)
}