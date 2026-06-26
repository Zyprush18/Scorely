package repouser

import (
	"context"
	"fmt"

	"github.com/Zyprush18/Scorely/models/entity"
	"gorm.io/gorm"
)

type UserRepo interface {
	GetAll(ctx context.Context, search, sort string, page, perpage int) ([]entity.Users, int64, error)
	Create(ctx context.Context, data *entity.Users) error
	Show(ctx context.Context, id int) (*entity.Users, error)
	Update(ctx context.Context, id int, data *entity.Users) error
	Delete(ctx context.Context, id int) error
}

type UserMysql struct {
	db *gorm.DB
}

func NewUserDatabase(db *gorm.DB) UserRepo {
	return &UserMysql{db: db}
}

func (u *UserMysql) GetAll(ctx context.Context, search, sort string, page, perpage int) ([]entity.Users, int64, error) {
	var modeluser []entity.Users
	var count int64
	order := fmt.Sprintf("created_at %s", sort)
	offset := (page - 1) * perpage

	if err := u.db.WithContext(ctx).Model(&entity.Users{}).Preload("Role").Where("email LIKE ?", "%"+search+"%").Count(&count).Order(order).Limit(perpage).Offset(offset).Find(&modeluser).Error; err != nil {
		return nil, 0, err
	}

	return modeluser, count, nil
}

func (u *UserMysql) Create(ctx context.Context, data *entity.Users) error {
	if err := u.db.WithContext(ctx).Model(&entity.Users{}).Create(data).Error; err != nil {
		return err
	}
	return nil
}

func (u *UserMysql) Show(ctx context.Context, id int) (*entity.Users, error) {
	var model_user entity.Users

	if err := u.db.WithContext(ctx).Model(&entity.Users{}).Preload("Role").Where("id_user = ?", id).First(&model_user).Error; err != nil {
		return nil, err
	}

	return &model_user, nil
}

func (u *UserMysql) Update(ctx context.Context, id int, data *entity.Users) error {
	result := u.db.WithContext(ctx).Model(&entity.Users{}).Where("id_user = ?", id).Updates(data)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (u *UserMysql) Delete(ctx context.Context, id int) error {
	result := u.db.WithContext(ctx).Delete(&entity.Users{},id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}