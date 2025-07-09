package repomajor

import (
	"fmt"
	"time"

	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/models/entity"
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/models/response"
	"gorm.io/gorm"
)

type MajorRepo interface {
	GetAllData(search, sort string, page, perpage int) ([]response.Majors, int64, error)
	Create(data *request.Majors) error
	ShowById(id int) (*response.Majors, error)
	Updates(id int, data *request.Majors) error
	Deletes(id int) error
}

type MysqlStruct struct {
	db *gorm.DB
}

func ConnectDb(db *gorm.DB) MysqlStruct {
	return MysqlStruct{db: db}
}

func (m *MysqlStruct) GetAllData(search, sort string, page, perpage int) ([]response.Majors, int64, error) {
	var modelmajor []response.Majors
	var count int64
	offset := (page - 1) * perpage
	order := fmt.Sprintf("created_at %s", sort)

	if err := m.db.Table("majors").Where("major LIKE ?", "%"+search+"%").Count(&count).Order(order).Limit(perpage).Offset(offset).Find(&modelmajor).Error; err != nil {
		return nil, 0, err
	}

	return modelmajor, count, nil
}

func (m *MysqlStruct) Create(data *request.Majors) error {
	now := time.Now()
	major := &request.Majors{
		Major:             data.Major,
		MajorAbbreviation: data.MajorAbbreviation,
		Models: helper.Models{
			CreatedAt: now,
		},
	}

	if err := m.db.Table("majors").Create(&major).Error; err != nil {
		return err
	}

	return nil
}

func (m *MysqlStruct) ShowById(id int) (*response.Majors, error) {
	var modelmajor response.Majors
	if err := m.db.Table("majors").Where("id_major = ?", id).First(&modelmajor).Error; err != nil {
		return nil, err
	}

	return &modelmajor, nil
}

func (m *MysqlStruct) Updates(id int, data *request.Majors) error {
	var modelmajor entity.Majors

	if err := m.db.Table("majors").Where("id_major = ?", id).First(&modelmajor).Error; err != nil {
		return err
	}

	if err := m.db.Table("majors").Where("id_major = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

func (m *MysqlStruct) Deletes(id int) error {
	var modelmajor entity.Majors

	if err := m.db.Model(&modelmajor).Where("id_major = ?", id).First(&modelmajor).Error; err != nil {
		return err
	}

	if err := m.db.Delete(&modelmajor).Error; err != nil {
		return err
	}

	return nil
}
