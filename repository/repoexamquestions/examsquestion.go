package repoexamquestions

import (
	"fmt"

	"github.com/Zyprush18/Scorely/models/entity"
	"github.com/Zyprush18/Scorely/models/response"
	"gorm.io/gorm"
)

type RepoExamQuest interface {
	GetAll(Search,Sort,coderole string, Page,Perpage,user_id,idexam int) ([]response.Exam_Questions, int64,error)
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


func CheckRole(d *gorm.DB,userid,idexam int,coderole string) *gorm.DB {
	query := d.Model(&entity.Exam_Questions{}).Preload("Exam")
	switch coderole {
	case "teacher":
		return query.Joins("JOIN exams AS e ON e.id_exam = exam_questions.exam_id").Joins("JOIN teacher_subjects AS ts ON ts.id_teacher_subject = e.teacher_subject_id").Joins("JOIN teachers AS t ON t.id_teacher = ts.id_teachers").Where("t.user_id = ? AND exam_id = ?",userid,idexam)
	default:
		return query.Where("exam_id = ?", idexam)
	}
}
