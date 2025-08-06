package repoexamquestions

import (
	"fmt"
	"time"

	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/models/entity"
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/models/response"
	"gorm.io/gorm"
)

type RepoExamQuest interface {
	GetAll(Search,Sort,coderole string, Page,Perpage,user_id,idexam int) ([]response.Exam_Questions, int64,error)
	Create(data *request.Exam_Questions,userid,id_exam int,coderole string) error
	Show(id,user_id,exam_id int,coderole string) (*response.Exam_Questions,error)
}

type MysqlStruct struct {
	db *gorm.DB
}

func ConnectDB(d *gorm.DB) MysqlStruct {
	return MysqlStruct{db: d}
}

func (m *MysqlStruct) GetAll(Search,Sort,coderole string, Page,Perpage,user_id,idexam int) ([]response.Exam_Questions, int64,error) {
	var modelexamquestion []entity.Exam_Questions
	var count int64
	order := fmt.Sprintf("created_at %s", Sort)
	offset := (Page - 1) * Perpage
	query := CheckRole(m.db,user_id,idexam,coderole)

	if err:= query.Count(&count).Order(order).Limit(Perpage).Offset(offset).Find(&modelexamquestion).Error;err != nil {
		return nil,0, err
	}

	return response.ParseExamsQuest(modelexamquestion),count,nil
}

func (m *MysqlStruct) Create(data *request.Exam_Questions,userid,id_exam int,coderole string) error  {
	now := time.Now()
	req := &request.Exam_Questions{
		Question: data.Question,
		ExamId: uint(id_exam),
		Models: helper.Models{
			CreatedAt: now,
		},
	}


	query := CheckRole(m.db,userid,id_exam,coderole)
	if err:= query.First(&entity.Exam_Questions{}).Error;err != nil {
		return err
	}

	if err:= m.db.Debug().Table("exam_questions").Create(req).Error;err != nil {
		return err
	}

	return nil
}

func (m *MysqlStruct) Show(id,user_id,exam_id int,coderole string) (*response.Exam_Questions,error){
	var modelexamquestion entity.Exam_Questions
	query := CheckRole(m.db,user_id,exam_id,coderole)
	if err:= query.Where("id_exam_question = ?", id).First(&modelexamquestion).Error;err != nil {
		return nil,err
	}

	return &response.Exam_Questions{
		IdExamQuestion: modelexamquestion.IdExamQuestion,
		Question: modelexamquestion.Question,
		ExamId: modelexamquestion.ExamId,
		Exam: response.Exams{
			IdExam: modelexamquestion.Exam.IdExam,
			NameExams: modelexamquestion.Exam.NameExams,
			Dates: modelexamquestion.Exam.Dates,
			StartLesson: modelexamquestion.Exam.StartLesson,
			EndLesson: modelexamquestion.Exam.EndLesson,
			TeacherSubjectId: modelexamquestion.Exam.TeacherSubjectId,
			Subject: response.Subjects(modelexamquestion.Exam.TeacherSubject.Subject),
			Models: modelexamquestion.Exam.Models,
		},
		Models: modelexamquestion.Models,
	},nil
}


func CheckRole(d *gorm.DB,userid,idexam int,coderole string) *gorm.DB {
	query := d.Model(&entity.Exam_Questions{}).Preload("Exam.TeacherSubject.Subject").Debug()
	fmt.Println(coderole)
	switch coderole {
	case "teacher":
		return query.Joins("JOIN exams AS e ON e.id_exam = exam_questions.exam_id").Joins("JOIN teacher_subjects AS ts ON ts.id_teacher_subject = e.teacher_subject_id").Joins("JOIN teachers AS t ON t.id_teacher = ts.id_teachers").Where("t.user_id = ? AND e.id_exam = ?",userid,idexam)
	default:
		return query.Where("exam_id = ?", idexam)
	}
}
