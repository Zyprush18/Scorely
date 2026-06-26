package repoclass

import (
	"context"
	"fmt"

	"github.com/Zyprush18/Scorely/models/entity"
	"gorm.io/gorm"
)

type RepoClass interface {
	GetAll(ctx context.Context, search, sort string, page, perpage int) ([]entity.Class, int64, error)
	Create(ctx context.Context, data *entity.Class) error
	Show(ctx context.Context, id int) (*entity.Class, error)
	Update(ctx context.Context, id int, data *entity.Class) error
	Delete(ctx context.Context, id int) error
}

type MysqlStruct struct {
	db *gorm.DB
}

func ConnectDb(d *gorm.DB) RepoClass {
	return &MysqlStruct{db: d}
}

func (m *MysqlStruct) GetAll(ctx context.Context, search, sort string, page, perpage int) ([]entity.Class, int64, error) {
	var findclass []entity.Class
	var count int64
	order := fmt.Sprintf("created_at %s", sort)
	offset := (page - 1) * perpage

	if err := m.db.WithContext(ctx).Model(&entity.Class{}).Preload("Level").Preload("Major").Debug().Where("name LIKE ?", "%"+search+"%").Count(&count).Order(order).Limit(perpage).Offset(offset).Find(&findclass).Error; err != nil {
		return nil, 0, err
	}

	return findclass, count, nil
}

func (m *MysqlStruct) Create(ctx context.Context, data *entity.Class) error {
	if err := m.db.WithContext(ctx).Model(&entity.Class{}).Debug().Create(data).Error; err != nil {
		return err
	}
	return nil
}

func (m *MysqlStruct) Show(ctx context.Context, id int) (*entity.Class, error) {
	var find_class entity.Class
	if err := m.db.WithContext(ctx).Model(&entity.Class{}).Preload("Level").Preload("Major").Debug().Where("id_class = ?", id).First(&find_class).Error; err != nil {
		return nil, err
	}
	return &find_class, nil
}

func (m *MysqlStruct) Update(ctx context.Context, id int, data *entity.Class) error {
	result := m.db.WithContext(ctx).Model(&entity.Class{}).Debug().Where("id_class = ?", id).Updates(data)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (m *MysqlStruct) Delete(ctx context.Context, id int) error {
	result := m.db.WithContext(ctx).Model(&entity.Class{}).Debug().Delete(id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}