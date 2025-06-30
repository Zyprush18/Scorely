package response

import "github.com/Zyprush18/Scorely/helper"

type Majors struct {
	IdMajor           uint   `json:"id_major"`
	Major             string `json:"major"`
	MajorAbbreviation string `json:"major_abbriviation"`

	// has many to class table
	// Class []Class `gorm:"foreignKey:MajorId;references:IdMajor;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	helper.Models
}