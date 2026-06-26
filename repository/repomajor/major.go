package repomajor

import (
	"context"
	"fmt"

	"github.com/Zyprush18/Scorely/models/entity"
	"github.com/Zyprush18/Scorely/models/request"
	"gorm.io/gorm"
)

type MajorRepo interface {
	GetAllData(ctx context.Context, search, sort string, page, perpage int) ([]entity.Majors, int64, error)
	Create(ctx context.Context, data *request.Majors) error
	ShowById(ctx context.Context, id int) (*entity.Majors, error)
	Updates(ctx context.Context, id int, data *request.Majors) error
	Deletes(ctx context.Context, id int) error
}

type MysqlStruct struct {
	db *gorm.DB
}

func ConnectDb(db *gorm.DB) MajorRepo {
	return &MysqlStruct{db: db}
}


func (m *MysqlStruct) GetAllData(ctx context.Context, search, sort string, page, perpage int) ([]entity.Majors, int64, error) {
	var modelmajor []entity.Majors
	var count int64
	offset := (page - 1) * perpage
	order := fmt.Sprintf("created_at %s", sort)

	if err := m.db.WithContext(ctx).Model(&entity.Majors{}).Where("major LIKE ?", "%"+search+"%").Count(&count).Order(order).Limit(perpage).Offset(offset).Find(&modelmajor).Error; err != nil {
		return nil, 0, err
	}

	return modelmajor, count, nil
}

func (m *MysqlStruct) Create(ctx context.Context, data *request.Majors) error {
	if err := m.db.WithContext(ctx).Model(&entity.Majors{}).Create(data).Error; err != nil {
		return err
	}
	return nil
}

func (m *MysqlStruct) ShowById(ctx context.Context, id int) (*entity.Majors, error) {
	var modelmajor entity.Majors
	if err := m.db.WithContext(ctx).Model(&entity.Majors{}).Where("id_major = ?", id).First(&modelmajor).Error; err != nil {
		return nil, err
	}
	return &modelmajor, nil
}

func (m *MysqlStruct) Updates(ctx context.Context, id int, data *request.Majors) error {
	result := m.db.WithContext(ctx).Model(&entity.Majors{}).Where("id_major = ?", id).Updates(data)
	if result.Error != nil {
		return  result.Error
	}


	if result.RowsAffected == 0 {
		return  gorm.ErrRecordNotFound
	}

	return nil
}

func (m *MysqlStruct) Deletes(ctx context.Context, id int) error {
	result :=  m.db.Model(&entity.Majors{}).Delete(id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return  gorm.ErrRecordNotFound
	}

	return nil
}
