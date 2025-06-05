package entity

import "github.com/Zyprush18/Scorely/service"

// table users
type Users struct {
	IdUser uint `json:"id_user" gorm:"primaryKey;autoIncrement"`
	Email string `json:"email" gorm:"unique;type:varchar(50)"`
	Password string `json:"password" gorm:"type:varchar(50)"`
	IdRole uint `json:"id_role"` 
	// belongs to role table
	Role Roles `gorm:"foreignKey:IdRole;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// has many to teacher table
	Teacher []*Teachers `gorm:"foreignKey:IdUser;references:IdUser"`
	// has many to student table
	Student []*Students `gorm:"foreignKey:IdUser;references:IdUser"`
	service.Models
}