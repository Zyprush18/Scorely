package reporole

import (
	"errors"
	"regexp"

	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/stretchr/testify/assert"
)

type tableRoleTest struct {
	Name          string
	DataRows      *sqlmock.Rows
	Count         func(search string, expcount int64) *sqlmock.ExpectedQuery
	FindData      func(search string, perpage, offset int, data *sqlmock.Rows) *sqlmock.ExpectedQuery
	Search, Sort  string
	Page, Perpage int
	ExpectedCount int64
	RequestRole   *request.Roles
	MockExec      func(name string)
	Id            int
	Err           bool
}

func TestGetAllData(t *testing.T) {
	db, mock, err := SetupMockDb()
	assert.NoError(t, err)
	repo := RolesMysql(db)

	data := []tableRoleTest{
		{
			Name: "Success Get All Data",
			DataRows: sqlmock.NewRows([]string{
				"id_role",
				"name_role",
			}).AddRow(1, "Admin").AddRow(2, "User"),
			Count: func(search string, expcount int64) *sqlmock.ExpectedQuery {
				return mock.ExpectQuery(regexp.QuoteMeta("SELECT count(*) FROM `roles` WHERE name_role LIKE ?")).WithArgs("%" + search + "%").WillReturnRows(sqlmock.NewRows([]string{
					"count",
				}).AddRow(expcount))
			},
			FindData: func(search string, perpage, offset int, data *sqlmock.Rows) *sqlmock.ExpectedQuery {
				if offset == 0 {
					return mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `roles` WHERE name_role LIKE ? ORDER BY created_at ASC LIMIT ?")).
						WithArgs("%"+search+"%", perpage).
						WillReturnRows(data)
				}
				return mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `roles` WHERE name_role LIKE ? ORDER BY created_at ASC LIMIT ? OFFSET ?")).
					WithArgs("%"+search+"%", perpage, offset).
					WillReturnRows(data)

			},
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
			Count: func(search string, expcount int64) *sqlmock.ExpectedQuery {
				return mock.ExpectQuery(regexp.QuoteMeta("SELECT count(*) FROM `roles` WHERE name_role LIKE ?")).WithArgs("%" + search + "%").WillReturnRows(sqlmock.NewRows([]string{
					"count",
				}).AddRow(expcount))
			},
			FindData: func(search string, perpage, offset int, data *sqlmock.Rows) *sqlmock.ExpectedQuery {
				return mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `roles` WHERE name_role LIKE ? ORDER BY created_at ASC LIMIT ?")).WillReturnError(db.Error)
			},
			Search:        "",
			Sort:          "ASC",
			Page:          1,
			Perpage:       10,
			ExpectedCount: 0,
			Err:           true,
		},
	}

	for _, v := range data {
		t.Run(v.Name, func(t *testing.T) {
			offset := (v.Page - 1) * v.Perpage

			v.Count(v.Search, v.ExpectedCount)
			v.FindData(v.Search, v.Perpage, offset, v.DataRows)

			resp, count, err := repo.GetAllDataRole(v.Search, v.Sort, v.Page, v.Perpage)
			if v.Err {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
			}
			assert.Equal(t, int64(v.ExpectedCount), count)
			assert.NoError(t, mock.ExpectationsWereMet())
		})

	}
}

func TestCreateRole(t *testing.T) {
	db, mock, err := SetupMockDb()
	assert.NoError(t, err)

	repo := RolesMysql(db)

	data := []tableRoleTest{
		{
			Name: "Success Create User",
			RequestRole: &request.Roles{
				NameRole: "Admin",
			},
			MockExec: func(name string) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `roles`")).WithArgs(name, sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			Err: false,
		},
		{
			Name: "Failed Create User",
			RequestRole: &request.Roles{
				NameRole: "User",
			},
			MockExec: func(name string) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `roles`")).WithArgs(name, sqlmock.AnyArg()).WillReturnError(sqlmock.ErrCancelled)
				mock.ExpectRollback()
			},
			Err: true,
		},
	}

	for _, v := range data {
		t.Run(v.Name, func(t *testing.T) {
			v.MockExec(v.RequestRole.NameRole)
			err := repo.CreateRole(v.RequestRole)
			if v.Err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestShowRole(t *testing.T) {
	db, mock, err := SetupMockDb()
	assert.NoError(t, err)
	repo := RolesMysql(db)

	data := []tableRoleTest{
		{
			Name: "Success Show Role By id",
			DataRows: sqlmock.NewRows([]string{
				"id_role",
				"name_role",
			}).AddRow(1, "Admin"),
			FindData: func(search string, perpage, offset int, data *sqlmock.Rows) *sqlmock.ExpectedQuery {
				return mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `roles` WHERE id_role = ? ORDER BY `roles`.`id_role` LIMIT ?")).
					WithArgs(1, 1).
					WillReturnRows(data)
			},
			Id:  1,
			Err: false,
		},
		{
			Name: "Failed Show Role By id",
			DataRows: sqlmock.NewRows([]string{
				"id_role",
				"name_role",
			}).AddRow(1, "Admin"),
			FindData: func(search string, perpage, offset int, data *sqlmock.Rows) *sqlmock.ExpectedQuery {
				return mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `roles` WHERE id_role = ? ORDER BY `roles`.`id_role` LIMIT ?")).
					WithArgs(2, 1).
					WillReturnError(errors.New("Not Found id: 2"))
			},
			Id:  2,
			Err: true,
		},
	}

	for _, v := range data {
		t.Run(v.Name, func(t *testing.T) {
			v.FindData("", 0, 0, v.DataRows)
			resp, err := repo.ShowById(v.Id)

			if v.Err {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, uint(v.Id), resp.IdRole)
				assert.Equal(t, "Admin", resp.NameRole)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestUpdateRole(t *testing.T) {
	db, mock, err := SetupMockDb()
	assert.NoError(t, err)

	repo := RolesMysql(db)

	data := []tableRoleTest{
		{
			Name: "Success Update Role",
			DataRows: sqlmock.NewRows([]string{
				"id_role",
				"name_role",
			}).AddRow(1, "Admin"),
			FindData: func(search string, perpage, offset int, data *sqlmock.Rows) *sqlmock.ExpectedQuery {
				return mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `roles` WHERE id_role = ? ORDER BY `roles`.`id_role` LIMIT ?")).
					WithArgs(1, 1).
					WillReturnRows(data)
			},
			RequestRole: &request.Roles{
				NameRole: "AdminUpdate",
			},
			MockExec: func(name string) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta("UPDATE `roles` ")).
					WithArgs(name, sqlmock.AnyArg(), 1).WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
			Id:  1,
			Err: false,
		},
		{
			Name: "Failed Update Role",
			DataRows: sqlmock.NewRows([]string{
				"id_role",
				"name_role",
			}).AddRow(1, "Admin"),
			FindData: func(search string, perpage, offset int, data *sqlmock.Rows) *sqlmock.ExpectedQuery {
				return mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `roles` WHERE id_role = ? ORDER BY `roles`.`id_role` LIMIT ?")).
					WithArgs(2, 1).
					WillReturnError(errors.New("Not Found id: 2"))
			},
			RequestRole: &request.Roles{
				NameRole: "UserUpdate",
			},
			MockExec: nil,
			Id:       2,
			Err:      true,
		},
	}

	for _, v := range data {
		t.Run(v.Name, func(t *testing.T) {
			v.FindData("", 0, 0, v.DataRows)

			if v.MockExec != nil {
				v.MockExec(v.RequestRole.NameRole)
			}

			err := repo.UpdateRole(v.Id, v.RequestRole)
			if v.Err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestDeleteRole(t *testing.T) {
	db, mock, err := SetupMockDb()
	assert.NoError(t, err)

	repo := RolesMysql(db)

	data := []tableRoleTest{
		{
			Name: "Success Delete Role",
			DataRows: sqlmock.NewRows([]string{
				"id_role",
				"name_role",
			}).AddRow(1, "Admin"),
			FindData: func(search string, perpage, offset int, data *sqlmock.Rows) *sqlmock.ExpectedQuery {
				return mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `roles` WHERE id_role = ? ORDER BY `roles`.`id_role` LIMIT ?")).
					WithArgs(1, 1).
					WillReturnRows(data)
			},
			MockExec: func(name string) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta("DELETE FROM `roles` ")).
					WithArgs(1).WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
			Id: 1,
			Err: false,
		},
		{
			Name: "Failed Delete Role",
			DataRows: sqlmock.NewRows([]string{
				"id_role",
				"name_role",
			}).AddRow(1, "Admin"),
			FindData: func(search string, perpage, offset int, data *sqlmock.Rows) *sqlmock.ExpectedQuery {
				return mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `roles` WHERE id_role = ? ORDER BY `roles`.`id_role` LIMIT ?")).
					WithArgs(2, 1).
					WillReturnError(errors.New("Not Found id: 2"))
			},
			MockExec: nil,
			Id: 2,
			Err: true,
		},
	}

	for _, v := range data {
		t.Run(v.Name, func(t *testing.T) {
			v.FindData("",0,0,v.DataRows)
			if v.MockExec != nil {
				v.MockExec("")
			}

			err := repo.DeleteRole(v.Id)
			if v.Err {
				assert.Error(t, err)
			}else{
				assert.NoError(t, err)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
