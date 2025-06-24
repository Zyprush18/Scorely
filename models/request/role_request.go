package request

import "github.com/Zyprush18/Scorely/helper"


type Roles struct {
	NameRole string `json:"name_role" validate:"required,min=3"`
	helper.Models
}