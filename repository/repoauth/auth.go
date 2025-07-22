package repoauth

import (
	"fmt"

	"github.com/Zyprush18/Scorely/config"
	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/models/entity"
	"github.com/Zyprush18/Scorely/models/request"
	"gorm.io/gorm"
)

type RepoAuth interface {
	Login(loginreq *request.Login) (string,error)
}

type MysqlStruct struct {
	db *gorm.DB
}

func ConnectDb(d *gorm.DB) MysqlStruct {
	return MysqlStruct{db: d}
}

func (m *MysqlStruct) Login(loginreq *request.Login) (string,error) {
	var model_user entity.Users
	if err := m.db.Model(&model_user).Preload("Role").Debug().Where("email = ?", loginreq.Email).First(&model_user).Error;err!= nil {
		return "",err
	}

	if err := helper.DecryptPassword(model_user.Password,loginreq.Password);err != nil {
		return "",err
	}

	
	token, err:= config.GenerateToken(model_user.IdUser,model_user.Role.CodeRole)
	if err != nil {
		fmt.Println("ini errpr")
		return "",err
	}
	return token,nil
}