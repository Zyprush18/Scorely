package response

import "github.com/Zyprush18/Scorely/helper"



type Roles struct {
	IdRole   uint   `json:"id_role"`
	NameRole string `json:"name_role" validate:"required,min=3"`
	// User []Users	`json:"user"`
	helper.Models
	
}
