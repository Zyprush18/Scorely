package database

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Zyprush18/Scorely/models/entity"
	"gorm.io/driver/gaussdb"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func checkDBConn(conn string) (*gorm.DB, error) {
	// membuat semuah huruf mnjadi kecil agar ketika di env db connectionnya MYSQL jadi mysql
	nameConn := strings.ToLower(conn)

	// ambil .env
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	namedb := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USERNAME")
	pass := os.Getenv("DB_PASSWORD")


	switch nameConn {
	case "postgres" :
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",host,user,pass,namedb,port)
		return gorm.Open(postgres.Open(dsn), &gorm.Config{})
	case "gauss":
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",host,user,pass,namedb,port)
		return gorm.Open(gaussdb.Open(dsn), &gorm.Config{})
	case "sqlite":
		return gorm.Open(sqlite.Open(namedb), &gorm.Config{})
	default:
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",user,pass,host,port,namedb)
		return gorm.Open(mysql.Open(dsn), &gorm.Config{})
	}
}


func Connect() (*gorm.DB,error) {
	db, err := checkDBConn(os.Getenv("DB_CONNECTION"))
	if err != nil {
		log.Println("Failed Connect DB")
		return nil,err
	}


	ers := db.AutoMigrate(
		&entity.Roles{},
		&entity.Users{},
		&entity.Subjects{},
		&entity.Teachers{},
		&entity.TeacherSubjects{},
		&entity.Students{},
		&entity.Class{},      
		&entity.Levels{},
		&entity.Majors{},
		&entity.Exam_Questions{},
		&entity.Option_Questions{},
		&entity.Answer_Questions{},
		&entity.Exams{},
	)
	if ers != nil {
		log.Println("Failed Migrate database")
		return nil,ers
	}

	return db,nil
}