package repouser

import (
	"time"

	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/models/entity"
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/models/response"
	"gorm.io/gorm"
)

type UserRepo interface {
	GetAll() ([]response.Users, error)
	Create(data *request.User) error
	Show(id int) (*response.Users, error)
}

type UserMysql struct {
	db *gorm.DB
}

func NewUserDatabase(db *gorm.DB) UserMysql {
	return UserMysql{db: db}
}

func (u *UserMysql) GetAll() ([]response.Users, error) {
	var modeluser []entity.Users
	if err := u.db.Model(&modeluser).Find(&modeluser).Error; err != nil {
		return nil, err
	}

	resp := []response.Users{}
	for _, u := range modeluser {
		resp = append(resp, response.Users{
			IdUser:   u.IdUser,
			Email:    u.Email,
			Password: u.Password,
			RoleId:   u.RoleId,
			Models: helper.Models{
				CreatedAt: u.CreatedAt,
				UpdatedAt: u.UpdatedAt,
			},
		})
	}

	return resp, nil
}

func (u *UserMysql) Create(data *request.User) error {
	req := &request.User{
		Email: data.Email,
		Password: data.Password,
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


