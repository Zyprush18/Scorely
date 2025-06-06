package database

import (
	"fmt"


	"github.com/Zyprush18/Scorely/repository/entity"
	"github.com/Zyprush18/Scorely/service"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {
	var err error
	dsn := "root:@tcp(127.0.0.1:3306)/scorely?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), 
	})
	if err != nil {
		
		service.Logfile(err.Error())
		panic("Failed Connect Database!!")
	}
	fmt.Println("Success Connect")

	errs := DB.AutoMigrate(
		&entity.Roles{},
		&entity.Users{},
		&entity.Teachers{},
		&entity.Students{},
		&entity.Class{},      
		&entity.Levels{},
		&entity.Majors{},
		&entity.Exam_Questions{},
		&entity.Option_Questions{},
		&entity.Answer_Questions{},
		&entity.Exams{},
		&entity.Subjects{},
	)
	if errs != nil {
		service.Logfile(errs.Error())
		panic("Failed Migrate Table")
	}

	fmt.Println("Success Migrate")
}