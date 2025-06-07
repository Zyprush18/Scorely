package bench

import (
	"testing"

	"github.com/Zyprush18/Scorely/test/services/mocks"
	"github.com/Zyprush18/Scorely/test/services/mocks/repo"
	"github.com/stretchr/testify/mock"
)

var Rolemock = &mocks.ServiceMock{Mock: mock.Mock{}}
var Reporole = &repo.RoleService{Repo: Rolemock}

func BenchmarkCreateRole(b *testing.B)  {
	Rolemock.Mock.On("AddRole", mock.AnythingOfType("Role")).Return("Success Create a New Role",nil)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_,_ =  Reporole.Create(mocks.Role{})
	}
}