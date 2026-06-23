package repoexamquestions

import (
	"context"
	"fmt"

	"github.com/Zyprush18/Scorely/models/entity"
	"gorm.io/gorm"
)

type RepoExamQuest interface {
	GetAll(ctx context.Context, Search, Sort, coderole string, Page, Perpage, user_id, idexam int) ([]entity.Exam_Questions, int64, error)
	Create(ctx context.Context, data *entity.Exam_Questions) error
	Show(ctx context.Context, id, user_id, exam_id int, coderole string) (*entity.Exam_Questions, error)
	VerifyExamAccess(ctx context.Context, userid, examId int, coderole string) error
}

type MysqlStruct struct {
	db *gorm.DB
}

func ConnectDB(d *gorm.DB) RepoExamQuest {
	return &MysqlStruct{db: d}
}

func (m *MysqlStruct) GetAll(ctx context.Context, Search, Sort, coderole string, Page, Perpage, user_id, idexam int) ([]entity.Exam_Questions, int64, error) {
	var modelexamquestion []entity.Exam_Questions
	var count int64
	order := fmt.Sprintf("created_at %s", Sort)
	offset := (Page - 1) * Perpage
	query := m.CheckRole(m.db.WithContext(ctx), user_id, idexam, coderole)

	if err := query.Count(&count).Order(order).Limit(Perpage).Offset(offset).Find(&modelexamquestion).Error; err != nil {
		return nil, 0, err
	}

	return modelexamquestion, count, nil
}

func (m *MysqlStruct) VerifyExamAccess(ctx context.Context, userid, examId int, coderole string) error {
	query := m.CheckRole(m.db.WithContext(ctx), userid, examId, coderole)
	return query.First(&entity.Exam_Questions{}).Error
}

func (m *MysqlStruct) Create(ctx context.Context, data *entity.Exam_Questions) error {
	if err := m.db.WithContext(ctx).Debug().Table("exam_questions").Create(data).Error; err != nil {
		return err
	}
	return nil
}

func (m *MysqlStruct) Show(ctx context.Context, id, user_id, exam_id int, coderole string) (*entity.Exam_Questions, error) {
	var modelexamquestion entity.Exam_Questions
	query := m.CheckRole(m.db.WithContext(ctx), user_id, exam_id, coderole)
	if err := query.Where("id_exam_question = ?", id).First(&modelexamquestion).Error; err != nil {
		return nil, err
	}
	return &modelexamquestion, nil
}

func (m *MysqlStruct) CheckRole(d *gorm.DB, userid, idexam int, coderole string) *gorm.DB {
	query := d.Model(&entity.Exam_Questions{}).Preload("Exam.TeacherSubject.Subject").Debug()
	switch coderole {
	case "teacher":
		return query.Joins("JOIN exams AS e ON e.id_exam = exam_questions.exam_id").Joins("JOIN teacher_subjects AS ts ON ts.id_teacher_subject = e.teacher_subject_id").Joins("JOIN teachers AS t ON t.id_teacher = ts.id_teachers").Where("t.user_id = ? AND e.id_exam = ?", userid, idexam)
	default:
		return query.Where("exam_id = ?", idexam)
	}
}
