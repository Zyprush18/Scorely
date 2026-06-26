package repostudent

import (
	"context"
	"fmt"

	"github.com/Zyprush18/Scorely/models/entity"
	"gorm.io/gorm"
)

type StudentRepo interface {
	GetAll(ctx context.Context, Search, Sort string, Page, Perpage int) ([]entity.Students, int64, error)
	Create(ctx context.Context, data *entity.Students) error
	Show(ctx context.Context, id int) (*entity.Students, error)
	Update(ctx context.Context, id int, data *entity.Students) error
	Delete(ctx context.Context, id int) error
}

type MysqlStruct struct {
	db *gorm.DB
}

func ConnectDb(d *gorm.DB) StudentRepo {
	return &MysqlStruct{db: d}
}

func (m *MysqlStruct) GetAll(ctx context.Context, Search, Sort string, Page, Perpage int) ([]entity.Students, int64, error) {
	var findstudent []entity.Students
	var count int64
	order := fmt.Sprintf("created_at %s", Sort)
	offset := (Page - 1) * Perpage

	if err := m.db.WithContext(ctx).Model(&entity.Students{}).Preload("Class.Major").Preload("Class.Level").Debug().Where("nisn LIKE ?", "%"+Search+"%").Count(&count).Order(order).Limit(Perpage).Offset(offset).Find(&findstudent).Error; err != nil {
		return nil, 0, err
	}

	return findstudent, count, nil
}

func (m *MysqlStruct) Create(ctx context.Context, data *entity.Students) error {
	if err := m.db.WithContext(ctx).Model(&entity.Students{}).Create(data).Error; err != nil {
		return err
	}
	return nil
}

func (m *MysqlStruct) Show(ctx context.Context, id int) (*entity.Students, error) {
	var findstudent entity.Students
	if err := m.db.WithContext(ctx).Model(&entity.Students{}).Preload("Class.Major").Preload("Class.Level").Where("id_student = ?", id).First(&findstudent).Error; err != nil {
		return nil, err
	}
	return &findstudent, nil
}

func (m *MysqlStruct) Update(ctx context.Context, id int, data *entity.Students) error {
	result := m.db.WithContext(ctx).Model(&entity.Students{}).Where("id_student = ?", id).Updates(data)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (m *MysqlStruct) Delete(ctx context.Context, id int) error {
	result := m.db.WithContext(ctx).Model(&entity.Students{}).Delete(id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}