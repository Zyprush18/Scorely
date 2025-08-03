package response

import (
	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/models/entity"
)

type Exam_Questions struct {
	IdExamQuestion uint   `json:"id_exam_question"`
	Question       string `json:"question"`
	ExamId         uint   `json:"exam_id"`

	Exam Exams 

	
	// OptionQuestion []Option_Questions 
	// AnswerQuestion []Answer_Questions 
	helper.Models
}

func ParseExamsQuest(data []entity.Exam_Questions) (resp []Exam_Questions)  {
	for _, v := range data {
		resp = append(resp, Exam_Questions{
			IdExamQuestion: v.IdExamQuestion,
			Question: v.Question,
			ExamId: v.ExamId,
			Exam: Exams{
				IdExam: v.Exam.IdExam,
				NameExams: v.Exam.NameExams,
				Dates: v.Exam.Dates,
				StartLesson: v.Exam.StartLesson,
				EndLesson: v.Exam.EndLesson,
				TeacherSubjectId: v.Exam.TeacherSubjectId,
				Models: v.Exam.Models,
			},
			Models: v.Models,
		})
	}

	return resp
}