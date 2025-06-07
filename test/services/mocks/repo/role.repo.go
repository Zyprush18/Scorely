package repo

import (
	"errors"

	"github.com/Zyprush18/Scorely/test/services/mocks"
)

type RoleService struct {
	Repo *mocks.ServiceMock
}

func (r RoleService) Create(data mocks.Role) (string,error)  {
	msg, err := r.Repo.AddRole(data)
	if err != nil {
		return "", errors.New(err.Error())
	}
	return msg,nil
}