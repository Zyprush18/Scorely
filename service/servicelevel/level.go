package servicelevel

import (
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/models/response"
	"github.com/Zyprush18/Scorely/repository/repolevel"
)

type LevelService interface {
	GetAllLevel(search,sort string, page,perpage int) ([]response.Levels, int64, error)
	CreateLevel(data *request.Levels) error
	ShowLevel(id int) (*response.Levels, error)
}

type Repolevels struct {
	repo repolevel.LevelRepo
}

func ConnectRepo(r repolevel.LevelRepo) LevelService  {
	return &Repolevels{repo: r}
}

func (r *Repolevels) GetAllLevel(search,sort string, page,perpage int) ([]response.Levels, int64, error) {
	return r.repo.GetAll(search,sort,page,perpage)
}

func (r *Repolevels) CreateLevel(data *request.Levels) error  {
	return  r.repo.Create(data)
}

func (r *Repolevels) ShowLevel(id int) (*response.Levels, error) {
	return r.repo.Show(id)
}