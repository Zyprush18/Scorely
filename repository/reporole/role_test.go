package reporole

import (
	"regexp"

	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/stretchr/testify/assert"
)

func TestCreateRole(t *testing.T) {
	db, mock, err := SetupMockDb()
	assert.NoError(t, err)

	repo := RolesMysql(db)

	role := &request.Roles{
		NameRole: "Admin",
	}

	t.Run("Succes Create Role", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `roles` ")).WithArgs(sqlmock.AnyArg(), role.NameRole, sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		err := repo.CreateRole(role)

		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())

	})

	t.Run("Failed Create Role", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `roles`")).WithArgs(sqlmock.AnyArg(), role.NameRole, sqlmock.AnyArg()).WillReturnError(sqlmock.ErrCancelled)
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

	t.Run("Success to Show Role By Id", func(t *testing.T) {
		id_success := 1

		mock.ExpectQuery("SELECT").WillReturnRows(dataRole)

		data, err := repo.ShowById(id_success)
		assert.NoError(t, err)
		assert.Equal(t, uint(id_success), data.IdRole)
		assert.Equal(t, "Admin", data.NameRole)

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Failed to Show Role by Id", func(t *testing.T) {
		id_failed := 2

		mock.ExpectQuery("SELECT").WillReturnRows(dataRole)

		data, err := repo.ShowById(id_failed)
		assert.Error(t, err)
		assert.Nil(t, data)

		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
