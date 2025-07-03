package servicelevel

import (
	"github.com/Zyprush18/Scorely/models/response"
	"github.com/Zyprush18/Scorely/repository/repolevel"
)

type LevelService interface {
	GetAllLevel(search,sort string, page,perpage int) ([]response.Levels, int64, error)
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