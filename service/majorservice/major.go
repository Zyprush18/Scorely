package majorservice

import (
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/models/response"
	"github.com/Zyprush18/Scorely/repository/repomajor"
)

type MajorService interface {
	GetAllMajor(search, sort string, page, perpage int) ([]response.Majors, int64, error)
	CreateMajor(data *request.Majors) error
	ShowMajor(id int) (*response.Majors, error)
	UpdatedMajor(id int, data *request.Majors) error
	DeleteMajor(id int) error
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

func (m *MajorRepo) CreateMajor(data *request.Majors) error {
	return  m.Repo.Create(data)
}

func (m *MajorRepo) ShowMajor(id int) (*response.Majors, error) {
	return m.Repo.ShowById(id)
}

func (m *MajorRepo) UpdatedMajor(id int, data *request.Majors) error {
	return m.Repo.Updates(id, data)
}

func (m *MajorRepo) DeleteMajor(id int) error {
	return m.Repo.Deletes(id)
}