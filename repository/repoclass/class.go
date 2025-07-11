package repoclass

import (
	"fmt"
	"time"

	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/models/entity"
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/models/response"
	"gorm.io/gorm"
)

type RepoClass interface {
	GetAll(Search,Sort string, Page,Perpage int) ([]response.Class, int64, error)
	Create(data *request.Class) error
	Show(id int)(*response.Class, error)
	Update(id int, data *request.Class) error
	Delete(id int) error
}

type MysqlStruct struct {
	db *gorm.DB
}

func ConnectDb(d *gorm.DB) MysqlStruct  {
	return MysqlStruct{db: d,}
}

func (m *MysqlStruct) GetAll(Search,Sort string, Page,Perpage int) ([]response.Class, int64, error)  {
	var findclass []entity.Class
	var count int64
	order := fmt.Sprintf("created_at %s", Sort)
	offset := (Page - 1) * Perpage

	// query := m.db.Table("classes")
	if err := m.db.Model(&entity.Class{}).Preload("Level").Preload("Major").Debug().Where("name LIKE ?", "%"+Search+"%").Count(&count).Order(order).Limit(Perpage).Offset(offset).Find(&findclass).Error;err != nil {
		return nil,0,err
	}


	return ParseResponse(findclass),count,nil

}

func (m *MysqlStruct) Create(data *request.Class) error {
	now := time.Now()
	classreq := &request.Class{
		Name: data.Name,
		LevelId: data.LevelId,
		MajorId: data.MajorId,
		Models: helper.Models{
			CreatedAt: now,
		},
	}

	if err := m.db.Table("classes").Debug().Create(classreq).Error;err != nil {
		return err
	}

	return nil
}

func (m *MysqlStruct) Show(id int) (*response.Class, error)  {
	var find_class entity.Class
	if err:= m.db.Model(&entity.Class{}).Preload("Level").Preload("Major").Debug().Where("id_class = ?", id).First(&find_class).Error;err != nil {
		return nil, err
	}
	resp := &response.Class{
		IdClass: find_class.IdClass,
		Name: find_class.Name,
		LevelId: find_class.LevelId,
		MajorId: find_class.MajorId,
		Level: response.Levels{
			IdLevel: find_class.Level.IdLevel,
			Level: find_class.Level.Level,
			Models: find_class.Models,
		},
		Major: response.Majors{
			IdMajor: find_class.Major.IdMajor,
			Major: find_class.Major.Major,
			MajorAbbreviation: find_class.Major.MajorAbbreviation,
			Models: find_class.Major.Models,
		},
		Models: find_class.Models,
	}

	return resp, nil
}

func (m *MysqlStruct) Update(id int, data *request.Class) error {
	now := time.Now()
	classreq := &request.Class{
		Name: data.Name,
		LevelId: data.LevelId,
		MajorId: data.MajorId,
		Models: helper.Models{
			UpdatedAt: now,
		},
	}

	if err := m.db.Table("classes").Debug().Where("id_class = ?", id).First(&entity.Class{}).Updates(classreq).Error;err != nil {
		return err
	}

	return nil
}

func (m *MysqlStruct) Delete(id int) error {
	var modelclass entity.Class
	if err:= m.db.Table("classes").Debug().Where("id_class = ?", id).First(&modelclass).Delete(&modelclass).Error;err != nil {
		return err
	}
	return nil
}

func ParseResponse(c []entity.Class) (resp []response.Class)  {
	for _, v := range c {
		resp = append(resp, response.Class{
			IdClass: v.IdClass,
			Name: v.Name,
			LevelId: v.LevelId,
			MajorId: v.MajorId,
			Level: response.Levels{
				IdLevel: v.LevelId,
				Level: v.Level.Level,
				Models: v.Level.Models,
			},
			Major: response.Majors{
				IdMajor: v.MajorId,
				Major: v.Major.Major,
				MajorAbbreviation: v.Major.MajorAbbreviation,
				Models: v.Major.Models,
			},
			Models: v.Models,
		})
	}


	return resp
}