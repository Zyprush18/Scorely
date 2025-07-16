package response

import (
	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/models/entity"
)

type Subjects struct {
	IdSubject   uint   `json:"id_subject"`
	NameSubject string `json:"name_subject"`
	Semester    string `json:"semester"`

	// has many to teacher table (many to many with teacher table)
	// Teacher []Teachers `gorm:"many2many:teacher_subjects;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	// has many to exam table
	// Exam []Exams 

	helper.Models
}

func Subjectsresp(data []entity.Subjects) (resp []Subjects)  {
	for _, v := range data {
		resp = append(resp, Subjects{
			IdSubject: v.IdSubject,
			NameSubject: v.NameSubject,
			Semester: v.Semester,
			Models: v.Models,
		})
	}
	
	return resp
}