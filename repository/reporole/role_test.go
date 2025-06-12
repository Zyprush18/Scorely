package reporole

import (
	"errors"

	"regexp"

	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/stretchr/testify/assert"
)

func TestGetAllDAtaRole(t *testing.T) {
	db, mock, err := SetupMockDb()
	assert.NoError(t, err)
	repo := RolesMysql(db)

	dataRole := sqlmock.NewRows([]string{
		"id_role",
		"name_role",
	}).AddRow(1, "Admin")

	dataUser := sqlmock.NewRows([]string{
		"id_user",
		"email",
		"password",
		"role_id",
	}).AddRow(1, "admin@gmail.com", "admin123", 1).
		AddRow(2, "user@gmail.com", "user1232", 1)

	t.Run("Success Get All Data Role", func(t *testing.T) {
		mock.ExpectQuery("SELECT (.*)").
			WillReturnRows(dataRole)

		mock.ExpectQuery("SELECT (.*)").
			WillReturnRows(dataUser)

		data, err := repo.GetAllDataRole()
		assert.NoError(t, err)
		assert.NotNil(t, data)

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Failed Get All Data Role", func(t *testing.T) {
		mock.ExpectQuery("SELECT (.*)").WillReturnError(errors.New("db error"))

		data, err := repo.GetAllDataRole()
		assert.Error(t, err)
		assert.Nil(t, data)

		assert.NoError(t, mock.ExpectationsWereMet())
	})

}

func TestCreateRole(t *testing.T) {
	db, mock, err := SetupMockDb()
	assert.NoError(t, err)

	repo := RolesMysql(db)

	role := &request.Roles{
		NameRole: "Admin",
	}

	t.Run("Succes Create Role", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `roles` ")).WithArgs(role.NameRole).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		err := repo.CreateRole(role)

		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())

	})

	t.Run("Failed Create Role", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `roles`")).WithArgs(role.NameRole).WillReturnError(sqlmock.ErrCancelled)
		mock.ExpectRollback()

		err := repo.CreateRole(role)
		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())

	})
}

func TestShowRoleById(t *testing.T) {
	db, mock, err := SetupMockDb()
	assert.NoError(t, err)

	repo := RolesMysql(db)

	dataRole := sqlmock.NewRows([]string{
		"id_role",
		"name_role",
	}).AddRow(1, "Admin")

	dataUser := sqlmock.NewRows([]string{
		"id_user",
		"email",
		"password",
		"role_id",
	}).AddRow(1, "admin@gmail.com", "admin123", 1)

	t.Run("Success to Show Role By Id", func(t *testing.T) {
		id_success := 1

		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `roles` WHERE id_role = ? ORDER BY `roles`.`id_role` LIMIT ?")).
			WithArgs(id_success, 1).
			WillReturnRows(dataRole)
			
			mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE `users`.`role_id` = ?")).
			WithArgs(id_success).
			WillReturnRows(dataUser)

		data, err := repo.ShowById(id_success)
		assert.NoError(t, err)
		assert.Equal(t, uint(id_success), data.IdRole)
		assert.Equal(t, "Admin", data.NameRole)

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Failed to Show Role by Id", func(t *testing.T) {
		id_failed := 2
		mock.ExpectBegin()
		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `roles` WHERE id_role = ? ORDER BY `roles`.`id_role` LIMIT ?")).
			WithArgs(id_failed, 2).
			WillReturnRows(dataRole)
		mock.ExpectCommit()

		mock.ExpectBegin()
		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE `users`.`role_id` = ?")).
			WithArgs(id_failed).
			WillReturnRows(dataUser)
		mock.ExpectCommit()

		data, err := repo.ShowById(id_failed)
		assert.Error(t, err)
		assert.Nil(t, data)

		assert.Error(t, mock.ExpectationsWereMet())
	})
}

func TestUpdateRole(t *testing.T) {
	db, mock, err := SetupMockDb()
	assert.NoError(t, err)

	repo := RolesMysql(db)

	t.Run("Success Update Role", func(t *testing.T) {
		rolereq := &request.Roles{
			NameRole: "AdminUpdate",
		}
		id_succes := 1

		mock.ExpectBegin()

		mock.ExpectExec(regexp.QuoteMeta("UPDATE `roles` ")).
			WithArgs(rolereq.NameRole, id_succes).WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectCommit()

		err := repo.UpdateRole(id_succes, rolereq)
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Failed Update Role", func(t *testing.T) {
		rolereq := &request.Roles{
			NameRole: "AdminUpdate",
		}
		id_failed := 2

		mock.ExpectBegin()

		mock.ExpectExec(regexp.QuoteMeta("UPDATE `roles` ")).
			WithArgs(rolereq.NameRole, id_failed).WillReturnError(sqlmock.ErrCancelled)
		mock.ExpectRollback()

		err := repo.UpdateRole(id_failed, rolereq)
		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
