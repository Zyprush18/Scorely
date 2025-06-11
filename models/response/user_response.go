package response

import "github.com/Zyprush18/Scorely/helper"

type Users struct {
	IdUser   uint   `json:"id_user" `
	Email    string `json:"email"`
	Password string `json:"password"`
	RoleId   uint   `json:"role_id"`
	// belongs to role table
	// Role Roles 
	// has many to teacher table
	// Teacher []Teachers `gorm:"foreignKey:UserId;references:IdUser"`
	// // has many to student table
	// Student []Students `gorm:"foreignKey:UserId;references:IdUser"`
	helper.Models
}