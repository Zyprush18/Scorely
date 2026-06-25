package servicelevel

import (
	"context"
	"time"

	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/models/entity"
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/models/response"
	"github.com/Zyprush18/Scorely/repository/repolevel"
)

type LevelService interface {
	GetAllLevel(ctx context.Context, search, sort string, page, perpage int) ([]response.Levels, int64, error)
	CreateLevel(ctx context.Context, data *request.Levels) error
	ShowLevel(ctx context.Context, id int) (*response.Levels, error)
	UpdateLevel(ctx context.Context, id int, data *request.Levels) error
	DeleteLevel(ctx context.Context, id int) error
}

type Repolevels struct {
	repo repolevel.LevelRepo
}

func ConnectRepo(r repolevel.LevelRepo) LevelService {
	return &Repolevels{repo: r}
}

func (r *Repolevels) GetAllLevel(ctx context.Context, search, sort string, page, perpage int) ([]response.Levels, int64, error) {
	entities, count, err := r.repo.GetAll(ctx, search, sort, page, perpage)
	if err != nil {
		return nil, 0, err
	}
	return parseLevelResponse(entities), count, nil
}

func (r *Repolevels) CreateLevel(ctx context.Context, data *request.Levels) error {
	ent := &entity.Levels{
		Level: data.Level,
		Models: helper.Models{
			CreatedAt: time.Now(),
		},
	}
	return r.repo.Create(ctx, ent)
}

func (r *Repolevels) ShowLevel(ctx context.Context, id int) (*response.Levels, error) {
	ent, err := r.repo.Show(ctx, id)
	if err != nil {
		return nil, err
	}
	return &response.Levels{
		IdLevel: ent.IdLevel,
		Level:   ent.Level,
		Model:  ent.Models,
	}, nil
}

func (r *Repolevels) UpdateLevel(ctx context.Context, id int, data *request.Levels) error {
	ent := &entity.Levels{
		Level: data.Level,
		Models: helper.Models{
			UpdatedAt: time.Now(),
		},
	}
	return r.repo.Update(ctx, id, ent)
}

func (r *Repolevels) DeleteLevel(ctx context.Context, id int) error {
	return r.repo.Delete(ctx, id)
}

func parseLevelResponse(entities []entity.Levels) []response.Levels {
	result := make([]response.Levels, len(entities))
	for i, v := range entities {
		result[i] = response.Levels{
			IdLevel: v.IdLevel,
			Level:   v.Level,
			Model:  v.Models,
		}
	}
	return result
}