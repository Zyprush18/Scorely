package entity

import "github.com/Zyprush18/Scorely/service"

type Levels struct {
	IdLevel uint `json:"id_level" gorm:"primaryKey;autoIcrement"`
	Level string `json:"level" gorm:"type:varchar(100)"`

	// has many to class table
	Class []Class `gorm:"foreignKey:LevelId;references:IdLevel;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	service.Models
}