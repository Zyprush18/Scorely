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
	GetAll(Search, Sort string, Page,Perpage int) ([]response.Teachers, int64, error)
	Create(data *request.Teachers) error
}

type StructMysql struct {
	db *gorm.DB
}

func ConnectDb(d *gorm.DB) StructMysql {
	return StructMysql{db: d}
}

func (m *StructMysql) GetAll(Search, Sort string, Page,Perpage int) ([]response.Teachers, int64, error)  {
	var finddata []entity.Teachers
	var count int64
	order := fmt.Sprintf("created_at %s",Sort)
	offset := (Page - 1) * Perpage

	if err := m.db.Model(&entity.Teachers{}).Preload("Subject").Debug().Where("nip LIKE ?", "%"+Search+"%").Count(&count).Order(order).Limit(Perpage).Offset(offset).Find(&finddata).Error; err != nil {
		return  nil, 0, err
	}
	

	return response.RespGetALl(finddata),count,nil
}

func (m *StructMysql) Create(data *request.Teachers) error {
	now := time.Now()

	// mengecek data di dalam subject apakah ada atau nggak
	var subjectfind []entity.Subjects
	if err := m.db.Model(&entity.Subjects{}).Where("id_subject In ?", data.SubjectId).Find(&subjectfind).Error; err != nil {
		return  err
	}

		teacherreq := &entity.Teachers{
			Name: data.Name,
			Nip: data.Nip,
			Gender: data.Gender,
			Address: data.Address,
			Phone: data.Phone,
			UserId: data.UserId,
			Subject: subjectfind,
			Models: helper.Models{
				CreatedAt: now,
			},
		}	


	if err:= m.db.Model("teachers").Create(teacherreq).Error;err != nil {
		return err
	}

	return nil
}