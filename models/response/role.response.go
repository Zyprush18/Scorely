package response

import "github.com/Zyprush18/Scorely/utils"

type Roles struct {
	utils.Models
	NameRole string `json:"name_role" validate:"required,min=3"`
}