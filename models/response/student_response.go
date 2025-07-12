package response

import "github.com/Zyprush18/Scorely/helper"

type Students struct {
	IdStudent uint   `json:"id_student" `
	Name      string `json:"name"`
	Nisn      string `json:"nisn"`
	Gender    string `json:"gender"`
	Address   string `json:"address"`
	Phone     uint   `json:"phone"`
	UserId    uint   `json:"user_id"`
	ClassId   uint   `json:"class_id"`

	Class     Class `json:"class"`
	helper.Models
}
