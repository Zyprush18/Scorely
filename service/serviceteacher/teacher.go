package serviceteacher

import (
	"context"
	"time"

	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/models/entity"
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/models/response"
	"github.com/Zyprush18/Scorely/repository/repoteacher"
	"gorm.io/gorm"
)

type ServiceTeacher interface {
	GetAllTeacher(ctx context.Context, Search, Sort string, Page, Perpage int) ([]response.Teachers, int64, error)
	CreateTeacher(ctx context.Context, data *request.Teachers) error
	ShowTeacher(ctx context.Context, id int) (*response.Teachers, error)
	UpdateTeacher(ctx context.Context, id int, data *request.Teachers) error
	DeleteTeacher(ctx context.Context, id int) error
}

type RepoTeacherStruct struct {
	repo repoteacher.RepoTeacher
	db   *gorm.DB
}

func ConnectRepo(r repoteacher.RepoTeacher, db *gorm.DB) ServiceTeacher {
	return &RepoTeacherStruct{repo: r, db: db}
}

func (r *RepoTeacherStruct) GetAllTeacher(ctx context.Context, Search, Sort string, Page, Perpage int) ([]response.Teachers, int64, error) {
	entities, count, err := r.repo.GetAll(ctx, Search, Sort, Page, Perpage)
	if err != nil {
		return nil, 0, err
	}
	return response.RespGetALl(entities), count, nil
}

func (r *RepoTeacherStruct) CreateTeacher(ctx context.Context, data *request.Teachers) error {
	var subjectfind []entity.Subjects
	if err := r.db.WithContext(ctx).Model(&entity.Subjects{}).Where("id_subject IN ?", data.SubjectId).Find(&subjectfind).Error; err != nil {
		return err
	}

	ent := &entity.Teachers{
		Name:    data.Name,
		Nip:     data.Nip,
		Gender:  data.Gender,
		Address: data.Address,
		Phone:   data.Phone,
		UserId:  data.UserId,
		Subject: subjectfind,
		Models: helper.Models{
			CreatedAt: time.Now(),
		},
	}
	return r.repo.Create(ctx, ent)
}

func (r *RepoTeacherStruct) ShowTeacher(ctx context.Context, id int) (*response.Teachers, error) {
	ent, err := r.repo.Show(ctx, id)
	if err != nil {
		return nil, err
	}
	return &response.Teachers{
		IdTeacher: ent.IdTeacher,
		Name:      ent.Name,
		Nip:       ent.Nip,
		Gender:    ent.Gender,
		Address:   ent.Address,
		Phone:     ent.Phone,
		UserId:    ent.UserId,
		Subject:   response.Subjectsresp(ent.Subject),
		Model:    ent.Models,
	}, nil
}

func (r *RepoTeacherStruct) UpdateTeacher(ctx context.Context, id int, data *request.Teachers) error {
	var findsubject []entity.Subjects
	if err := r.db.WithContext(ctx).Model(&entity.Subjects{}).Where("id_subject IN ?", data.SubjectId).Find(&findsubject).Error; err != nil {
		return err
	}

	ent := &entity.Teachers{
		Name:    data.Name,
		Nip:     data.Nip,
		Gender:  data.Gender,
		Address: data.Address,
		Phone:   data.Phone,
		UserId:  data.UserId,
		Subject: findsubject,
		Models: helper.Models{
			UpdatedAt: time.Now(),
		},
	}
	return r.repo.Update(ctx, id, ent)
}

func (r *RepoTeacherStruct) DeleteTeacher(ctx context.Context, id int) error {
	return r.repo.Delete(ctx, id)
}