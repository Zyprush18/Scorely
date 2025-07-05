package repouser

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

type repoTest struct {
	Name                                         string
	DataRows, DataRelation                       *sqlmock.Rows
	Count                                        func(search string, expcount int64) *sqlmock.ExpectedQuery
	FindDataAll                                  func(search string, perpage, offset int, data *sqlmock.Rows) *sqlmock.ExpectedQuery
	Find                                         func(iduser, idrole int, datauser, datarole *sqlmock.Rows)
	Request                                      *request.User
	Mocks                                        func(email string, idrole,iduser int)
	Search, Sort                                 string
	Page, Perpage, expectedCount, iduser, idrole int
	Err                                          bool
}

func TestGetAll(t *testing.T) {
	database, mock, err := SetupDBForUser()
	assert.NoError(t, err)
	repouser := NewUserDatabase(database)

	data := []repoTest{
		{
			Name: "Success Get All Data With Order Asceding",
			DataRows: sqlmock.NewRows([]string{
				"id_user",
				"email",
				"password",
				"role_id",
			}).AddRow(1, "Admin@gmail.com", "Admin124", 1).AddRow(2, "user@gmail.com", "user123", 2),
			Count: func(search string, expcount int64) *sqlmock.ExpectedQuery {
				return mock.ExpectQuery(regexp.QuoteMeta("SELECT count(*) FROM `users` WHERE email LIKE ?")).WithArgs("%" + search + "%").WillReturnRows(sqlmock.NewRows([]string{
					"count",
				}).AddRow(expcount))
			},
			FindDataAll: func(search string, perpage, offset int, data *sqlmock.Rows) *sqlmock.ExpectedQuery {
				if offset != 0 {
					return mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE email LIKE ? ORDER BY created_at ASC LIMIT ? OFFSET ?")).WithArgs("%"+search+"%", perpage, offset).WillReturnRows(data)
				}
				return mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE email LIKE ? ORDER BY created_at ASC LIMIT ?")).WithArgs("%"+search+"%", perpage).WillReturnRows(data)
			},
			Search:        "",
			Sort:          "ASC",
			Page:          1,
			Perpage:       10,
			expectedCount: 2,
			Err:           false,
		},
		{
			Name: "Success Get All Data With Order Desceding",
			DataRows: sqlmock.NewRows([]string{
				"id_user",
				"email",
				"password",
				"role_id",
			}).AddRow(2, "user@gmail.com", "user123", 2).AddRow(1, "Admin@gmail.com", "Admin124", 1),
			Count: func(search string, expcount int64) *sqlmock.ExpectedQuery {
				return mock.ExpectQuery(regexp.QuoteMeta("SELECT count(*) FROM `users` WHERE email LIKE ?")).WithArgs("%" + search + "%").WillReturnRows(sqlmock.NewRows([]string{
					"count",
				}).AddRow(expcount))
			},
			FindDataAll: func(search string, perpage, offset int, data *sqlmock.Rows) *sqlmock.ExpectedQuery {
				if offset != 0 {
					return mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE email LIKE ? ORDER BY created_at DESC LIMIT ? OFFSET ?")).WithArgs("%"+search+"%", perpage, offset).WillReturnRows(data)
				}
				return mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE email LIKE ? ORDER BY created_at DESC LIMIT ?")).WithArgs("%"+search+"%", perpage).WillReturnRows(data)
			},
			Search:        "",
			Sort:          "DESC",
			Page:          1,
			Perpage:       10,
			expectedCount: 2,
			Err:           false,
		},
		{
			Name: "Failed Get All Data",
			DataRows: sqlmock.NewRows([]string{
				"id_user",
				"email",
				"password",
				"role_id",
			}).AddRow(1, "Admin@gmail.com", "Admin124", 1).AddRow(2, "user@gmail.com", "user123", 2),
			Count: func(search string, expcount int64) *sqlmock.ExpectedQuery {
				return mock.ExpectQuery(regexp.QuoteMeta("SELECT count(*) FROM `users` WHERE email LIKE ?")).WithArgs("%" + search + "%").WillReturnRows(sqlmock.NewRows([]string{
					"count",
				}).AddRow(expcount))
			},
			FindDataAll: func(search string, perpage, offset int, data *sqlmock.Rows) *sqlmock.ExpectedQuery {
				return mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE email LIKE ? ORDER BY created_at ASC LIMIT ?")).WithArgs("%"+search+"%", perpage).WillReturnError(database.Error)
			},
			Search:        "",
			Sort:          "ASC",
			Page:          1,
			Perpage:       10,
			expectedCount: 0,
			Err:           true,
		},
	}

	for _, v := range data {
		t.Run(v.Name, func(t *testing.T) {
			offset := (v.Page - 1) * v.Perpage
			v.Count(v.Search, int64(v.expectedCount))
			v.FindDataAll(v.Search, v.Perpage, offset, v.DataRows)

			resp, count, err := repouser.GetAll(v.Search, v.Sort, v.Page, v.Perpage)
			if v.Err {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, int64(v.expectedCount), count)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestCreateUser(t *testing.T) {
	databases, mocks, err := SetupDBForUser()
	assert.NoError(t, err)
	repouser := NewUserDatabase(databases)

	data := []repoTest{
		{
			Name: "Success Create User",
			Request: &request.User{
				Email:    "Admin11@gmail.com",
				Password: "admin123",
				RoleId:   1,
			},
			Mocks: func(email string, idrole,iduser int) {
				mocks.ExpectBegin()
				mocks.ExpectExec(regexp.QuoteMeta("INSERT INTO `users`")).WithArgs(email, sqlmock.AnyArg(), idrole, sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 1))
				mocks.ExpectCommit()
			},
			Err: false,
		},
		{
			Name: "Failed Create User",
			Request: &request.User{
				Email:    "Admin1@gmail.com",
				Password: "admin123",
				RoleId:   1,
			},
			Mocks: func(email string, idrole,iduser int) {
				mocks.ExpectBegin()
				mocks.ExpectExec(regexp.QuoteMeta("INSERT INTO `users`")).WithArgs(email, sqlmock.AnyArg(), idrole, sqlmock.AnyArg()).WillReturnError(databases.Error)
				mocks.ExpectRollback()
			},
			Err: true,
		},
	}

	for _, v := range data {
		t.Run(v.Name, func(t *testing.T) {
			v.Mocks(v.Request.Email, int(v.Request.RoleId),v.iduser)
			err := repouser.Create(v.Request)
			if v.Err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.NoError(t, mocks.ExpectationsWereMet())
		})
	}
}

func TestShowUser(t *testing.T) {
	databases, mocks, err := SetupDBForUser()
	assert.NoError(t, err)
	repouser := NewUserDatabase(databases)

	data := []repoTest{
		{
			Name: "Success Show User",
			DataRows: sqlmock.NewRows([]string{
				"id_user",
				"email",
				"password",
				"role_id",
			}).AddRow(1, "Admin@gmail.com", "admin123", 1),
			DataRelation: sqlmock.NewRows([]string{
				"id_role",
				"name_role",
			}).AddRow(1, "Admin"),
			Find: func(iduser, idrole int, datauser, datarole *sqlmock.Rows) {
				mocks.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE id_user = ? ORDER BY `users`.`id_user` LIMIT ?")).WithArgs(iduser, 1).WillReturnRows(datauser)
				mocks.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `roles` WHERE `roles`.`id_role` = ?")).WithArgs(idrole).WillReturnRows(datarole)
			},
			iduser: 1,
			idrole: 1,
			Err:    false,
		},
		{
			Name: "Failed Show User",
			DataRows: sqlmock.NewRows([]string{
				"id_user",
				"email",
				"password",
				"role_id",
			}).AddRow(1, "admin@gmail.com", "admin123", 1),
			DataRelation: sqlmock.NewRows([]string{
				"id_role",
				"name_role",
			}).AddRow(1, "admin"),
			Find: func(iduser, idrole int, datauser, datarole *sqlmock.Rows) {
				mocks.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE id_user = ? ORDER BY `users`.`id_user` LIMIT ?")).WithArgs(iduser, 1).WillReturnError(gorm.ErrRecordNotFound)
			},
			iduser: 3,
			idrole: 1,
			Err:    true,
		},
	}
	for _, v := range data {
		t.Run(v.Name, func(t *testing.T) {
			v.Find(v.iduser, v.idrole, v.DataRows, v.DataRelation)
			data, err := repouser.Show(v.iduser)
			if v.Err {
				assert.Error(t, err)
				assert.Nil(t, data)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, uint(v.iduser), data.IdUser)
			}

			assert.NoError(t, mocks.ExpectationsWereMet())
		})
	}
}

func TestUpdateUser(t *testing.T) {
	database, mock, err := SetupDBForUser()
	assert.NoError(t, err)
	repouser := NewUserDatabase(database)

	data := []repoTest{
		{
			Name: "Success Update User",
			DataRows: sqlmock.NewRows([]string{
				"id_user",
				"email",
				"password",
				"role_id",
			}).AddRow(1, "Admin@gmail.com", "admin123", 1),
			Request: &request.User{
				Email: "AdminUpdate@gmail.com",
			},
			Find: func(iduser, idrole int, datauser, datarole *sqlmock.Rows) {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE id_user = ? ORDER BY `users`.`id_user` LIMIT ?")).WithArgs(iduser, 1).WillReturnRows(datauser)
			},
			Mocks: func(email string, idrole,iduser int) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta("UPDATE `users`")).WithArgs(email, sqlmock.AnyArg(), idrole).WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
			iduser: 1,
			idrole: 1,
			Err:    false,
		},
		{
			Name: "Failed Update User",
			DataRows: sqlmock.NewRows([]string{
				"id_user",
				"email",
				"password",
				"role_id",
			}).AddRow(1, "Admin@gmail.com", "admin123", 1),
			Request: &request.User{
				Email: "admins@gmail.com",
			},
			Find: func(iduser, idrole int, datauser, datarole *sqlmock.Rows) {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE id_user = ? ORDER BY `users`.`id_user` LIMIT ?")).WithArgs(iduser, 1).WillReturnRows(datauser)
			},
			Mocks: func(email string, idrole,iduser int) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta("UPDATE `users`")).WithArgs(email, sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnError(sqlmock.ErrCancelled)
				mock.ExpectRollback()
			},
			iduser: 2,
			Err:    true,
		},
	}

	for _, v := range data {
		t.Run(v.Name, func(t *testing.T) {
			v.Find(v.iduser, v.idrole, v.DataRows, nil)
			v.Mocks(v.Request.Email, v.idrole,v.iduser)

			err := repouser.Update(v.iduser, v.Request)
			if v.Err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestDeleteUser(t *testing.T) {
	database, mock, err := SetupDBForUser()
	assert.NoError(t, err)
	repouser := NewUserDatabase(database)

	data := []repoTest{
		{
			Name: "Success Delete User",
			DataRows: sqlmock.NewRows([]string{
				"id_user",
				"email",
				"password",
				"role_id",
			}).AddRow(1, "Admin@gmail.com", "admin123", 1),
			Find: func(iduser, idrole int, datauser, datarole *sqlmock.Rows) {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE id_user = ? ORDER BY `users`.`id_user` LIMIT ?")).WithArgs(iduser, 1).WillReturnRows(datauser)
			},
			Mocks: func(email string, idrole,iduser int) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta("DELETE FROM `users`")).WithArgs(iduser).WillReturnResult(sqlmock.NewResult(0,1))
				mock.ExpectCommit()
			},
			iduser: 1,
			Err: false,
		},
		{
			Name: "Failed Delete User",
			DataRows: sqlmock.NewRows([]string{
				"id_user",
				"email",
				"password",
				"role_id",
			}).AddRow(1, "Admin@gmail.com", "admin123", 1),
			Find: func(iduser, idrole int, datauser, datarole *sqlmock.Rows) {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE id_user = ? ORDER BY `users`.`id_user` LIMIT ?")).WithArgs(iduser, 1).WillReturnRows(datauser)
			},
			Mocks: func(email string, idrole,iduser int) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta("DELETE FROM `users`")).WithArgs(iduser).WillReturnError(sqlmock.ErrCancelled)
				mock.ExpectRollback()
			},
			iduser: 1,
			Err: true,
		},
	}

	for _, v := range data {
		t.Run(v.Name,func(t *testing.T) {
			v.Find(v.iduser,0,v.DataRows,nil)
			v.Mocks("", 0,v.iduser)

			err:= repouser.Delete(v.iduser)
			if v.Err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)

			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}