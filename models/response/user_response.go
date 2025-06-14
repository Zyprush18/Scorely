package response

import "github.com/Zyprush18/Scorely/helper"

type Users struct {
	IdUser   uint   `json:"id_user" `
	Email    string `json:"email"`
	Password string `json:"password"`
	RoleId   uint   `json:"role_id"`
	helper.Models
}