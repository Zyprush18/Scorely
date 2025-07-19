package reposubject

import (
	"fmt"
	"time"

	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/models/entity"
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/models/response"
	"gorm.io/gorm"
)

type RepoSubject interface {
	GetAll(Search, Sort string,Page,Perpage int)([]response.Subjects, int64, error)
	Create(data *request.Subjects) error
	Show(id int) (*response.Subjects, error)
	Update(id int, data *request.Subjects) error
	Delete(id int) error
}

type MysqlStruct struct {
	db *gorm.DB
}

func ConnectDb(d *gorm.DB) MysqlStruct  {
	return MysqlStruct{db: d}
}

func (m *MysqlStruct) GetAll(Search, Sort string,Page,Perpage int)([]response.Subjects, int64, error) {
	var subjectmodel []entity.Subjects
	var count int64
	order := fmt.Sprintf("created_at %s", Sort)
	offset := (Page - 1) * Perpage

	if err := m.db.Model(&entity.Subjects{}).Debug().Where("name_subject LIKE ?", "%"+Search+"%").Count(&count).Order(order).Limit(Perpage).Offset(offset).Find(&subjectmodel).Error;err!= nil {
		return nil,0,err
	}

	return response.Subjectsresp(subjectmodel),count,nil
}

func (m *MysqlStruct) Create(data *request.Subjects) error {
	now := time.Now()
	reqsubject := &request.Subjects{
		NameSubject: data.NameSubject,
		Semester: data.Semester,
		Models: helper.Models{
			CreatedAt: now,
		},
	}

	if err := m.db.Table("subjects").Debug().Create(reqsubject).Error; err != nil {
		return err
	}

	return nil
}

func (m *MysqlStruct) Show(id int) (*response.Subjects, error) {
	var finddata entity.Subjects

	if err := m.db.Model(&finddata).Debug().Where("id_subject = ?", id).First(&finddata).Error; err != nil {
		return nil, err
	}

	return &response.Subjects{
		IdSubject: finddata.IdSubject,
		NameSubject: finddata.NameSubject,
		Semester: finddata.Semester,
		Models: finddata.Models,
	},nil
}

func (m *MysqlStruct) Update(id int, data *request.Subjects) error {
	var finddata response.Subjects
	now := time.Now()
	subjectreq := &request.Subjects{
		NameSubject: data.NameSubject,
		Semester: data.Semester,
		Models: helper.Models{
			UpdatedAt: now,
		},
	}

	if err:= m.db.Table("subjects").Debug().Where("id_subject = ?", id).First(&finddata).Updates(subjectreq).Error;err != nil {
		return err
	}

	return nil
}

func (m *MysqlStruct) Delete(id int) error {
	var finddata entity.Subjects
	if err:= m.db.Table("teacher_subjects").Where("id_subject = ?", id).Debug().Delete("teacher_subjects").Error;err != nil {
		return err
	}
	if err := m.db.Table("subjects").Where("id_subject = ?",id).First(&finddata).Delete(finddata).Error;err != nil {
		return err
	}

	return nil
}