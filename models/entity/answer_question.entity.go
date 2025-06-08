package entity

import "github.com/Zyprush18/Scorely/helper"

type Answer_Questions struct {
	IdAnswerQuestion uint   `json:"id_answer_question" gorm:"primaryKey;autoIncrement"`
	Option           string `json:"option" gorm:"type:varchar(1)"`
	IsCorrect        bool   `json:"is_correct" gorm:"type:bool;default:false"`
	StudentId        uint   `json:"student_id"`
	ExamId           uint   `json:"exam_id"`
	ExamQuestionId   uint   `json:"exam_question_id"`

	// belongs to student table
	Student Students `gorm:"foreignKey:StudentId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// belongs to student table
	Exam Exams `gorm:"foreignKey:ExamId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// belongs to student table
	ExamQuestion Exam_Questions `gorm:"foreignKey:ExamQuestionId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	helper.Models
}
