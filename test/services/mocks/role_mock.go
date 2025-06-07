package mocks

import "github.com/stretchr/testify/mock"

type Role struct {
	IdRole   uint 
	NameRole string 
}

// mock
type MockService interface {
	AddRole(r Role) (string,error)
}

type ServiceMock struct {
	Mock mock.Mock
}

func (m *ServiceMock) AddRole(r Role) (string,error) {
	args := m.Mock.Called(r)
	return args.String(0),args.Error(1)
}


