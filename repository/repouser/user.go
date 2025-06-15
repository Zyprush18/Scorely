package repouser

import (
	"github.com/Zyprush18/Scorely/models/entity"
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/models/response"
	"gorm.io/gorm"
)

type UserRepo interface {
	Create(data *request.User) error
	Show(id int) (*response.Users, error)
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

func (u *UserMysql) Show(id int) (*response.Users, error)  {
	var model_user entity.Users

	if err:= u.db.Model(&model_user).Preload("Role").Where("id_user = ?", id).First(&model_user).Error;err != nil {
		return nil, err
	}

	resp := &response.Users{
		IdUser: model_user.IdUser,
		Email: model_user.Email,
		Password: model_user.Password,
		RoleId: model_user.RoleId,
	}

	return resp, nil

}