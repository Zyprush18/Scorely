package reporole

import (
	"fmt"
	"time"

	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/models/entity"
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/models/response"

	"gorm.io/gorm"
)

type RoleService interface {
	GetAllDataRole(search,sort string,page,perpage int) ([]response.Roles, int64,error)
	CreateRole(data *request.Roles) error
	ShowById(id int) (*response.Roles, error)
	UpdateRole(id int, data *request.Roles) error
	DeleteRole(id int) error
}

type RoleMysql struct {
	db *gorm.DB
}

func RolesMysql(db *gorm.DB) RoleMysql {
	return RoleMysql{db: db}
}

// GetAllData
func (r RoleMysql) GetAllDataRole(search,sort string, page,perpage int) ([]response.Roles, int64 ,error) {
	var RoleModel []entity.Roles
	var count int64
	offset := (page - 1 ) * perpage
	order := fmt.Sprintf("created_at %s", sort)

	query := r.db.Table("roles").Where("name_role LIKE ?", "%"+search+"%")
	query.Count(&count)
	if  query.Error != nil {
		return nil, 0, query.Error
	}

	query.Order(order).Limit(perpage).Offset(offset).Find(&RoleModel)
	resp := []response.Roles{}

	for _, r := range RoleModel {
		resp = append(resp, response.Roles{
			IdRole: r.IdRole,
			NameRole: r.NameRole,
			// Users: ResponseRole(r.Users),
			Models: helper.Models{
				CreatedAt: r.CreatedAt,
				UpdatedAt: r.UpdatedAt,
			},
		})
	}

	return resp, count,nil
}

// create
func (r RoleMysql) CreateRole(data *request.Roles) error {
	role:= &request.Roles{
		NameRole: data.NameRole,
		Models: helper.Models{
			CreatedAt: time.Now(),
		},
	}
	if err := r.db.Table("roles").Create(&role).Error; err != nil {
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
		// Users: ResponseRole(rolemodel.Users),
		Models: helper.Models{
			CreatedAt: rolemodel.CreatedAt,
			UpdatedAt: rolemodel.UpdatedAt,
		},
	}

	return &resp, nil

}

func (r RoleMysql) UpdateRole(id int, data *request.Roles) error  {
	var rolemodel entity.Roles
	if err := r.db.Model(&rolemodel).Where("id_role = ?", id).First(&rolemodel).Error; err != nil {
		return  err
	}

	if err:= r.db.Table("roles").Where("id_role = ?", id).Updates(data).Error;err != nil {
		return err
	}

	return nil
}

func (r RoleMysql) DeleteRole(id int) error {
	var modelrole entity.Roles
	if err:= r.db.Model(&modelrole).Where("id_role = ?", id).First(&modelrole).Error;err != nil {
		return err
	}

	if err:= r.db.Delete(&modelrole).Error;err != nil {
		return err
	}

	return nil
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