package response

import "github.com/Zyprush18/Scorely/helper"

type Roles struct {
	helper.Models
	NameRole string `json:"name_role" validate:"required,min=3"`
}
