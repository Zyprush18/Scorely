package entity

import "github.com/Zyprush18/Scorely/utils"

type Class struct {
	IdClass uint   `json:"id_class" gorm:"primaryKey;autoIncrement"`
	Name    string `json:"name" gorm:"varchar(100)"`
	LevelId uint   `json:"level_id"`
	MajorId uint   `json:"major_id"`

	// has many to student table
	Student []Students `gorm:"foreignKey:ClassId;references:IdClass;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// belongs to level table
	Level Levels `gorm:"foreignKey:LevelId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// belongs to level table
	Major Majors `gorm:"foreignKey:MajorId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	utils.Models
}
