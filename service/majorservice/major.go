package majorservice

import (
	"context"
	"time"

	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/models/entity"
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/models/response"
	"github.com/Zyprush18/Scorely/repository/repomajor"
)

type MajorService interface {
	GetAllMajor(ctx context.Context,search, sort string, page, perpage int) ([]response.Majors, int64, error)
	CreateMajor(ctx context.Context,data *request.Majors) error
	ShowMajor(ctx context.Context, id int) (*response.Majors, error)
	UpdatedMajor(ctx context.Context, id int, data *request.Majors) error
	DeleteMajor(ctx context.Context, id int) error
}

type MajorRepo struct {
	Repo repomajor.MajorRepo
}

func RepoMajorConn(r repomajor.MajorRepo) MajorService {
	return &MajorRepo{Repo: r}
}

func (m *MajorRepo) GetAllMajor(ctx context.Context,search, sort string, page, perpage int) ([]response.Majors, int64, error) {
	entities, count, err := m.Repo.GetAllData(ctx,search, sort, page, perpage)
	if err != nil {
		return nil, 0, err
	}
	return m.parseMajorResponse(entities), count, nil
}

func (m *MajorRepo) CreateMajor(ctx context.Context, data *request.Majors) error {
	ent := &request.Majors{
		Major:             data.Major,
		MajorAbbreviation: data.MajorAbbreviation,
		Model: helper.Models{
			CreatedAt: time.Now(),
		},
	}
	return m.Repo.Create(ctx,ent)
}

func (m *MajorRepo) ShowMajor(ctx context.Context, id int) (*response.Majors, error) {
	ent, err := m.Repo.ShowById(ctx, id)
	if err != nil {
		return nil, err
	}

	result := m.mapMajorResponse(ent)

	return  &result, nil
}

func (m *MajorRepo) UpdatedMajor(ctx context.Context, id int, data *request.Majors) error {
	ent := &request.Majors{
		Major:             data.Major,
		MajorAbbreviation: data.MajorAbbreviation,
		Model: helper.Models{
			UpdatedAt: time.Now(),
		},
	}
	return m.Repo.Updates(ctx, id, ent)
}

func (m *MajorRepo) DeleteMajor(ctx context.Context, id int) error {
	return m.Repo.Deletes(ctx, id)
}

func (m *MajorRepo) mapMajorResponse(ent *entity.Majors) response.Majors  {
	return  response.Majors{
			IdMajor:           ent.IdMajor,
			Major:             ent.Major,
			MajorAbbreviation: ent.MajorAbbreviation,
			Model:            ent.Model,
		}
}

func (m *MajorRepo) parseMajorResponse(entities []entity.Majors) []response.Majors {
	result := make([]response.Majors, len(entities))
	for i, v := range entities {
		result[i] = m.mapMajorResponse(&v)
	}
	return result
}