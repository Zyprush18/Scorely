package entity

import "github.com/Zyprush18/Scorely/service"

type Option_Questions struct {
	IdOptionQuestion uint `json:"id_option_question" gorm:"primaryKey;autoIncrement"`
	Option string `json:"option" gorm:"type:varchar(1)"`
	DescriptionOption string `json:"description_option" gorm:"varchar(225)"`
	IsCorrect bool	`json:"is_correct" gorm:"type:bool;default:false"`
	IdExamQuestion uint `json:"Id_exam_question"`

	// belongs to exam table
	ExamQuestion Exam_Questions `gorm:"foreignKey:IdExamQuestion;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	service.Models
}