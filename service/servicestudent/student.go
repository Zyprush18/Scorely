package servicestudent

import (
	"context"
	"time"

	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/models/entity"
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/models/response"
	"github.com/Zyprush18/Scorely/repository/repostudent"
)

type ServiceStudent interface {
	GetAllStudent(ctx context.Context, Search, Sort string, Page, Perpage int) ([]response.Students, int64, error)
	CreateStudent(ctx context.Context, data *request.Students) error
	ShowStudent(ctx context.Context, id int) (*response.Students, error)
	UpdateStudent(ctx context.Context, id int, data *request.Students) error
	DeleteStudent(ctx context.Context, id int) error
}

type RepoStudent struct {
	repo repostudent.StudentRepo
}

func NewServiceStudent(r repostudent.StudentRepo) ServiceStudent {
	return &RepoStudent{repo: r}
}

func (r *RepoStudent) GetAllStudent(ctx context.Context, Search, Sort string, Page, Perpage int) ([]response.Students, int64, error) {
	entities, count, err := r.repo.GetAll(ctx, Search, Sort, Page, Perpage)
	if err != nil {
		return nil, 0, err
	}
	return parseStudentResponse(entities), count, nil
}

func (r *RepoStudent) CreateStudent(ctx context.Context, data *request.Students) error {
	phone, _ := parseUint(data.Phone)
	ent := &entity.Students{
		Name:    data.Name,
		Nisn:    data.Nisn,
		Gender:  data.Gender,
		Address: data.Address,
		Phone:   phone,
		UserId:  data.UserId,
		ClassId: data.ClassId,
		Model: helper.Models{
			CreatedAt: time.Now(),
		},
	}
	return r.repo.Create(ctx, ent)
}

func (r *RepoStudent) ShowStudent(ctx context.Context, id int) (*response.Students, error) {
	ent, err := r.repo.Show(ctx, id)
	if err != nil {
		return nil, err
	}
	return mapStudentToResponse(ent), nil
}

func (r *RepoStudent) UpdateStudent(ctx context.Context, id int, data *request.Students) error {
	phone, _ := parseUint(data.Phone)
	ent := &entity.Students{
		Name:    data.Name,
		Nisn:    data.Nisn,
		Gender:  data.Gender,
		Address: data.Address,
		Phone:   phone,
		UserId:  data.UserId,
		ClassId: data.ClassId,
		Model: helper.Models{
			UpdatedAt: time.Now(),
		},
	}
	return r.repo.Update(ctx, id, ent)
}

func (r *RepoStudent) DeleteStudent(ctx context.Context, id int) error {
	return r.repo.Delete(ctx, id)
}

func parseUint(s string) (uint, error) {
	var val uint
	for _, c := range []byte(s) {
		if c < '0' || c > '9' {
			return 0, nil
		}
		val = val*10 + uint(c-'0')
	}
	return val, nil
}

func parseStudentResponse(entities []entity.Students) []response.Students {
	result := make([]response.Students, len(entities))
	for i, v := range entities {
		result[i] = *mapStudentToResponse(&v)
	}
	return result
}

func mapStudentToResponse(v *entity.Students) *response.Students {
	if v.Class == nil {
		v.Class = &entity.Class{}
	}
	if v.Class.Level == nil {
		v.Class.Level = &entity.Levels{}
	}
	if v.Class.Major == nil {
		v.Class.Major = &entity.Majors{}
	}
	return &response.Students{
		IdStudent: v.IdStudent,
		Name:      v.Name,
		Nisn:      v.Nisn,
		Gender:    v.Gender,
		Address:   v.Address,
		Phone:     v.Phone,
		UserId:    v.UserId,
		ClassId:   v.ClassId,
		Class: response.Class{
			IdClass: v.Class.IdClass,
			Name:    v.Class.Name,
			LevelId: v.Class.LevelId,
			MajorId: v.Class.MajorId,
			Level: response.Levels{
				IdLevel: v.Class.Level.IdLevel,
				Level:   v.Class.Level.Level,
				Model:  v.Class.Level.Model,
			},
			Major: response.Majors{
				IdMajor:           v.Class.Major.IdMajor,
				Major:             v.Class.Major.Major,
				MajorAbbreviation: v.Class.Major.MajorAbbreviation,
				Model:            v.Class.Major.Model,
			},
		},
		Model: v.Model,
	}
}