package repomajor

import (
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDbMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	Sqldb,mock,err := sqlmock.New()
	if err != nil {
		return  nil, nil, err
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: Sqldb,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})

	if err != nil {
		return nil, nil,err
	}

	return  db,mock,nil
}