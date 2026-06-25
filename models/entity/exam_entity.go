package entity

import (
	"time"

	"github.com/Zyprush18/Scorely/helper"
)

type Exams struct {
	IdExam      uint      `json:"id_exam" gorm:"primaryKey;autoIncrement"`
	NameExams   string    `json:"name_exam" gorm:"type:varchar(200)"`
	Dates       time.Time `json:"date" gorm:"type:date"`
	StartLesson string 	  `json:"start_lesson" gorm:"type:varchar(20)"`
	EndLesson   string    `json:"end_lesson" gorm:"type:varchar(20)"`
	// SubjectId   uint      `json:"subject_id"`
	TeacherSubjectId uint `json:"teacher_subject_id"`

	// belongs to subjects table
	TeacherSubject *TeacherSubjects `gorm:"foreignKey:TeacherSubjectId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	// has many to exam question table
	ExamQuestion []Exam_Questions `gorm:"foreignKey:ExamId;references:IdExam;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// has many to answer question table
	AnswerQuestion []Answer_Questions `gorm:"foreignKey:ExamId;references:IdExam;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	helper.Models
}
