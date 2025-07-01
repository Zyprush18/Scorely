package repouser

import (
	"fmt"
	"time"

	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/models/entity"
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/models/response"
	"gorm.io/gorm"
)

type UserRepo interface {
	GetAll(search, sort string, page,perpage int) ([]response.Users, int64,error)
	Create(data *request.User) error
	Show(id int) (*response.Users, error)
	Update(id int, data *request.User) error
	Delete(id int) error
}

type UserMysql struct {
	db *gorm.DB
}

func NewUserDatabase(db *gorm.DB) UserMysql {
	return UserMysql{db: db}
}

func (u *UserMysql) GetAll(search, sort string, page,perpage int) ([]response.Users, int64,error) {
	var modeluser []response.Users
	var count int64
	order := fmt.Sprintf("created_at %s", sort)
	offset := (page - 1) * perpage

	if err := u.db.Table("users").Where("email LIKE ?","%"+search+"%").Count(&count).Order(order).Limit(perpage).Offset(offset).Find(&modeluser).Error; err != nil {
		return nil, 0,err
	}

	return modeluser, count, nil
}

func (u *UserMysql) Create(data *request.User) error {
	req := &request.User{
		Email: data.Email,
		Password: helper.HashingPassword(data.Password),
		RoleId: data.RoleId,
		Models: helper.Models{
			CreatedAt: time.Now().Local(),
		},
	}
	if err := u.db.Table("users").Create(req).Error; err != nil {
		return err
	}

	return nil
}

func (u *UserMysql) Show(id int) (*response.Users, error) {
	var model_user entity.Users

	if err := u.db.Model(&model_user).Preload("Role").Where("id_user = ?", id).First(&model_user).Error; err != nil {
		return nil, err
	}

	resp := &response.Users{
		IdUser:   model_user.IdUser,
		Email:    model_user.Email,
		Password: model_user.Password,
		RoleId:   model_user.RoleId,
		Models: helper.Models{
			CreatedAt: model_user.CreatedAt,
			UpdatedAt: model_user.UpdatedAt,
		},
	}

	return resp, nil

}

func(u *UserMysql) Update(id int, data *request.User) error{
	var user_model entity.Users
	if err:= u.db.Model(&user_model).Where("id_user = ?", id).First(&user_model).Error;err != nil {
		return err
	}

	if err:= u.db.Table("users").Where("id_user = ?", id).Updates(&data).Error;err != nil {
		return err
	}

	return nil
}

func (u *UserMysql) Delete(id int) error {
	var model_user entity.Users
	if err:= u.db.Model(&model_user).Where("id_user = ?",id).First(&model_user).Error;err != nil {
		return err
	}

	if err:= u.db.Delete(&model_user).Error;err != nil {
		return err
	}

	return  nil

}