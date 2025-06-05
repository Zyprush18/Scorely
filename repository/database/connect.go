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

	errs := Migration(DB,
		&entity.Roles{},
		&entity.Users{},
		&entity.Teachers{},
		&entity.Students{},
		&entity.Levels{},
		&entity.Majors{},
		&entity.Class{},
		&entity.Subjects{},
		&entity.Exams{},
		&entity.Exam_Questions{},
		&entity.Option_Questions{},
		&entity.Answer_Questions{},
	)
	if errs != nil {
		service.Logfile(errs.Error())
		panic("Failed Migrate Table")
	}

	
	fmt.Println("Success Migrate")
}


func Migration(D *gorm.DB, model ...interface{}) error {
	for _, m := range model {
		if err := D.AutoMigrate(m); err != nil {
			return err
		}
	}

	return nil
}