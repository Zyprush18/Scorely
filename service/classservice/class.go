package classservice

import (
	"context"
	"time"

	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/models/entity"
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/models/response"
	"github.com/Zyprush18/Scorely/repository/repoclass"
)

type ServiceClass interface {
	AllData(ctx context.Context, search, sort string, page, perpage int) ([]response.Class, int64, error)
	CreateClass(ctx context.Context, data *request.Class) error
	ShowClass(ctx context.Context, id int) (*response.Class, error)
	UpdateClass(ctx context.Context, id int, data *request.Class) error
	DeleteClass(ctx context.Context, id int) error
}

type ClassService struct {
	repo repoclass.RepoClass
}

func NewClassService(r repoclass.RepoClass) ServiceClass {
	return &ClassService{repo: r}
}

func (c *ClassService) AllData(ctx context.Context, search, sort string, page, perpage int) ([]response.Class, int64, error) {
	entities, count, err := c.repo.GetAll(ctx, search, sort, page, perpage)
	if err != nil {
		return nil, 0, err
	}
	return c.parseClassResponse(entities), count, nil
}

func (c *ClassService) CreateClass(ctx context.Context, data *request.Class) error {
	entity := &entity.Class{
		Name:    data.Name,
		LevelId: data.LevelId,
		MajorId: data.MajorId,
		Models: helper.Models{
			CreatedAt: time.Now(),
		},
	}
	return c.repo.Create(ctx, entity)
}

func (c *ClassService) ShowClass(ctx context.Context, id int) (*response.Class, error) {
	ent, err := c.repo.Show(ctx, id)
	if err != nil {
		return nil, err
	}
	return c.mapClassToResponse(ent), nil
}

func (c *ClassService) UpdateClass(ctx context.Context, id int, data *request.Class) error {
	entity := &entity.Class{
		Name:    data.Name,
		LevelId: data.LevelId,
		MajorId: data.MajorId,
		Models: helper.Models{
			UpdatedAt: time.Now(),
		},
	}
	return c.repo.Update(ctx, id, entity)
}

func (c *ClassService) DeleteClass(ctx context.Context, id int) error {
	return c.repo.Delete(ctx, id)
}

func (c *ClassService) parseClassResponse(entities []entity.Class) []response.Class {
	result := make([]response.Class, 0, len(entities))
	for _, v := range entities {
		result = append(result, *c.mapClassToResponse(&v))
	}
	return result
}

func (c *ClassService) mapClassToResponse(v *entity.Class) *response.Class {
	if v.Level == nil {
		v.Level = &entity.Levels{}
	}
	if v.Major == nil {
		v.Major = &entity.Majors{}
	}
	return &response.Class{
		IdClass: v.IdClass,
		Name:    v.Name,
		LevelId: v.LevelId,
		MajorId: v.MajorId,
		Level: response.Levels{
			IdLevel: v.Level.IdLevel,
			Level:   v.Level.Level,
			Model:  v.Level.Models,
		},
		Major: response.Majors{
			IdMajor:           v.Major.IdMajor,
			Major:             v.Major.Major,
			MajorAbbreviation: v.Major.MajorAbbreviation,
			Model:            v.Major.Models,
		},
		Model: v.Models,
	}
}