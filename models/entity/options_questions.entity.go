package entity

import "github.com/Zyprush18/Scorely/helper"

type Option_Questions struct {
	IdOptionQuestion  uint   `json:"id_option_question" gorm:"primaryKey;autoIncrement"`
	Option            string `json:"option" gorm:"type:varchar(1)"`
	DescriptionOption string `json:"description_option" gorm:"varchar(225)"`
	IsCorrect         bool   `json:"is_correct" gorm:"type:bool;default:false"`
	ExamQuestionId    uint   `json:"exam_question_id"`

	// belongs to exam table
	ExamQuestion Exam_Questions `gorm:"foreignKey:ExamQuestionId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	helper.Models
}
