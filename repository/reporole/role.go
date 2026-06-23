package reporole

import (
	"context"
	"fmt"
	"time"

	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/models/entity"
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/models/response"

	"gorm.io/gorm"
)

type RoleService interface {
	GetAllDataRole(ctx context.Context, search,sort string,page,perpage int) ([]response.Roles, int64,error)
	CreateRole(ctx context.Context, data *request.Roles) error
	ShowById(ctx context.Context, id int) (*response.Roles, error)
	UpdateRole(ctx context.Context, id int, data *request.Roles) error
	DeleteRole(ctx context.Context, id int) error
}

type RoleMysql struct {
	db *gorm.DB
}

func RolesMysql(db *gorm.DB) RoleService {
	return &RoleMysql{db: db}
}

// GetAllData
func (r *RoleMysql) GetAllDataRole(ctx context.Context, search,sort string, page,perpage int) ([]response.Roles, int64 ,error) {
	var RoleModel []response.Roles
	var count int64
	offset := (page - 1 ) * perpage
	order := fmt.Sprintf("created_at %s", sort)

	if  err := r.db.WithContext(ctx).Model(&entity.Roles{}).Where("name_role LIKE ?", "%"+search+"%").Count(&count).Order(order).Limit(perpage).Offset(offset).Find(&RoleModel).Error; err != nil {
		return nil, 0, err
	}
	
	return RoleModel, count,nil
}

// create
func (r *RoleMysql) CreateRole(ctx context.Context, data *request.Roles) error {
	role:= &request.Roles{
		NameRole: data.NameRole,
		CodeRole: data.CodeRole,
		Model: helper.Models{
			CreatedAt: time.Now(),
		},
	}
	if err := r.db.WithContext(ctx).Model(&entity.Roles{}).Create(role).Error; err != nil {
		return err
	}

	return nil
}

// show
func (r *RoleMysql) ShowById(ctx context.Context, id int) (*response.Roles, error) {
	var rolemodel response.Roles

	if err := r.db.WithContext(ctx).Model(&entity.Roles{}).Where("id_role = ?", id).First(&rolemodel).Error; err != nil {
		return nil, err
	}


	return &rolemodel, nil

}

func (r *RoleMysql) UpdateRole(ctx context.Context, id int, data *request.Roles) error  {
	result := r.db.WithContext(ctx).Model(&entity.Roles{}).Where("id_role = ?", id).Updates(data)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *RoleMysql) DeleteRole(ctx context.Context, id int) error {
	result := r.db.WithContext(ctx).Model(&entity.Roles{}).Delete(id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}