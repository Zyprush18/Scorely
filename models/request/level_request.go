package request

import "github.com/Zyprush18/Scorely/helper"

type Levels struct {
	Level   string `json:"level" validate:"required"`
	helper.Models
}