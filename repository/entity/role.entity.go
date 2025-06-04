package entity

import "github.com/Zyprush18/Scorely/service"

// table role
type Roles struct {
	IdRole uint `json:"id_role" gorm:"primaryKey;autoIncrement"`
	NameRole string `json:"name_role" gorm:"not null"`
	// has many to users table
	User []Users	`gorm:"foreignKey:IdRole"`
	service.Models
}