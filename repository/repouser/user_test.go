package repouser

import (
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T)  {
	databases, mocks, err := SetipDBForUser()
	assert.NoError(t, err)
	repouser := NewUserDatabase(databases)

	t.Run("Success Create a New User", func(t *testing.T) {
		reqUser := &request.User{
			Email: "Admin@gmail.com",
			Password: "admin123",
			RoleId: 1,
		}
		mocks.ExpectBegin()
		mocks.ExpectExec(regexp.QuoteMeta("INSERT INTO `users` ")).WithArgs(reqUser.Email,reqUser.Password,reqUser.RoleId).WillReturnResult(sqlmock.NewResult(1,1))
		mocks.ExpectCommit()

		err := repouser.Create(reqUser)
		assert.NoError(t, err)
		assert.NoError(t, mocks.ExpectationsWereMet())
	})

	t.Run("Failed Create a New User", func(t *testing.T) {
		reqUser := &request.User{
			Email: "Admin@gmail.com",
			Password: "admin123",
			RoleId: 1,
		}
		mocks.ExpectBegin()
		mocks.ExpectExec(regexp.QuoteMeta("INSERT INTO `users` ")).WithArgs(reqUser.Email,reqUser.Password,reqUser.RoleId).WillReturnError(errors.New("Database is Refused"))
		mocks.ExpectRollback()

		err := repouser.Create(reqUser)
		assert.Error(t, err)
		assert.NoError(t, mocks.ExpectationsWereMet())
	})

}