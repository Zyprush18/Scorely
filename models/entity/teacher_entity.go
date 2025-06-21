package entity

import "github.com/Zyprush18/Scorely/helper"

// table teacher
type Teachers struct {
	IdTeacher uint   `json:"id_teacher" gorm:"primaryKey;autoIncrement"`
	Name      string `json:"name" gorm:"not null;type:varchar(50)"`
	Nip       string `json:"nip" gorm:"type:varchar(50)"`
	Gender    string `json:"gender" gorm:"type:varchar(10)"`
	Address   string `json:"address" gorm:"type:varchar(50)"`
	Phone     uint   `json:"phone" gorm:"type:bigint;unique"`
	UserId    uint   `json:"user_id"`

	// di ubah jadi relasi one to one
	// belongs to users table
	// User Users `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`


	// has many to subjects table (many to many with subjects table)
	Subject []*Subjects `gorm:"many2many:teacher_subjects;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	helper.Models
}
