package serviceexamquest

import (
	"context"
	"time"

	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/models/entity"
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/models/response"
	"github.com/Zyprush18/Scorely/repository/repoexamquestions"
)

type ServiceExamQuest interface {
	GetAllExamQuest(ctx context.Context, Search, Sort, coderole string, Page, Perpage, user_id, idexam int) ([]response.Exam_Questions, int64, error)
	CreateExamQuest(ctx context.Context, data *request.Exam_Questions, userid, id_exam int, coderole string) error
	ShowExamQuest(ctx context.Context, id, user_id, exam_id int, coderole string) (*response.Exam_Questions, error)
}

type RepoExamQuest struct {
	repo repoexamquestions.RepoExamQuest
}

func ConnectRepo(r repoexamquestions.RepoExamQuest) ServiceExamQuest {
	return &RepoExamQuest{repo: r}
}

func (r *RepoExamQuest) GetAllExamQuest(ctx context.Context, Search, Sort, coderole string, Page, Perpage, user_id, idexam int) ([]response.Exam_Questions, int64, error) {
	entities, count, err := r.repo.GetAll(ctx, Search, Sort, coderole, Page, Perpage, user_id, idexam)
	if err != nil {
		return nil, 0, err
	}
	return response.ParseExamsQuest(entities), count, nil
}

func (r *RepoExamQuest) CreateExamQuest(ctx context.Context, data *request.Exam_Questions, userid, id_exam int, coderole string) error {
	if err := r.repo.VerifyExamAccess(ctx, userid, id_exam, coderole); err != nil {
		return err
	}

	ent := &entity.Exam_Questions{
		Question: data.Question,
		ExamId:   uint(id_exam),
		Model: helper.Models{
			CreatedAt: time.Now(),
		},
	}
	return r.repo.Create(ctx, ent)
}

func (r *RepoExamQuest) ShowExamQuest(ctx context.Context, id, user_id, exam_id int, coderole string) (*response.Exam_Questions, error) {
	ent, err := r.repo.Show(ctx, id, user_id, exam_id, coderole)
	if err != nil {
		return nil, err
	}
	return mapExamQuestToResponse(ent), nil
}

func mapExamQuestToResponse(v *entity.Exam_Questions) *response.Exam_Questions {
	if v.Exam.TeacherSubject == nil {
		v.Exam.TeacherSubject = &entity.TeacherSubjects{}
	}
	return &response.Exam_Questions{
		IdExamQuestion: v.IdExamQuestion,
		Question:       v.Question,
		ExamId:         v.ExamId,
		Exam: response.Exams{
			IdExam:            v.Exam.IdExam,
			NameExams:         v.Exam.NameExams,
			Dates:             v.Exam.Dates,
			StartLesson:       v.Exam.StartLesson,
			EndLesson:         v.Exam.EndLesson,
			TeacherSubjectId:  v.Exam.TeacherSubjectId,
			Subject:           response.Subjects(v.Exam.TeacherSubject.Subject),
			Model:            v.Exam.Model,
		},
		Model: v.Model,
	}
}