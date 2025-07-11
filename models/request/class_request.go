package request

import "github.com/Zyprush18/Scorely/helper"

type Class struct {
	Name    string `json:"name" validate:"required"`
	LevelId uint   `json:"level_id" validate:"required"`
	MajorId uint   `json:"major_id" validate:"required"`

	helper.Models
}