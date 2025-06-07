package database

import (
	"fmt"

	"github.com/Zyprush18/Scorely/repository/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	// "gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() error {
	var err error
	dsn := "root:@tcp(127.0.0.1:3306)/scorely?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info), 
	})
	if err != nil {
		fmt.Println("Failed Connect Database")
		return err
	}


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
		fmt.Println("Failed Migrate Table Tp Database")
		return err
	}

	


	return nil
}