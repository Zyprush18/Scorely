package entity

import "github.com/Zyprush18/Scorely/helper"

type Levels struct {
	IdLevel uint   `json:"id_level" gorm:"primaryKey;autoIcrement"`
	Level   string `json:"level" gorm:"type:varchar(100);unique"`

	// has many to class table
	Class []Class `gorm:"foreignKey:LevelId;references:IdLevel;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	helper.Models
}
