package repouser

import (
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDBForUser() (*gorm.DB, sqlmock.Sqlmock, error) {
	sqlDb, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}
	// defer sqlDb.Close()

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      sqlDb,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})

	if err != nil {
		return nil, nil, err
	}

	return db, mock, nil
}
