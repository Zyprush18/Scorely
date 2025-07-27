package entity

type TeacherSubjects struct {
	IdTeacherSubject uint `json:"id_teacher_subject" gorm:"primaryKey;autoIncrement"`
	IdTeachers	int `gorm:"index"`
	IdSubjects	int `gorm:"index"`
	Exam 		[]Exams `gorm:"foreignKey:TeacherSubjectId;references:IdTeacherSubject;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Subject 	Subjects `gorm:"foreignKey:IdSubjects;references:IdSubject;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Teacher 	Teachers `gorm:"foreignKey:IdTeachers;references:IdTeacher;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}