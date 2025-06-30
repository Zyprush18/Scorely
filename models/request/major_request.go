package request

import "github.com/Zyprush18/Scorely/helper"

type Majors struct {
	Major             string `json:"major" validate:"required"`
	MajorAbbreviation string `json:"major_abbriviation" validate:"required"`
	helper.Models
}