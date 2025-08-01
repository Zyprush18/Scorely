package request

import "github.com/Zyprush18/Scorely/helper"

type Exams struct {
	NameExams   string    `json:"name_exam" validate:"required"`
	Dates       string 	  `json:"date" validate:"required"`
	StartLesson string 	  `json:"start_lesson" validate:"required"`
	EndLesson   string 	  `json:"end_lesson" validate:"required"`
	TeacherSubjectId uint `json:"teacher_subject_id"`
	TeacherId *uint 	  `json:"teacher_id,omitempty" gorm:"-"`
	SubjectId *uint 	  `json:"subject_id,omitempty" gorm:"-"`

	helper.Models
}