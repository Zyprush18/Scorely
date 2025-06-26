package reporole

import (
	"regexp"

	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/stretchr/testify/assert"
)

type tableroleTest struct {
	Name          string
	DataRows      *sqlmock.Rows
	Search, Sort  string
	Page, Perpage int
	ExpectedCount int64
	Err           bool
}

func TestGetAllData(t *testing.T) {
	db, mock, err := SetupMockDb()
	assert.NoError(t, err)
	repo := RolesMysql(db)

	data := []tableroleTest{
		{
			Name: "Success Get All Data",
			DataRows: sqlmock.NewRows([]string{
				"id_role",
				"name_role",
			}).AddRow(1, "Admin").AddRow(2, "User"),
			Search:        "",
			Sort:          "ASC",
			Page:          1,
			Perpage:       10,
			ExpectedCount: 2,
			Err:           false,
		},
		{
			Name: "Failed Get All Data",
			DataRows: sqlmock.NewRows([]string{
				"id_role",
				"name_role",
			}).AddRow(1, "Admin").AddRow(2, "User"),
			Search:        "",
			Sort:          "ASC",
			Page:          1,
			Perpage:       10,
			ExpectedCount: 2,
			Err:           true,
		},
	}

	for _, v := range data {
		offset := (v.Page - 1) * v.Perpage

		t.Run(v.Name, func(t *testing.T) {
			if v.Err {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT count(*) FROM `roles` WHERE name_role LIKE ?")).WithArgs("%" + v.Search + "%").WillReturnRows(sqlmock.NewRows([]string{
					"count",
				}).AddRow(v.ExpectedCount))
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `roles` WHERE name_role LIKE ? ORDER BY created_at ASC LIMIT ?")).WillReturnError(db.Error)
				resp, count, err := repo.GetAllDataRole(v.Search, v.Sort, v.Page, v.Perpage)
				assert.Error(t, err)
				assert.Nil(t, resp)
				assert.NotEqual(t, int64(v.ExpectedCount), count)
				assert.NoError(t, mock.ExpectationsWereMet())
			} else {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT count(*) FROM `roles` WHERE name_role LIKE ?")).WithArgs("%" + v.Search + "%").WillReturnRows(sqlmock.NewRows([]string{
					"count",
				}).AddRow(v.ExpectedCount))

				if offset == 0 {
					mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `roles` WHERE name_role LIKE ? ORDER BY created_at ASC LIMIT ?")).
						WithArgs("%"+v.Search+"%", v.Perpage).
						WillReturnRows(v.DataRows)
				} else {
					mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `roles` WHERE name_role LIKE ? ORDER BY created_at ASC LIMIT ? OFFSET ?")).
						WithArgs("%"+v.Search+"%", v.Perpage, offset).
						WillReturnRows(v.DataRows)
				}

				resp, count, err := repo.GetAllDataRole(v.Search, v.Sort, v.Page, v.Perpage)
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, int64(v.ExpectedCount), count)
				assert.NoError(t, mock.ExpectationsWereMet())
			}
		})
	}
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
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `roles` ")).WithArgs(role.NameRole, sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		err := repo.CreateRole(role)

		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())

	})

	t.Run("Failed Create Role", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `roles`")).WithArgs(role.NameRole, sqlmock.AnyArg()).WillReturnError(sqlmock.ErrCancelled)
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

	dataRole := sqlmock.NewRows([]string{
		"id_role",
		"name_role",
	}).AddRow(1, "Admin")
	t.Run("Success Update Role", func(t *testing.T) {

		rolereq := &request.Roles{
			NameRole: "AdminUpdate",
		}
		id_succes := 1

		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `roles` WHERE id_role = ? ORDER BY `roles`.`id_role` LIMIT ?")).
			WithArgs(id_succes, 1).
			WillReturnRows(dataRole)

		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("UPDATE `roles` ")).
			WithArgs(rolereq.NameRole,sqlmock.AnyArg(), id_succes).WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectCommit()

		err := repo.UpdateRole(id_succes, rolereq)
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Failed Update Role", func(t *testing.T) {
		rolereq := &request.Roles{
			NameRole: "AdminUpdate1",
		}
		id_failed := 2

		mock.ExpectBegin()
		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `roles` WHERE id_role = ? ORDER BY `roles`.`id_role` LIMIT ?")).
			WithArgs(id_failed, 2).
			WillReturnRows(dataRole)
		mock.ExpectCommit()

		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("UPDATE `roles` ")).
			WithArgs(rolereq.NameRole, id_failed,sqlmock.AnyArg()).WillReturnError(sqlmock.ErrCancelled)
		mock.ExpectRollback()

		err := repo.UpdateRole(id_failed, rolereq)
		assert.Error(t, err)
		assert.Error(t, mock.ExpectationsWereMet())
	})
}

func TestDeleteRole(t *testing.T) {
	db, mock, err := SetupMockDb()
	assert.NoError(t, err)

	repo := RolesMysql(db)

	dataRole := sqlmock.NewRows([]string{
		"id_role",
		"name_role",
	}).AddRow(1, "Admin")

	t.Run("Success Delete Role", func(t *testing.T) {

		id_succes := 1

		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `roles` WHERE id_role = ? ORDER BY `roles`.`id_role` LIMIT ?")).
			WithArgs(id_succes, 1).
			WillReturnRows(dataRole)

		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("DELETE FROM `roles` ")).
			WithArgs(id_succes).WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectCommit()

		err := repo.DeleteRole(id_succes)
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Failed Delete Role", func(t *testing.T) {

		id_failed := 2

		mock.ExpectBegin()
		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `roles` WHERE id_role = ? ORDER BY `roles`.`id_role` LIMIT ?")).
			WithArgs(id_failed, 2).
			WillReturnRows(dataRole)
		mock.ExpectCommit()

		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("DELETE FROM `roles` ")).
			WithArgs(id_failed).WillReturnError(sqlmock.ErrCancelled)
		mock.ExpectRollback()

		err := repo.DeleteRole(id_failed)
		assert.Error(t, err)
		assert.Error(t, mock.ExpectationsWereMet())
	})
}
