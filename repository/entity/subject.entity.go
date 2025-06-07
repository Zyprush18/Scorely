package entity

import "github.com/Zyprush18/Scorely/utils"

type Subjects struct {
	IdSubject   uint   `json:"id_subject" gorm:"primaryKey;autoIncrement"`
	NameSubject string `json:"name_subject" gorm:"type:varchar(100)"`
	Semester    string `json:"semester" gorm:"type:varchar(100)"`

	// has many to teacher table (many to many with teacher table)
	Teacher []Teachers `gorm:"many2many:teacher_subjects;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	// has many to exam table
	Exam []Exams `gorm:"foreignKey:SubjectId;references:IdSubject;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	utils.Models
}
