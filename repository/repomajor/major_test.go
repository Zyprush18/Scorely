package repomajor

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)


type testRepoMajor struct {
	Name string
	Id   int
	DataRows  *sqlmock.Rows
	CountData func (search string, expectationcount int64) *sqlmock.ExpectedQuery
	FindDataAll func (src string, perpage,offset int, data *sqlmock.Rows) *sqlmock.ExpectedQuery
	FindDataById func (id int, data *sqlmock.Rows)
	ReqMock func (datareq *request.Majors)
	ReqMajor	*request.Majors
	Page,Perpage int
	Search,Sort string
	ExpCount int64
	IsErr bool
}	

func TestGetAll(t *testing.T)  {
	db,mock,err := NewDbMock()
	assert.NoError(t, err)
	repomajor := ConnectDb(db)

	data := []testRepoMajor{
		{
			Name: "Success Get All Data: ASC",
			DataRows: sqlmock.NewRows([]string{
				"id_major",
				"major",
				"major_abbreviation",
			}).AddRow(1, "Computer Engineer", "CE").AddRow(2, "Information System", "SI"),
			CountData: func(search string, expectationcount int64) *sqlmock.ExpectedQuery {
				return mock.ExpectQuery(regexp.QuoteMeta("SELECT count(*) FROM `majors` WHERE major LIKE ?")).WithArgs("%"+search+"%").WillReturnRows(sqlmock.NewRows([]string{
					"count",
				}).AddRow(expectationcount))
			},
			FindDataAll: func(src string, perpage, offset int, data *sqlmock.Rows) *sqlmock.ExpectedQuery {
				if offset != 0 {
					return mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `majors` WHERE major LIKE ? ORDER BY created_at ASC LIMIT ? OFFSET ?")).WithArgs("%"+src+"%", perpage,offset).WillReturnRows(data)
				}
				return mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `majors` WHERE major LIKE ? ORDER BY created_at ASC LIMIT ?")).WithArgs("%"+src+"%", perpage).WillReturnRows(data)
			},
			Page: 1,
			Perpage: 10,
			Search: "",
			Sort: "ASC",
			ExpCount: 2,
			IsErr: false,
		},
		{
			Name: "Success Get All Data: DESC",
			DataRows: sqlmock.NewRows([]string{
				"id_major",
				"major",
				"major_abbreviation",
			}).AddRow(1, "Computer Engineer", "CE").AddRow(2, "Information System", "SI"),
			CountData: func(search string, expectationcount int64) *sqlmock.ExpectedQuery {
				return mock.ExpectQuery(regexp.QuoteMeta("SELECT count(*) FROM `majors` WHERE major LIKE ?")).WithArgs("%"+search+"%").WillReturnRows(sqlmock.NewRows([]string{
					"count",
				}).AddRow(expectationcount))
			},
			FindDataAll: func(src string, perpage, offset int, data *sqlmock.Rows) *sqlmock.ExpectedQuery {
				if offset != 0 {
					return mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `majors` WHERE major LIKE ? ORDER BY created_at DESC LIMIT ? OFFSET ?")).WithArgs("%"+src+"%", perpage,offset).WillReturnRows(data)
				}
				return mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `majors` WHERE major LIKE ? ORDER BY created_at DESC LIMIT ?")).WithArgs("%"+src+"%", perpage).WillReturnRows(data)
			},
			Page: 1,
			Perpage: 10,
			Search: "",
			Sort: "DESC",
			ExpCount: 2,
			IsErr: false,
		},
		{
			Name: "Failed Get All Data",
			DataRows: sqlmock.NewRows([]string{
				"id_major",
				"major",
				"major_abbreviation",
			}).AddRow(1, "Computer Engineer", "CE").AddRow(2, "Information System", "SI"),
			CountData: func(search string, expectationcount int64) *sqlmock.ExpectedQuery {
				return mock.ExpectQuery(regexp.QuoteMeta("SELECT count(*) FROM `majors` WHERE major LIKE ?")).WithArgs("%"+search+"%").WillReturnRows(sqlmock.NewRows([]string{
					"count",
				}).AddRow(expectationcount))
			},
			FindDataAll: func(src string, perpage, offset int, data *sqlmock.Rows) *sqlmock.ExpectedQuery {
				if offset != 0 {
					return mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `majors` WHERE major LIKE ? ORDER BY created_at ASC LIMIT ? OFFSET ?")).WithArgs("%"+src+"%", perpage,offset).WillReturnError(db.Error)
				}
				return mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `majors` WHERE major LIKE ? ORDER BY created_at ASC LIMIT ?")).WithArgs("%"+src+"%", perpage).WillReturnError(db.Error)
			},
			Page: 1,
			Perpage: 10,
			Search: "",
			Sort: "ASC",
			ExpCount: 2,
			IsErr: true,
		},
	}

	for _, v := range data {
		t.Run(v.Name, func(t *testing.T) {
			offset := (v.Page - 1) * v.Perpage
			v.CountData(v.Search,v.ExpCount)
			v.FindDataAll(v.Search,v.Perpage,offset, v.DataRows)

			resp, count, err := repomajor.GetAllData(v.Search,v.Sort,v.Page,v.Perpage)
			if v.IsErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, v.ExpCount, count)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}

}

func TestCreate(t *testing.T)  {
	db,mock,err := NewDbMock()
	assert.NoError(t, err)
	repomajor := ConnectDb(db)

	data := []testRepoMajor{
		{
			Name: "Success Create New Major",
			ReqMajor: &request.Majors{
				Major: "Information Technology",
				MajorAbbreviation: "IT",
			},
			ReqMock: func(datareq *request.Majors) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `majors`")).WithArgs(datareq.Major, datareq.MajorAbbreviation, sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1,1))
				mock.ExpectCommit()
			},
			IsErr: false,
		},
		{
			Name: "Failed Create New Major",
			ReqMajor: &request.Majors{
				Major: "Information Technology",
				MajorAbbreviation: "IT",
			},
			ReqMock: func(datareq *request.Majors) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `majors`")).WithArgs(datareq.Major, datareq.MajorAbbreviation, sqlmock.AnyArg()).WillReturnError(db.Error)
				mock.ExpectRollback()
			},
			IsErr: true,
		},
	}

	for _, v := range data {
		t.Run(v.Name,func(t *testing.T) {
			v.ReqMock(v.ReqMajor)
			err := repomajor.Create(v.ReqMajor)
			if v.IsErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestShow(t *testing.T)  {
	db, mock, err := NewDbMock()
	assert.NoError(t, err)
	repomajor := ConnectDb(db)
	data := []testRepoMajor{
		{
			Name: "Success Show Data",
			DataRows: sqlmock.NewRows([]string{
				"id_major",
				"major",
				"major_abbreviation",
			}).AddRow(1, "Informatics Engineering", "IE"),
			Id: 1,
			FindDataById: func(id int, data *sqlmock.Rows) {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `majors` WHERE id_major = ? ORDER BY `majors`.`id_major` LIMIT ?")).WithArgs(id, 1).WillReturnRows(data)
			},
			IsErr: false,
		},
		{
			Name: "Failed Show Data",
			DataRows: sqlmock.NewRows([]string{
				"id_major",
				"major",
				"major_abbreviation",
			}).AddRow(1, "Informatics Engineering", "IE"),
			Id: 2,
			FindDataById: func(id int, data *sqlmock.Rows) {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `majors` WHERE id_major = ? ORDER BY `majors`.`id_major` LIMIT ?")).WithArgs(id,1).WillReturnError(gorm.ErrRecordNotFound)
			},
			IsErr: true,
		},
	}

	for _, v := range data {
		t.Run(v.Name, func(t *testing.T) {
			v.FindDataById(v.Id, v.DataRows)
			resp, err := repomajor.ShowById(v.Id)
			if v.IsErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestUpdate(t *testing.T)  {
	db,mock,err := NewDbMock()
	assert.NoError(t, err)
	repomajor := ConnectDb(db)
	data := []testRepoMajor{
		{
			Name: "Success Update Major",
			ReqMajor: &request.Majors{
				Major: "System Information",
				MajorAbbreviation: "SI",
			},
			DataRows: sqlmock.NewRows([]string{
				"id_major",
				"major",
				"major_abbreviation",
			}).AddRow(1, "Informatics Engineering","IE"),
			FindDataById: func(id int, data *sqlmock.Rows) {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `majors` WHERE id_major = ? ORDER BY `majors`.`id_major` LIMIT ?")).WithArgs(id, 1).WillReturnRows(data)
			},
			ReqMock: func(datareq *request.Majors) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta("UPDATE `majors`")).WithArgs(datareq.Major,datareq.MajorAbbreviation,sqlmock.AnyArg(), 1).WillReturnResult(sqlmock.NewResult(0,1))
				mock.ExpectCommit()
			},
			Id: 1,
			IsErr: false,
		},
		{
			Name: "Failed Update Major: Not Found",
			ReqMajor: &request.Majors{
				Major: "System Information",
				MajorAbbreviation: "SI",
			},
			DataRows: sqlmock.NewRows([]string{
				"id_major",
				"major",
				"major_abbreviation",
			}).AddRow(2, "Informatics Engineering","IE"),
			FindDataById: func(id int, data *sqlmock.Rows) {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `majors` WHERE id_major = ? ORDER BY `majors`.`id_major` LIMIT ?")).WithArgs(id, 1).WillReturnError(gorm.ErrRecordNotFound)
			},
			ReqMock: func(datareq *request.Majors) {},
			Id: 2,
			IsErr: true,
		},
		{
			Name: "Failed Update Major: Database error",
			ReqMajor: &request.Majors{
				Major: "System Information",
				MajorAbbreviation: "SI",
			},
			DataRows: sqlmock.NewRows([]string{
				"id_major",
				"major",
				"major_abbreviation",
			}).AddRow(2, "Informatics Engineering","IE"),
			FindDataById: func(id int, data *sqlmock.Rows) {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `majors` WHERE id_major = ? ORDER BY `majors`.`id_major` LIMIT ?")).WithArgs(id, 1).WillReturnRows(data)
			},
			ReqMock: func(datareq *request.Majors) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta("UPDATE `majors`")).WithArgs(datareq.Major,datareq.MajorAbbreviation,sqlmock.AnyArg(), 2).WillReturnError(db.Error)
				mock.ExpectRollback()
			},
			Id: 2,
			IsErr: true,
		},
	}
	for _, v := range data {
		t.Run(v.Name, func(t *testing.T) {
			v.FindDataById(v.Id,v.DataRows)
			v.ReqMock(v.ReqMajor)

			err := repomajor.Updates(v.Id, v.ReqMajor)
			if v.IsErr {
				assert.Error(t, err)
			}else{
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestDelete(t *testing.T)  {
	db, mock, err := NewDbMock()
	assert.NoError(t, err)
	repomajor := ConnectDb(db)
	data := []testRepoMajor{
		{
			Name: "Success Delete Major",
			Id: 1,
			DataRows: sqlmock.NewRows([]string{
				"id_major",
				"major",
				"major_abbreviation",
			}).AddRow(1, "System Information", "SI"),
			FindDataById: func(id int, data *sqlmock.Rows) {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `majors` WHERE id_major = ? ORDER BY `majors`.`id_major` LIMIT ?")).WithArgs(id, 1).WillReturnRows(data)
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta("DELETE FROM `majors`")).WithArgs(id).WillReturnResult(sqlmock.NewResult(0,1))
				mock.ExpectCommit()
			},
			IsErr: false,
		},
		{
			Name: "Failed Delete Major: Not Found",
			Id: 2,
			DataRows: sqlmock.NewRows([]string{
				"id_major",
				"major",
				"major_abbreviation",
			}).AddRow(1, "System Information", "SI"),
			FindDataById: func(id int, data *sqlmock.Rows) {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `majors` WHERE id_major = ? ORDER BY `majors`.`id_major` LIMIT ?")).WithArgs(id, 1).WillReturnError(gorm.ErrRecordNotFound)
			},
			IsErr: true,
		},
		{
			Name: "Failed Delete Major: Database Error",
			Id: 3,
			DataRows: sqlmock.NewRows([]string{
				"id_major",
				"major",
				"major_abbreviation",
			}).AddRow(3, "System Information", "SI"),
			FindDataById: func(id int, data *sqlmock.Rows) {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `majors` WHERE id_major = ? ORDER BY `majors`.`id_major` LIMIT ?")).WithArgs(id, 1).WillReturnRows(data)
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta("DELETE FROM `majors`")).WithArgs(id).WillReturnError(db.Error)
				mock.ExpectRollback()
			},
			IsErr: true,
		},
	}

	for _, v := range data {
		t.Run(v.Name, func(t *testing.T) {
			v.FindDataById(v.Id, v.DataRows)
			err := repomajor.Deletes(v.Id)
			if v.IsErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}