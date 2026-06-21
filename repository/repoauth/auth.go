package repoauth

import (
	"github.com/Zyprush18/Scorely/models/entity"
	"github.com/Zyprush18/Scorely/models/request"
	"gorm.io/gorm"
)

type RepoAuth interface {
	Login(email string) (*entity.Users,error)
	Register(data *request.Register) error
}

type MysqlStruct struct {
	db *gorm.DB
}

func ConnectDb(d *gorm.DB) MysqlStruct {
	return MysqlStruct{db: d}
}

func (m *MysqlStruct) Login(email string) (*entity.Users,error) {
	var model_user entity.Users
	if err := m.db.Model(&model_user).Preload("Role").Debug().Where("email = ?", email).First(&model_user).Error;err!= nil {
		return nil,err
	}

	return &model_user, nil
}

func (m *MysqlStruct) Register(data *request.Register) error {
	if err := m.db.Table("users").Create(data).Error; err != nil {
		return err
	}

	return nil
}