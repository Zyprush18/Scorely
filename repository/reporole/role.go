package reporole

import (

	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/models/entity"
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/models/response"

	"gorm.io/gorm"
)

type RoleService interface {
	GetAllDataRole() ([]response.Roles, error)
	CreateRole(data *request.Roles) error
	ShowById(id int) (*response.Roles, error)
}

type RoleMysql struct {
	db *gorm.DB
}

func RolesMysql(db *gorm.DB) RoleMysql {
	return RoleMysql{db: db}
}

// GetAllData
func (r RoleMysql) GetAllDataRole() ([]response.Roles, error) {
	var RoleModel []entity.Roles
	if err := r.db.Table("roles").Preload("Users").Find(&RoleModel).Error;err != nil {
		return nil, err
	}

	resp := []response.Roles{}

	for _, r := range RoleModel {
		resp = append(resp, response.Roles{
			IdRole: r.IdRole,
			NameRole: r.NameRole,
			Users: ResponseRole(r.Users),
			Models: helper.Models{
				CreatedAt: r.CreatedAt,
				UpdatedAt: r.UpdatedAt,
			},
		})
	}

	return resp, nil
}

// create
func (r RoleMysql) CreateRole(data *request.Roles) error {
	// respRole := &response.Roles{
	// 	NameRole: data.NameRole,
	// 	Models: helper.Models{
	// 		CreatedAt: time.Now(),
	// 	},
	// }

	if err := r.db.Table("roles").Create(&data).Error; err != nil {
		return err
	}

	return nil
}

// show
func (r RoleMysql) ShowById(id int) (*response.Roles, error) {
	var rolemodel entity.Roles

	if err := r.db.Model(&rolemodel).Preload("Users").Where("id_role = ?", id).First(&rolemodel).Error; err != nil {
		return nil, err
	}

	resp := response.Roles{
		IdRole: rolemodel.IdRole,
		NameRole: rolemodel.NameRole,
		Users: ResponseRole(rolemodel.Users),
		Models: helper.Models{
			CreatedAt: rolemodel.CreatedAt,
			UpdatedAt: rolemodel.UpdatedAt,
		},
	}

	return &resp, nil

}

func ResponseRole(data []entity.Users) []response.Users {
	var result []response.Users
	for _, d := range data {
			result = append(result, response.Users{
				IdUser: d.IdUser,
				Email: d.Email,
				Password: d.Password,
				RoleId: d.RoleId,
			})
	}
	return result
}
