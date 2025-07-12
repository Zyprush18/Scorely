package request

import "github.com/Zyprush18/Scorely/helper"

type Students struct {
	Name      string `json:"name" validate:"required,min=3"`
	Nisn      string `json:"nisn" validate:"required,min=10,max=10"`
	Gender    string `json:"gender" validate:"required"`
	Address   string `json:"address" validate:"required,min=3"`
	Phone     string  `json:"phone" validate:"required,min=10,numeric"`
	UserId    uint   `json:"user_id" validate:"required"`
	ClassId   uint   `json:"class_id" validate:"required"`

	helper.Models
}