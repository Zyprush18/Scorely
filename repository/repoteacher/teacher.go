package repoteacher

import (
	"fmt"
	"time"

	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/models/entity"
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/models/response"
	"gorm.io/gorm"
)

type RepoTeacher interface {
	GetAll(Search, Sort string, Page, Perpage int) ([]response.Teachers, int64, error)
	Create(data *request.Teachers) error
	Show(id int) (*response.Teachers, error)
	Update(id int, data *request.Teachers) error
	Delete(id int) error
}

type StructMysql struct {
	db *gorm.DB
}

func ConnectDb(d *gorm.DB) StructMysql {
	return StructMysql{db: d}
}

func (m *StructMysql) GetAll(Search, Sort string, Page, Perpage int) ([]response.Teachers, int64, error) {
	var finddata []entity.Teachers
	var count int64
	order := fmt.Sprintf("created_at %s", Sort)
	offset := (Page - 1) * Perpage

	if err := m.db.Model(&entity.Teachers{}).Preload("Subject").Debug().Where("nip LIKE ?", "%"+Search+"%").Count(&count).Order(order).Limit(Perpage).Offset(offset).Find(&finddata).Error; err != nil {
		return nil, 0, err
	}

	return response.RespGetALl(finddata), count, nil
}

func (m *StructMysql) Create(data *request.Teachers) error {
	now := time.Now()

	// mengecek data di dalam subject apakah ada atau nggak
	var subjectfind []entity.Subjects
	if err := m.db.Model(&entity.Subjects{}).Debug().Where("id_subject IN ?", data.SubjectId).Find(&subjectfind).Error; err != nil {
		return err
	}

	teacherreq := &entity.Teachers{
		Name:    data.Name,
		Nip:     data.Nip,
		Gender:  data.Gender,
		Address: data.Address,
		Phone:   data.Phone,
		UserId:  data.UserId,
		Subject: subjectfind,
		Models: helper.Models{
			CreatedAt: now,
		},
	}

	if err := m.db.Model(&entity.Teachers{}).Debug().Create(teacherreq).Error; err != nil {
		return err
	}

	return nil
}
func (m *StructMysql) Show(id int) (*response.Teachers, error) {
	var finddata entity.Teachers
	if err := m.db.Model(&finddata).Preload("Subject").Where("id_teacher =?", id).First(&finddata).Error; err != nil {
		return nil, err
	}

	return &response.Teachers{
		IdTeacher: finddata.IdTeacher,
		Name:      finddata.Name,
		Nip:       finddata.Nip,
		Gender:    finddata.Gender,
		Address:   finddata.Address,
		Phone:     finddata.UserId,
		UserId:    finddata.UserId,
		Subject:   response.Subjectsresp(finddata.Subject),
		Models:    finddata.Models,
	}, nil
}

func (m *StructMysql) Update(id int, data *request.Teachers) error {
	now := time.Now()

	var findsubject []entity.Subjects
	if err := m.db.Model(&entity.Subjects{}).Debug().Where("id_subject IN ?", data.SubjectId).Find(&findsubject).Error; err != nil {
		return err
	}

	teacherreq := &entity.Teachers{
		Name:    data.Name,
		Nip:     data.Nip,
		Gender:  data.Gender,
		Address: data.Address,
		Phone:   data.Phone,
		UserId:  data.UserId,
		Subject: findsubject,
		Models: helper.Models{
			UpdatedAt: now,
		},
	}

	// update table teachers
	var teachermodel entity.Teachers
	if err := m.db.Model(&teachermodel).Debug().Where("id_teacher = ?", id).First(&teachermodel).Updates(teacherreq).Error; err != nil {
		return err
	}

	// update di relasi many to many nya
	if err := m.db.Model(&teachermodel).Association("Subject").Replace(findsubject); err != nil {
		return err
	}

	return nil
}

func (m *StructMysql) Delete(id int) error {
	var findteacher entity.Teachers

	// mengambil data teacher dan subject
	if err := m.db.Preload("Subject").First(&findteacher, id).Error; err != nil {
		return err
	}

	// menghapus relasi ke subject
	if err := m.db.Model(&findteacher).Association("Subject").Clear(); err != nil {
		return err
	}

	// menghapus data teacher nya
	if err := m.db.Delete(&findteacher).Error; err != nil {
		return err
	}

	return nil

}
