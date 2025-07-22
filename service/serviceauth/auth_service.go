package serviceauth

import (
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/repository/repoauth"
)


type AuthService interface {
	Loginuser(data *request.Login) (string,error)
}

type ServiceRepoAuth struct {
	repo repoauth.RepoAuth
}

func ConnectRepo(r repoauth.RepoAuth)  AuthService {
	return &ServiceRepoAuth{repo: r}
}

func (s *ServiceRepoAuth) Loginuser(data *request.Login) (string,error) {
	return s.repo.Login(data)
}