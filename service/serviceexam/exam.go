package serviceexam

import (
	"context"
	"fmt"
	"time"

	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/models/entity"
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/models/response"
	"github.com/Zyprush18/Scorely/repository/repoexams"
)

type ServiceExams interface {
	GetAllExams(ctx context.Context, Search, Sort string, Page, Perpage int) ([]response.Exams, int64, error)
	FindExamsbyIdTeacher(ctx context.Context, Search, Sort string, Page, Perpage, id int) ([]response.Exams, int64, error)
	CreateExams(ctx context.Context, data *request.Exams, role string, user_id, subject_id int) error
	ShowExams(ctx context.Context, id, userid int, coderole string) (*response.Exams, error)
	UpdateExam(ctx context.Context, data *request.Exams, role string, id, userid int) error
	DeleteExam(ctx context.Context, id, userid int, coderole string) error
}

type RepoExams struct {
	repo repoexams.RepoExams
}

func ConnectRepo(r repoexams.RepoExams) ServiceExams {
	return &RepoExams{repo: r}
}

func (r *RepoExams) GetAllExams(ctx context.Context, Search, Sort string, Page, Perpage int) ([]response.Exams, int64, error) {
	entities, count, err := r.repo.GetAll(ctx, Search, Sort, Page, Perpage)
	if err != nil {
		return nil, 0, err
	}
	return response.ParseExams(entities), count, nil
}

func (r *RepoExams) FindExamsbyIdTeacher(ctx context.Context, Search, Sort string, Page, Perpage, id int) ([]response.Exams, int64, error) {
	entities, count, err := r.repo.FindByidTeacher(ctx, Search, Sort, Page, Perpage, id)
	if err != nil {
		return nil, 0, err
	}
	return response.ParseExams(entities), count, nil
}

func (r *RepoExams) CreateExams(ctx context.Context, data *request.Exams, role string, user_id, subject_id int) error {
	return r.repo.TxExams(ctx, func(repo repoexams.RepoExams) error {
		date, err := time.Parse("2006-01-02", data.Dates)
		if err != nil {
			return fmt.Errorf("invalid date format: %w", err)
		}

		idts, err := repo.CheckRoleforCreateOrUpdate(ctx, role, user_id, subject_id, data.TeacherId)
		if err != nil {
			return err
		}

		ent := &entity.Exams{
			NameExams:        data.NameExams,
			Dates:            date,
			StartLesson:      data.StartLesson,
			EndLesson:        data.EndLesson,
			TeacherSubjectId: idts,
			Model: helper.Models{
				CreatedAt: time.Now(),
			},
		}
		return repo.Create(ctx, ent)

	})

}

func (r *RepoExams) ShowExams(ctx context.Context, id, userid int, coderole string) (*response.Exams, error) {
	ent, err := r.repo.Show(ctx, id, userid, coderole)
	if err != nil {
		return nil, err
	}
	return r.mapExamToResponse(ent), nil
}

func (r *RepoExams) UpdateExam(ctx context.Context, data *request.Exams, role string, id, userid int) error {
	return r.repo.TxExams(ctx, func(repo repoexams.RepoExams) error {
		date, err := time.Parse("2006-01-02", data.Dates)
		if err != nil {
			return fmt.Errorf("invalid date format: %w", err)
		}

		tsid, err := repo.CheckRoleforCreateOrUpdate(ctx, role, userid, int(*data.SubjectId), data.TeacherId)
		if err != nil {
			return err
		}

		ent := &entity.Exams{
			NameExams:        data.NameExams,
			Dates:            date,
			StartLesson:      data.StartLesson,
			EndLesson:        data.EndLesson,
			TeacherSubjectId: tsid,
			Model: helper.Models{
				UpdatedAt: time.Now(),
			},
		}
		return repo.Update(ctx, id, ent)
	})
}

func (r *RepoExams) DeleteExam(ctx context.Context, id, userid int, coderole string) error {
	return r.repo.Delete(ctx, id, userid, coderole)
}

func (r *RepoExams) mapExamToResponse(v *entity.Exams) *response.Exams {
	if v.TeacherSubject == nil {
		v.TeacherSubject = &entity.TeacherSubjects{}
	}
	if v.TeacherSubject.Subject.IdSubject == 0 {
		v.TeacherSubject.Subject = entity.Subjects{}
	}
	return &response.Exams{
		IdExam:           v.IdExam,
		NameExams:        v.NameExams,
		Dates:            v.Dates,
		StartLesson:      v.StartLesson,
		EndLesson:        v.EndLesson,
		TeacherSubjectId: v.TeacherSubjectId,
		Subject: response.Subjects{
			IdSubject:   v.TeacherSubject.Subject.IdSubject,
			NameSubject: v.TeacherSubject.Subject.NameSubject,
			Semester:    v.TeacherSubject.Subject.Semester,
			Model:       v.TeacherSubject.Subject.Model,
		},
		Model: v.Model,
	}
}
