package entity

import (
	"time"

	"github.com/Zyprush18/Scorely/helper"
)

type Exams struct {
	IdExam      uint      `json:"id_exam" gorm:"primaryKey;autoIncrement"`
	Dates       time.Time `json:"date" gorm:"type:date"`
	StartLesson time.Time `json:"start_lesson" gorm:"type:timestamp"`
	EndLesson   time.Time `json:"end_lesson" gorm:"type:timestamp"`
	SubjectId   uint      `json:"subject_id"`

	// belongs to subjects table
	Subject Subjects `gorm:"foreignKey:SubjectId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	// has many to exam question table
	ExamQuestion []Exam_Questions `gorm:"foreignKey:ExamId;references:IdExam;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// has many to answer question table
	AnswerQuestion []Answer_Questions `gorm:"foreignKey:ExamId;references:IdExam;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	helper.Models
}
