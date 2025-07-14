package repoteacher

import (
	"fmt"

	"github.com/Zyprush18/Scorely/models/entity"
	"github.com/Zyprush18/Scorely/models/response"
	"gorm.io/gorm"
)

type RepoTeacher interface {
	GetAll(Search, Sort string, Page,Perpage int) ([]response.Teachers, int64, error)
}

type StructMysql struct {
	db *gorm.DB
}

func ConnectDb(d *gorm.DB) StructMysql {
	return StructMysql{db: d}
}

func (m *StructMysql) GetAll(Search, Sort string, Page,Perpage int) ([]response.Teachers, int64, error)  {
	var finddata []response.Teachers
	var count int64
	order := fmt.Sprintf("created_at %s",Sort)
	offser := (Page - 1) * Perpage

	if err := m.db.Model(&entity.Teachers{}).Debug().Where("name = ?", "%"+Search+"%").Count(&count).Order(order).Limit(Perpage).Offset(offser).Find(&finddata).Error; err != nil {
		return  nil, 0, err
	}

	return finddata,count,nil
}