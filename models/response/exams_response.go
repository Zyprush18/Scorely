package response

import (
	"time"

	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/models/entity"
)


type Exams struct {
	IdExam      uint      `json:"id_exam"`
	NameExams   string    `json:"name_exam"`
	Dates       time.Time `json:"date"`
	StartLesson string `json:"start_lesson"`
	EndLesson   string `json:"end_lesson"`
	SubjectId   uint      `json:"subject_id"`

	// belongs to subjects table
	Subject Subjects 

	// has many to exam question table
	// ExamQuestion []Exam_Questions `gorm:"foreignKey:ExamId;references:IdExam;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// has many to answer question table
	// AnswerQuestion []Answer_Questions `gorm:"foreignKey:ExamId;references:IdExam;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	helper.Models
}

func ParseExams(data []entity.Exams) (resp []Exams) {
	for _, v := range data {
		resp = append(resp, Exams{
			IdExam: v.IdExam,
			NameExams: v.NameExams,
			Dates: v.Dates,
			StartLesson: v.StartLesson,
			EndLesson: v.EndLesson,
			SubjectId: v.SubjectId,
			Subject: Subjects{
				IdSubject: v.SubjectId,
				NameSubject: v.Subject.NameSubject,
				Semester: v.Subject.Semester,
				Models: v.Subject.Models,
			},
			Models: v.Models,
		})
	}
	return resp
}