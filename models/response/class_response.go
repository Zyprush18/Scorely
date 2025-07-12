package response

import "github.com/Zyprush18/Scorely/helper"

type Class struct {
	IdClass uint   `json:"id_class"`
	Name    string `json:"name"`
	LevelId uint   `json:"level_id"`
	MajorId uint   `json:"major_id"`
	Level 	Levels `json:"level"`
	Major	Majors `json:"major"`
	helper.Models
}

