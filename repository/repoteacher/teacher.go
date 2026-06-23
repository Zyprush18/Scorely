package repoteacher

import (
	"context"
	"fmt"

	"github.com/Zyprush18/Scorely/models/entity"
	"gorm.io/gorm"
)

type RepoTeacher interface {
	GetAll(ctx context.Context, Search, Sort string, Page, Perpage int) ([]entity.Teachers, int64, error)
	Create(ctx context.Context, data *entity.Teachers) error
	Show(ctx context.Context, id int) (*entity.Teachers, error)
	Update(ctx context.Context, id int, data *entity.Teachers) error
	Delete(ctx context.Context, id int) error
}

type StructMysql struct {
	db *gorm.DB
}

func ConnectDb(d *gorm.DB) RepoTeacher {
	return &StructMysql{db: d}
}

func (m *StructMysql) GetAll(ctx context.Context, Search, Sort string, Page, Perpage int) ([]entity.Teachers, int64, error) {
	var finddata []entity.Teachers
	var count int64
	order := fmt.Sprintf("created_at %s", Sort)
	offset := (Page - 1) * Perpage

	if err := m.db.WithContext(ctx).Model(&entity.Teachers{}).Preload("Subject").Debug().Where("nip LIKE ?", "%"+Search+"%").Count(&count).Order(order).Limit(Perpage).Offset(offset).Find(&finddata).Error; err != nil {
		return nil, 0, err
	}

	return finddata, count, nil
}

func (m *StructMysql) Create(ctx context.Context, data *entity.Teachers) error {
	if err := m.db.WithContext(ctx).Model(&entity.Teachers{}).Debug().Create(data).Error; err != nil {
		return err
	}
	return nil
}
func (m *StructMysql) Show(ctx context.Context, id int) (*entity.Teachers, error) {
	var finddata entity.Teachers
	if err := m.db.WithContext(ctx).Model(&entity.Teachers{}).Preload("Subject").Where("id_teacher = ?", id).First(&finddata).Error; err != nil {
		return nil, err
	}
	return &finddata, nil
}

func (m *StructMysql) Update(ctx context.Context, id int, data *entity.Teachers) error {
	return m.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var teachermodel entity.Teachers
		if err := tx.Model(&teachermodel).Debug().Where("id_teacher = ?", id).First(&teachermodel).Error; err != nil {
			return err
		}

		result := tx.Model(&entity.Teachers{}).Debug().Where("id_teacher = ?", id).Updates(data)
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return gorm.ErrRecordNotFound
		}

		if err := tx.Model(&teachermodel).Association("Subject").Replace(data.Subject); err != nil {
			return err
		}

		return nil
	})
}

func (m *StructMysql) Delete(ctx context.Context, id int) error {
	result := m.db.WithContext(ctx).Model(&entity.Teachers{}).Delete(id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil

}
