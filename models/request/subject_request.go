package request

import "github.com/Zyprush18/Scorely/helper"

type Subjects struct {
	NameSubject string `json:"name_subject" validate:"required"`
	Semester    string `json:"semester" validate:"required"`
	helper.Models
}
