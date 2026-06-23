package repoauth

import (
	"context"

	"github.com/Zyprush18/Scorely/models/entity"
	"github.com/Zyprush18/Scorely/models/request"
	"gorm.io/gorm"
)

type RepoAuth interface {
	Login(ctx context.Context, email string) (*entity.Users,error)
	Register(ctx context.Context, data *request.Register) error
}

type MysqlStruct struct {
	db *gorm.DB
}

func ConnectDb(d *gorm.DB) RepoAuth {
	return &MysqlStruct{db: d}
}

func (m *MysqlStruct) Login(ctx context.Context, email string) (*entity.Users,error) {
	var model_user entity.Users
	if err := m.db.WithContext(ctx).Model(&model_user).Preload("Role").Debug().Where("email = ?", email).First(&model_user).Error;err!= nil {
		return nil,err
	}

	return &model_user, nil
}

func (m *MysqlStruct) Register(ctx context.Context, data *request.Register) error {
	if err := m.db.WithContext(ctx).Table("users").Create(data).Error; err != nil {
		return err
	}

	return nil
}