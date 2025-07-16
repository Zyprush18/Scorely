package response

import (
	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/models/entity"
)

type Teachers struct {
	IdTeacher uint   `json:"id_teacher"`
	Name      string `json:"name"`
	Nip       string `json:"nip"`
	Gender    string `json:"gender"`
	Address   string `json:"address"`
	Phone     uint   `json:"phone"`
	UserId    uint   `json:"user_id"`

	// di ubah jadi relasi one to one
	// belongs to users table
	// User Users `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`


	// has many to subjects table (many to many with subjects table)
	Subject []Subjects 
	helper.Models
}

func RespGetALl(data []entity.Teachers) (resp []Teachers) {
	for _, v := range data {
		resp = append(resp, Teachers{
			IdTeacher: v.IdTeacher,
			Name: v.Name,
			Nip: v.Nip,
			Gender: v.Gender,
			Address: v.Address,
			Phone: v.Phone,
			UserId: v.UserId,
			Subject: Subjectsresp(v.Subject),
			Models: v.Models,
		})
	}

	return resp
}
