package entity

import "github.com/Zyprush18/Scorely/service"

type Answer_Questions struct {
	IdAnswerQuestion uint `json:"id_answer_question" gorm:"primaryKey;autoIncrement"`
	Option string `json:"option" gorm:"type:varchar(1)"`
	IsCorrect bool `json:"is_correct" gorm:"type:bool;default:false"`
	IdStudent uint `json:"id_student"`
	IdExam uint `json:"id_exam"`
	IdExamQuestion uint `json:"id_exam_question"`

	// belongs to student table
	Student Students `gorm:"foreignKey:IdStudent;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// belongs to student table
	Exam Exams `gorm:"foreignKey:IdExam;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// belongs to student table
	ExamQuestion Exam_Questions `gorm:"foreignKey:IdExamQuestion;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	service.Models
}