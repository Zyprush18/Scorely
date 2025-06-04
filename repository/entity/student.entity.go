package entity

import "github.com/Zyprush18/Scorely/service"

type Students struct {
	IdStudent uint `json:"id_student" gorm:"primaryKey;autoIncrement"`
	Name string	`json:"name" gorm:"type:varchar(50)"`
	Nisn string `json:"" gorm:"type:varchar(10);unique"`
	Gender string 	`json:"gender" gorm:"type:enum('male','female')"`
	Address string `json:"address" gorm:"type:varchar(50)"`
	Phone	uint `json:"phone" gorm:"type:bigint;unique"`
	IdUser uint	`json:"id_user"`
	IdClass uint `json:"id_class"`
	// belongs to users table
	User Users `gorm:"foreignKey:IdUser;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// belongst to class table
	Class Classs `gorm:"foreignKey:IdClass;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// has many to answer question table
	AnswerQuestion []Answer_Questions `gorm:"foreignKey:IdStudent"`
	service.Models
}