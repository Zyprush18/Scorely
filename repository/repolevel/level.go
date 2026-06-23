package repolevel

import (
	"context"
	"fmt"

	"github.com/Zyprush18/Scorely/models/entity"
	"gorm.io/gorm"
)

type LevelRepo interface {
	GetAll(ctx context.Context, search, sort string, page, perpage int) ([]entity.Levels, int64, error)
	Create(ctx context.Context, data *entity.Levels) error
	Show(ctx context.Context, id int) (*entity.Levels, error)
	Update(ctx context.Context, id int, data *entity.Levels) error
	Delete(ctx context.Context, id int) error
}

type MysqlStruct struct {
	db *gorm.DB
}

func ConnectDb(d *gorm.DB) LevelRepo {
	return &MysqlStruct{db: d}
}

func (m *MysqlStruct) GetAll(ctx context.Context, search, sort string, page, perpage int) ([]entity.Levels, int64, error) {
	var model_level []entity.Levels
	var count int64
	order := fmt.Sprintf("created_at %s", sort)
	offset := (page - 1) * perpage
	if err := m.db.WithContext(ctx).Model(&entity.Levels{}).Debug().Where("level LIKE ?", "%"+search+"%").Count(&count).Order(order).Limit(perpage).Offset(offset).Find(&model_level).Error; err != nil {
		return nil, 0, err
	}

	return model_level, count, nil
}

func (m *MysqlStruct) Create(ctx context.Context, data *entity.Levels) error {
	if err := m.db.WithContext(ctx).Model(&entity.Levels{}).Debug().Create(data).Error; err != nil {
		return err
	}
	return nil
}

func (m *MysqlStruct) Show(ctx context.Context, id int) (*entity.Levels, error) {
	var model_level entity.Levels
	if err := m.db.WithContext(ctx).Model(&entity.Levels{}).Debug().Where("id_level = ?", id).First(&model_level).Error; err != nil {
		return nil, err
	}
	return &model_level, nil
}

func (m *MysqlStruct) Update(ctx context.Context, id int, data *entity.Levels) error {
	result := m.db.WithContext(ctx).Model(&entity.Levels{}).Where("id_level = ?", id).Updates(data)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (m *MysqlStruct) Delete(ctx context.Context, id int) error {
	result := m.db.WithContext(ctx).Model(&entity.Levels{}).Delete(id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
