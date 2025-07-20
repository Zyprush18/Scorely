package entity

import "github.com/Zyprush18/Scorely/helper"

// table users
type Users struct {
	IdUser   uint   `json:"id_user" gorm:"primaryKey;autoIncrement"`
	Email    string `json:"email" gorm:"unique;type:varchar(50)"`
	Password string `json:"password" gorm:"type:varchar(255)"`
	RoleId   uint   `json:"role_id"`
	// belongs to role table
	Role Roles `gorm:"foreignKey:RoleId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// has many to teacher table
	Teacher Teachers `gorm:"foreignKey:UserId;references:IdUser;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// has many to student table
	Student Students `gorm:"foreignKey:UserId;references:IdUser;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	helper.Models
}
