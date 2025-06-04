package entity

import "github.com/Zyprush18/Scorely/service"

type Classs struct {
	IdClass uint `json:"id_class" gorm:"primaryKey;autoIncrement"`
	Name string `json:"name" gorm:"varchar(100)"`
	IdLevel uint `json:"id_level"`
	IdMajor uint `json:"id_major"`

	// has many to student table
	Student []Students `gorm:"foreignKey:IdClass"` 
	// belongs to level table
	Level Levels `gorm:"foreignKey:IdLevel;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// belongs to level table
	Major Majors `gorm:"foreignKey:IdMajor;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	service.Models
}