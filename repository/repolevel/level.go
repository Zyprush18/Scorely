package repolevel

import (
	"fmt"
	"time"

	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/models/entity"
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/models/response"
	"gorm.io/gorm"
)

type LevelRepo interface {
	GetAll(search, sort string, page,perpage int) ([]response.Levels, int64,error)
	Create(data *request.Levels) error
	Show(id int) (*response.Levels, error)
	Update(id int, data *request.Levels) error
	Delete(id int) error
}

type MysqlStruct struct {
	db *gorm.DB
}

func ConnectDb(d *gorm.DB) MysqlStruct  {
	return MysqlStruct{db: d}
}

func (m *MysqlStruct) GetAll(search, sort string, page,perpage int) ([]response.Levels, int64,error) {
	var model_level []response.Levels
	var count int64
	order := fmt.Sprintf("created_at %s", sort)
	offset := (page - 1) * perpage
	if err:= m.db.Table("levels").Debug().Where("level LIKE ?", "%"+search+"%").Count(&count).Order(order).Limit(perpage).Offset(offset).Find(&model_level).Error;err != nil {
		return nil, 0, err
	}

	return model_level,count, nil
}

func (m *MysqlStruct) Create(data *request.Levels) error {
	levelreq:= &request.Levels{
		Level: data.Level,
		Models: helper.Models{
			CreatedAt: time.Now(),
		},
	}

	if err:= m.db.Table("levels").Debug().Create(levelreq).Error;err != nil {
		return err
	}

	return  nil
}

func (m *MysqlStruct) Show(id int) (*response.Levels, error) {
	var model_level response.Levels
	if err:=m.db.Table("levels").Debug().Where("id_level = ?", id).First(&model_level).Error;err != nil {
		return nil, err
	}

	return &model_level,nil
}

func (m *MysqlStruct) Update(id int, data *request.Levels) error {
	var model_level response.Levels
	now := time.Now()
	levelreq:=&request.Levels{
		Level: data.Level,
		Models: helper.Models{
			UpdatedAt: now,
		},
	}
	if err:= m.db.Table("levels").Where("id_level = ?", id).First(&model_level).Error;err != nil {
		return  err
	}

	if err:= m.db.Table("levels").Where("id_level = ?", id).Updates(levelreq).Error;err != nil {
		return err
	}

	return nil
}

func (m *MysqlStruct) Delete(id int) error {
	var model_level entity.Levels
	if err:= m.db.Table("levels").Where("id_level = ?", id).First(&model_level).Error;err != nil {
		return  err
	}

	if err:=m.db.Delete(&model_level).Error;err != nil {
		return  err
	}
	return  nil
}
