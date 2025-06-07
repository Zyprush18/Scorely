package service

import (
	"time"

	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/models/response"
	"github.com/Zyprush18/Scorely/repository/database"
	"github.com/Zyprush18/Scorely/utils"
)

func AddRoleLogic(reqval *request.Roles) error {
	respRole := &response.Roles{
		NameRole: reqval.NameRole,
		Models: utils.Models{
			CreatedAt: time.Now(),
		},
	}

	if err:= database.DB.Table("roles").Create(&respRole).Error;err != nil {
		return err
	}

	return nil
}