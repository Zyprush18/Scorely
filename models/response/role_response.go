package response

import "github.com/Zyprush18/Scorely/helper"



type Roles struct {
	IdRole   uint   `json:"id_role"`
	NameRole string `json:"name_role"`
	// Users []Users	`json:"user"`
	helper.Models
	
}
