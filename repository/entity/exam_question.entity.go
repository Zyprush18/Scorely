package entity

import "github.com/Zyprush18/Scorely/service"

type Exam_Questions struct {
	IdExamQuestion uint `json:"id_exam_question" gorm:"primaryKey;autoIncrement"`
	Question string `json:"question" gorm:"type:text"`
	IdExam uint `json:"id_exam"`

	// belongs to exam table
	Exam Exams `gorm:"foreignKey:IdExam;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	// has many to option question table
	OptionQuestion []*Option_Questions `gorm:"foreignKey:IdExamQuestion;references:IdExamQuestion"`
	// has many to answer question table
	AnswerQuestion []*Answer_Questions `gorm:"foreignKey:IdExamQuestion;references:IdExamQuestion"`
	service.Models
}