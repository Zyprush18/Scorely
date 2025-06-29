package repomajor

import (
	"fmt"

	"github.com/Zyprush18/Scorely/models/response"
	"gorm.io/gorm"
)


type MajorRepo interface {
	GetAllData(search, sort string, page, perpage int) ([]response.Majors, int64, error)
}

type MysqlStruct struct {
	db *gorm.DB
}

func ConnectDb(db *gorm.DB) MysqlStruct  {
	return MysqlStruct{db: db}
}

func (m *MysqlStruct) GetAllData(search, sort string, page, perpage int) ([]response.Majors, int64, error)  {
	var modelmajor []response.Majors
	var count int64
	offset := (page - 1) *perpage
	order := fmt.Sprintf("created_at %s", sort)

	if err := m.db.Table("majors").Where("major LIKE ?", "%"+search+"%").Count(&count).Order(order).Limit(perpage).Offset(offset).Find(&modelmajor).Error; err != nil {
		return  nil, 0, err
	}

	return  modelmajor,count,nil	
}