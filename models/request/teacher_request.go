package request

import "github.com/Zyprush18/Scorely/helper"

type Teachers struct {
	Name      string `json:"name" validate:"required,min=3"`
	Nip       string `json:"nip" validate:"required,min=9,max=18"`
	Gender    string `json:"gender" validate:"required"`
	Address   string `json:"address" validate:"required"`
	Phone     uint   `json:"phone" validate:"required,min=10"`
	UserId    uint   `json:"user_id" validate:"required"`
	SubjectId   []int `json:"subject_id" validate:"required"`

	helper.Models
}