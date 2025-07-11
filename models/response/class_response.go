package response

import "github.com/Zyprush18/Scorely/helper"

type Class struct {
	IdClass uint   `json:"id_class" gorm:"primaryKey;autoIncrement"`
	Name    string `json:"name" gorm:"varchar(100)"`
	LevelId uint   `json:"level_id"`
	MajorId uint   `json:"major_id"`
	Level 	Levels
	Major	Majors
	helper.Models
}

