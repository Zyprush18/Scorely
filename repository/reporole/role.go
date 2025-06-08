package reporole

import (
	"time"

	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/models/response"

	"gorm.io/gorm"
)

type RoleService interface {
	CreateRole(data *request.Roles) error
}

type RoleMysql struct {
	db *gorm.DB
}

func RolesMysql(db *gorm.DB) RoleMysql  {
	return RoleMysql{db: db}
}

func (r RoleMysql) CreateRole(data *request.Roles) error  {
	respRole := &response.Roles{
		NameRole: data.NameRole,
		Models: helper.Models{
			CreatedAt: time.Now(),
		},
	}

	if err := r.db.Table("roles").Create(&respRole).Error; err != nil {
		return err
	}

	return nil
}

