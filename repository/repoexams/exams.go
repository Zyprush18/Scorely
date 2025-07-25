package repoexams

import (
	"fmt"

	"github.com/Zyprush18/Scorely/models/entity"
	"github.com/Zyprush18/Scorely/models/response"
	"gorm.io/gorm"
)


type RepoExams interface {
	GetAll(Search,Sort string, Page,Perpage int) ([]response.Exams, int64,error)
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

	if err:= m.db.Model(&entity.Exams{}).Preload("Subject").Where("name_exams LIKE ?", "%"+Search+"%").Count(&count).Order(order).Limit(Perpage).Offset(offset).Find(&modelexam).Error; err != nil {
		return nil,0,err
	}

	return response.ParseExams(modelexam),count,nil
}