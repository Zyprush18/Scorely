package subjectservice

import (
	"context"
	"time"

	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/models/entity"
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/models/response"
	"github.com/Zyprush18/Scorely/repository/reposubject"
)

type ServiceSubject interface {
	GetAllSubject(ctx context.Context, Search, Sort string, Page, Perpage int) ([]response.Subjects, int64, error)
	CreateSubject(ctx context.Context, data *request.Subjects) error
	ShowSubject(ctx context.Context, id int) (*response.Subjects, error)
	UpdateSubject(ctx context.Context, id int, data *request.Subjects) error
	DeleteSubject(ctx context.Context, id int) error
}

type RepoStruct struct {
	repo reposubject.RepoSubject
}

func ConnectRepo(r reposubject.RepoSubject) ServiceSubject {
	return &RepoStruct{repo: r}
}

func (r *RepoStruct) GetAllSubject(ctx context.Context, Search, Sort string, Page, Perpage int) ([]response.Subjects, int64, error) {
	entities, count, err := r.repo.GetAll(ctx, Search, Sort, Page, Perpage)
	if err != nil {
		return nil, 0, err
	}
	return response.Subjectsresp(entities), count, nil
}

func (r *RepoStruct) CreateSubject(ctx context.Context, data *request.Subjects) error {
	ent := &entity.Subjects{
		NameSubject: data.NameSubject,
		Semester:    data.Semester,
		Model: helper.Models{
			CreatedAt: time.Now(),
		},
	}
	return r.repo.Create(ctx, ent)
}

func (r *RepoStruct) ShowSubject(ctx context.Context, id int) (*response.Subjects, error) {
	ent, err := r.repo.Show(ctx, id)
	if err != nil {
		return nil, err
	}
	return &response.Subjects{
		IdSubject:   ent.IdSubject,
		NameSubject: ent.NameSubject,
		Semester:    ent.Semester,
		Model:      ent.Model,
	}, nil
}

func (r *RepoStruct) UpdateSubject(ctx context.Context, id int, data *request.Subjects) error {
	ent := &entity.Subjects{
		NameSubject: data.NameSubject,
		Semester:    data.Semester,
		Model: helper.Models{
			UpdatedAt: time.Now(),
		},
	}
	return r.repo.Update(ctx, id, ent)
}

func (r *RepoStruct) DeleteSubject(ctx context.Context, id int) error {
	return r.repo.Delete(ctx, id)
}