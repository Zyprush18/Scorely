package repoexams

import (
	"fmt"

	"github.com/Zyprush18/Scorely/models/entity"
	"github.com/Zyprush18/Scorely/models/response"
	"gorm.io/gorm"
)


type RepoExams interface {
	GetAll(Search,Sort string, Page,Perpage int) ([]response.Exams, int64,error)
	FindByidTeacher(Search,Sort string, Page,Perpage,id int) ([]response.Exams, int64,error)

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

	if err:= m.db.Model(&entity.Exams{}).Preload("TeacherSubject.Subject").Where("name_exams LIKE ?", "%"+Search+"%").Count(&count).Order(order).Limit(Perpage).Offset(offset).Find(&modelexam).Error; err != nil {
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

	if err := m.db.Model(&entity.Exams{}).Preload("TeacherSubject.Subject").Debug().Joins("JOIN teacher_subjects AS ts ON ts.id_teacher_subject = exams.teacher_subject_id").Joins("JOIN subjects AS s ON s.id_subject = ts.id_subjects").Joins("JOIN teachers AS t ON t.id_teacher = ts.id_teachers").Where("t.user_id = ?", id).Where("name_exams LIKE ?", "%"+Search+"%").Count(&count).Order(order).Limit(Perpage).Offset(offset).Find(&finddata).Error; err != nil {
		return nil,0,err
	}

	return response.ParseExams(finddata),count,nil

}