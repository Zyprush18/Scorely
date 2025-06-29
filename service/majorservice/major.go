package majorservice

import (
	"github.com/Zyprush18/Scorely/models/response"
	"github.com/Zyprush18/Scorely/repository/repomajor"
)

type MajorService interface {
	GetAllMajor(search, sort string, page, perpage int) ([]response.Majors, int64, error)
}

type MajorRepo struct {
	Repo repomajor.MajorRepo
}

func RepoMajorConn(r repomajor.MajorRepo) MajorService {
	return &MajorRepo{Repo: r}
}

func (m *MajorRepo) GetAllMajor(search, sort string, page, perpage int) ([]response.Majors, int64, error) {
	return m.Repo.GetAllData(search,sort,page,perpage)
}
