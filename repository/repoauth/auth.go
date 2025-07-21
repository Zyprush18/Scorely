package repoauth

import (


	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/models/entity"
	"github.com/Zyprush18/Scorely/models/request"
	"gorm.io/gorm"
)

type RepoAuth interface {
	Login(loginreq *request.Login) error
}

type MysqlStruct struct {
	db *gorm.DB
}

func ConnectDb(d *gorm.DB) MysqlStruct {
	return MysqlStruct{db: d}
}

func (m *MysqlStruct) Login(loginreq *request.Login) error {
	var model_user entity.Users
	if err := m.db.Table("users").Debug().Where("email = ?", loginreq.Email).First(&model_user).Error;err!= nil {
		return err
	}

	if err := helper.DecryptPassword(model_user.Password,loginreq.Password);err != nil {
		return err
	}

	return nil
}