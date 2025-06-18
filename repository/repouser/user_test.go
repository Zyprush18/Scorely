package repouser

import (
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	database, mocks, err := SetupDBForUser()
	assert.NoError(t, err)
	repousers := NewUserDatabase(database)

	userRow := sqlmock.NewRows([]string{
		"id_user",
		"email",
		"password",
		"role_id",
	}).AddRow(1,"Admin@gmail.com","Admin124", 1)

	t.Run("Success Get All User", func(t *testing.T) {
		mocks.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users`")).WillReturnRows(userRow)

		data, err := repousers.GetAll()
		assert.NoError(t, err)
		assert.NotNil(t, data)

		assert.NoError(t, mocks.ExpectationsWereMet())
	})

	t.Run("Failed Get All User", func(t *testing.T) {
		mocks.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users`")).WillReturnError(database.Error)

		data, err := repousers.GetAll()
		assert.Error(t, err)
		assert.Nil(t, data)

		assert.NoError(t, mocks.ExpectationsWereMet())
	})
}

func TestCreateUser(t *testing.T) {
	databases, mocks, err := SetupDBForUser()
	assert.NoError(t, err)
	repouser := NewUserDatabase(databases)

	t.Run("Success Create a New User", func(t *testing.T) {
		reqUser := &request.User{
			Email:    "Admin11@gmail.com",
			Password: "admin123",
			RoleId:   3,
		}
		mocks.ExpectBegin()
		mocks.ExpectExec(regexp.QuoteMeta("INSERT INTO `users` ")).WithArgs(reqUser.Email, reqUser.Password, reqUser.RoleId, sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 1))
		mocks.ExpectCommit()

		err := repouser.Create(reqUser)
		assert.NoError(t, err)
		assert.NoError(t, mocks.ExpectationsWereMet())
	})

	t.Run("Failed Create a New User", func(t *testing.T) {
		reqUser := &request.User{
			Email:    "Admin@gmail.com",
			Password: "admin123",
			RoleId:   1,
		}
		mocks.ExpectBegin()
		mocks.ExpectExec(regexp.QuoteMeta("INSERT INTO `users` ")).WithArgs(reqUser.Email, reqUser.Password, reqUser.RoleId).WillReturnError(errors.New("Database is Refused"))
		mocks.ExpectRollback()

		err := repouser.Create(reqUser)
		assert.Error(t, err)
		assert.Error(t, mocks.ExpectationsWereMet())
	})

}

func TestShowuserById(t *testing.T) {
	databases, mocks, err := SetupDBForUser()
	assert.NoError(t, err)
	repouser := NewUserDatabase(databases)

	dataUser := sqlmock.NewRows([]string{
		"id_user",
		"email",
		"password",
		"role_id",
	}).AddRow(1, "Admin@gmail.com", "admin123", 1)

	dataRole := sqlmock.NewRows([]string{
		"id_role",
		"name_role",
	}).AddRow(1, "Admin")
	t.Run("Success Show User By Id", func(t *testing.T) {
		id_success := 1

		mocks.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE id_user = ? ORDER BY `users`.`id_user` LIMIT ?")).WithArgs(id_success, 1).WillReturnRows(dataUser)
		mocks.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `roles` WHERE `roles`.`id_role` = ?")).
			WithArgs(id_success).
			WillReturnRows(dataRole)

		data, err := repouser.Show(id_success)
		assert.NoError(t, err)
		assert.Equal(t, uint(1), data.IdUser)

		assert.NoError(t, mocks.ExpectationsWereMet())
	})

	t.Run("Failed Show User By id", func(t *testing.T) {
		id_failed := 2

		mocks.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE id_user = ? ORDER BY `users`.`id_user` LIMIT ?")).WithArgs(id_failed, 2).WillReturnRows(dataUser)
		mocks.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `roles` WHERE `roles`.`id_role` = ?")).
			WithArgs(id_failed).
			WillReturnRows(dataRole)

		data, err := repouser.Show(id_failed)
		assert.Error(t, err)
		assert.Nil(t, data)

		assert.Error(t, mocks.ExpectationsWereMet())
	})
}
