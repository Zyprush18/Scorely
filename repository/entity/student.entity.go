package entity

import "github.com/Zyprush18/Scorely/service"

type Students struct {
	IdStudent uint `json:"id_student" gorm:"primaryKey;autoIncrement"`
	Name string	`json:"name" gorm:"type:varchar(50)"`
	Nisn string `json:"" gorm:"type:varchar(10);unique"`
	Gender    string `json:"gender" gorm:"type:varchar(10)"`
	Address string `json:"address" gorm:"type:varchar(50)"`
	Phone	uint `json:"phone" gorm:"type:bigint;unique"`
	UserId uint	`json:"user_id"`
	ClassId uint `json:"class_id"`
	// belongs to users table
	User Users `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// belongst to class table
	Class Class `gorm:"foreignKey:ClassId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// has many to answer question table
	AnswerQuestion []Answer_Questions `gorm:"foreignKey:StudentId;references:IdStudent;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	service.Models
}