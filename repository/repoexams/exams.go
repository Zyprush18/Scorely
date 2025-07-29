package repoexams

import (
	"fmt"
	"time"

	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/models/entity"
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/models/response"
	"gorm.io/gorm"
)


type RepoExams interface {
	GetAll(Search,Sort string, Page,Perpage int) ([]response.Exams, int64,error)
	FindByidTeacher(Search,Sort string, Page,Perpage,id int) ([]response.Exams, int64,error)
	Create(data *request.Exams, role string,user_id,subject_id int) error
}

type MysqlStruct struct {
	db *gorm.DB
}

func ConnectDb(d *gorm.DB) MysqlStruct {
	return MysqlStruct{db: d}
}

func (m *MysqlStruct) GetAll(Search,Sort string, Page,Perpage int) ([]response.Exams, int64,error) {
	var modelexam []entity.Exams
	var count int64
	order := fmt.Sprintf("created_at %s",Sort)
	offset := (Page - 1) * Perpage

	if err:= m.db.Model(&entity.Exams{}).Debug().Preload("TeacherSubject.Subject").Where("name_exams LIKE ?", "%"+Search+"%").Count(&count).Order(order).Limit(Perpage).Offset(offset).Find(&modelexam).Error; err != nil {
		return nil,0,err
	}

	return response.ParseExams(modelexam),count,nil
}


// find exams for teacher
func (m *MysqlStruct) FindByidTeacher(Search,Sort string, Page,Perpage,id int) ([]response.Exams, int64,error)  {
	var finddata []entity.Exams
	var count int64
	order := fmt.Sprintf("created_at %s", Sort)
	offset := (Page - 1) * Perpage

	if err := m.db.Model(&entity.Exams{}).Debug().Preload("TeacherSubject.Subject").Debug().Joins("JOIN teacher_subjects AS ts ON ts.id_teacher_subject = exams.teacher_subject_id").Joins("JOIN subjects AS s ON s.id_subject = ts.id_subjects").Joins("JOIN teachers AS t ON t.id_teacher = ts.id_teachers").Where("t.user_id = ?", id).Where("name_exams LIKE ?", "%"+Search+"%").Count(&count).Order(order).Limit(Perpage).Offset(offset).Find(&finddata).Error; err != nil {
		return nil,0,err
	}

	return response.ParseExams(finddata),count,nil

}

func (m *MysqlStruct) Create(data *request.Exams, role string,user_id,subject_id int) error  {
	now := time.Now()
	idts, err := checkRole(m.db,role,user_id,subject_id,int(data.TeacherId))
	if err != nil {
		return err
	}

	req := &request.Exams{
		NameExams: data.NameExams,
		Dates: data.Dates,
		StartLesson: data.StartLesson,
		EndLesson: data.EndLesson,
		TeacherSubjectId: idts,
		Models: helper.Models{
			CreatedAt: now,
		},
	}

	if err:= m.db.Table("exams").Create(req).Error;err != nil {
		return  err
	}

	return  nil
}

func checkRole(d *gorm.DB,role string,user_id,subject_id,teacher_id int) (uint, error) {
	var model_teachersubject entity.TeacherSubjects
	query := d.Model(&entity.TeacherSubjects{})
	switch role {
	case "teacher":
		if err:= query.Debug().Joins("JOIN teachers AS t ON t.id_teacher = teacher_subjects.id_teachers").Where("t.user_id = ? AND id_subjects = ?", user_id, subject_id).First(model_teachersubject).Error;err != nil {
			return 0, err
		}

		return model_teachersubject.IdTeacherSubject,nil
	default:
		if err := query.Debug().Where("id_teachers = ? AND id_subjects = ?", teacher_id,subject_id).Find(model_teachersubject).Error;err != nil {
			return 0, err
		}

		return model_teachersubject.IdTeacherSubject,nil 
	}
}