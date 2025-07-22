package entity

import "github.com/Zyprush18/Scorely/helper"

// table role
type Roles struct {
	IdRole   uint   `json:"id_role" gorm:"primaryKey;autoIncrement"`
	NameRole string `json:"name_role" gorm:"not null;unique"`
	CodeRole string `json:"code_role" gorm:"not null;unique"`
	// has many to users table
	Users []Users `gorm:"foreignKey:RoleId;references:IdRole;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	helper.Models
}
