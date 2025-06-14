package repouser

import (
	"github.com/Zyprush18/Scorely/models/request"
	"gorm.io/gorm"
)

type UserRepo interface {
	Create(data *request.User) error
}

type UserMysql struct {
	db *gorm.DB
}

func NewUserDatabase(db *gorm.DB) UserMysql {
	return UserMysql{db: db}
}

func (u *UserMysql) Create(data *request.User) error  {
	
	if err:= u.db.Table("users").Create(data).Error;err != nil {
		return err
	}

	return nil
}