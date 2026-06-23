package reposubject

import (
	"context"
	"fmt"

	"github.com/Zyprush18/Scorely/models/entity"
	"gorm.io/gorm"
)

type RepoSubject interface {
	GetAll(ctx context.Context, Search, Sort string, Page, Perpage int) ([]entity.Subjects, int64, error)
	Create(ctx context.Context, data *entity.Subjects) error
	Show(ctx context.Context, id int) (*entity.Subjects, error)
	Update(ctx context.Context, id int, data *entity.Subjects) error
	Delete(ctx context.Context, id int) error
}

type MysqlStruct struct {
	db *gorm.DB
}

func ConnectDb(d *gorm.DB) RepoSubject {
	return &MysqlStruct{db: d}
}

func (m *MysqlStruct) GetAll(ctx context.Context, Search, Sort string, Page, Perpage int) ([]entity.Subjects, int64, error) {
	var subjectmodel []entity.Subjects
	var count int64
	order := fmt.Sprintf("created_at %s", Sort)
	offset := (Page - 1) * Perpage

	if err := m.db.WithContext(ctx).Model(&entity.Subjects{}).Debug().Where("name_subject LIKE ?", "%"+Search+"%").Count(&count).Order(order).Limit(Perpage).Offset(offset).Find(&subjectmodel).Error; err != nil {
		return nil, 0, err
	}

	return subjectmodel, count, nil
}

func (m *MysqlStruct) Create(ctx context.Context, data *entity.Subjects) error {
	if err := m.db.WithContext(ctx).Table("subjects").Debug().Create(data).Error; err != nil {
		return err
	}
	return nil
}

func (m *MysqlStruct) Show(ctx context.Context, id int) (*entity.Subjects, error) {
	var finddata entity.Subjects

	if err := m.db.WithContext(ctx).Model(&entity.Subjects{}).Debug().Where("id_subject = ?", id).First(&finddata).Error; err != nil {
		return nil, err
	}

	return &finddata, nil
}

func (m *MysqlStruct) Update(ctx context.Context, id int, data *entity.Subjects) error {
	result := m.db.WithContext(ctx).Model(&entity.Subjects{}).Debug().Where("id_subject = ?", id).Updates(data)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (m *MysqlStruct) Delete(ctx context.Context, id int) error {
	result := m.db.WithContext(ctx).Model(&entity.Subjects{}).Delete(id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
