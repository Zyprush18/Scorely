package response

import (
	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/models/entity"
)

type Users struct {
	IdUser   uint   `json:"id_user" `
	Email    string `json:"email"`
	Role   string  `json:"role"`
	Model helper.Models
}

func MapUserResp(data *entity.Users) *Users  {
	return &Users{
		IdUser: data.IdUser,
		Email: data.Email,
		Role: data.Role.NameRole,
		Model: helper.Models{
			CreatedAt: data.CreatedAt,
			UpdatedAt: data.UpdatedAt,
		},
	}
}

func ParseUserResponse(entities []entity.Users) (data []Users) {
	for _, v := range entities {
		data = append(data, *MapUserResp(&v))
	}
	return data
}