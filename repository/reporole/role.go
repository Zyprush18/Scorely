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
	ShowById(id int) (*response.Roles, error)
}

type RoleMysql struct {
	db *gorm.DB
}

func RolesMysql(db *gorm.DB) RoleMysql {
	return RoleMysql{db: db}
}

// create
func (r RoleMysql) CreateRole(data *request.Roles) error {
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

// show

func (r RoleMysql) ShowById(id int) (*response.Roles, error) {
	var role response.Roles

	if err := r.db.Table("roles").Where("id_role = ?", id).First(&role).Error; err != nil {
		return nil, err
	}

	return &role, nil

}

// func ResponseRole(data []entity.Users) []response.Users {
// 	var result []response.Users
// 	for _, d := range data {
// 			result = append(result, response.Users{
// 				IdUser: d.IdUser,
// 				Email: d.Email,
// 				Password: d.Password,
// 				RoleId: d.RoleId,
// 			})
// 	}
// 	return result
// }
