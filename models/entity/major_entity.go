package entity

import "github.com/Zyprush18/Scorely/helper"

type Majors struct {
	IdMajor           uint   `json:"id_major" gorm:"primaryKey;autoIncrement"`
	Major             string `json:"major" gorm:"type:varchar(50)"`
	MajorAbbreviation string `json:"major_abbriviation" gorm:"varchar(20)"`

	// has many to class table
	Class []Class `gorm:"foreignKey:MajorId;references:IdMajor;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	helper.Models
}
