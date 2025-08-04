package request

import "github.com/Zyprush18/Scorely/helper"

type Exam_Questions struct {
	Question       string `json:"question" validate:"required"`
	ExamId         uint   `json:"exam_id"`

	helper.Models
}