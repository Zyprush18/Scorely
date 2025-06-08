package entity

import "github.com/Zyprush18/Scorely/helper"

type Exam_Questions struct {
	IdExamQuestion uint   `json:"id_exam_question" gorm:"primaryKey;autoIncrement"`
	Question       string `json:"question" gorm:"type:text"`
	ExamId         uint   `json:"exam_id"`

	// belongs to exam table
	Exam Exams `gorm:"foreignKey:ExamId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	// has many to option question table
	OptionQuestion []Option_Questions `gorm:"foreignKey:ExamQuestionId;references:IdExamQuestion;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// has many to answer question table
	AnswerQuestion []Answer_Questions `gorm:"foreignKey:ExamQuestionId;references:IdExamQuestion;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	helper.Models
}
