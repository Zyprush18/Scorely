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
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `roles` ")).WithArgs(sqlmock.AnyArg(), role.NameRole).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		err := repo.CreateRole(role)

		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())

	})

	t.Run("Failed Create Role", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `roles`")).WithArgs(sqlmock.AnyArg(), role.NameRole).WillReturnError(sqlmock.ErrCancelled)
		mock.ExpectRollback()

		err := repo.CreateRole(role)
		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())

	})

}
