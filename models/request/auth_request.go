package request

import "github.com/Zyprush18/Scorely/helper"

type Login struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type Register struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	RoleId   uint   `json:"role_id" validate:"required"`
	Model    helper.Models
}