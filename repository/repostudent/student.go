package repostudent

import (
	"fmt"
	"time"

	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/models/entity"
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/models/response"
	"gorm.io/gorm"
)

type StudentRepo interface {
	GetAll(Search,Sort string, Page,Perpage int) ([]response.Students, int64, error)
	Create(data *request.Students) error
	Show(id int) (*response.Students, error)
}

type MysqlStruct struct {
	db *gorm.DB
}

func ConnectDb(d *gorm.DB) MysqlStruct  {
	return MysqlStruct{db: d}
}

func (m *MysqlStruct) GetAll(Search,Sort string, Page,Perpage int) ([]response.Students, int64, error) {
	var findstudent []entity.Students
	var count int64
	order := fmt.Sprintf("created_at %s", Sort)
	offset := (Page - 1) * Perpage

	if err := m.db.Model(&entity.Students{}).Preload("Class.Major").Preload("Class.Level").Debug().Where("nisn LIKE ?", "%"+Search+"%").Count(&count).Order(order).Limit(Perpage).Offset(offset).Find(&findstudent).Error; err != nil {
		return nil, 0, err
	}

	return ParseResponse(findstudent),count,nil
}

func (m *MysqlStruct) Create(data *request.Students) error {
	now := time.Now()
	studentreq := &request.Students{
		Name: data.Name,
		Nisn: data.Nisn,
		Gender: data.Gender,
		Address: data.Address,
		Phone: data.Phone,
		UserId: data.UserId,
		ClassId: data.ClassId,
		Models: helper.Models{
			CreatedAt: now,
		},
	}

	if err := m.db.Table("students").Create(studentreq).Error;err != nil {
		return  err
	}

	return nil
}

func (m *MysqlStruct) Show(id int) (*response.Students, error) {
	var findstudent entity.Students
	if err:= m.db.Model(findstudent).Preload("Class.Major").Preload("Class.Level").Where("id_student = ?", id).First(&findstudent).Error;err != nil {
		return nil, err
	}

	return &response.Students{
		IdStudent: findstudent.IdStudent,
		Name: findstudent.Name,
		Nisn: findstudent.Nisn,
		Gender: findstudent.Gender,
		Address: findstudent.Address,
		Phone: findstudent.Phone,
		UserId: findstudent.UserId,
		ClassId: findstudent.ClassId,
		Class: response.Class{
			IdClass: findstudent.Class.IdClass,
			Name: findstudent.Class.Name,
			LevelId: findstudent.Class.LevelId,
			MajorId: findstudent.Class.MajorId,
			Level: response.Levels{
				IdLevel: findstudent.Class.Level.IdLevel,
				Level:  findstudent.Class.Level.Level,
				Models: findstudent.Class.Level.Models,
			},
			Major: response.Majors{
				IdMajor: findstudent.Class.Major.IdMajor,
				Major: findstudent.Class.Major.Major,
				MajorAbbreviation: findstudent.Class.Major.MajorAbbreviation,
				Models: findstudent.Class.Major.Models,
			},
		},
		Models: findstudent.Models,
	},nil
}

func ParseResponse(data []entity.Students) (resp []response.Students) {
	for _, v := range data {
		resp = append(resp, response.Students{
			IdStudent: v.IdStudent,
			Name: v.Name,
			Nisn: v.Nisn,
			Gender: v.Gender,
			Address: v.Address,
			Phone: v.Phone,
			UserId: v.UserId,
			ClassId: v.ClassId,
			Class: response.Class{
				IdClass: v.Class.IdClass,
				Name: v.Class.Name,
				LevelId: v.Class.LevelId,
				MajorId: v.Class.MajorId,
				Models: v.Class.Models,
				Level: response.Levels{
					IdLevel: v.Class.Level.IdLevel,
					Level: v.Class.Level.Level,
					Models: v.Class.Level.Models,
				},
				Major: response.Majors{
					IdMajor: v.Class.Major.IdMajor,
					Major: v.Class.Major.Major,
					MajorAbbreviation: v.Class.Major.MajorAbbreviation,
					Models: v.Class.Major.Models,
				},
			},
			Models: v.Models,
		})
	}

	return resp
}