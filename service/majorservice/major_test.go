package majorservice

import (
	"errors"
	"testing"
	"time"

	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/models/response"
	"github.com/stretchr/testify/assert"
)

type testMajorService struct {
	Name          string
	ResponseAll   []response.Majors
	ReqMajor      *request.Majors
	RespMajor     *response.Majors
	Search, Sort  string
	Page, Perpage int
	expectedCount int
	Id			  int
	ExpectedErr   error
}

func TestGetAll(t *testing.T) {
	mockmajor := RepoMajorMock{}
	servicemajor := RepoMajorConn(&mockmajor)
	data := []testMajorService{
		{
			Name: "Success Get All",
			ResponseAll: []response.Majors{
				{
					IdMajor:           1,
					Major:             "System Information",
					MajorAbbreviation: "SI",
					Models: helper.Models{
						CreatedAt: time.Now(),
						UpdatedAt: time.Now(),
					},
				},
				{
					IdMajor:           2,
					Major:             "Informatics Engineering",
					MajorAbbreviation: "IE",
					Models: helper.Models{
						CreatedAt: time.Now(),
						UpdatedAt: time.Now(),
					},
				},
			},
			Search:        "",
			Sort:          "asc",
			Page:          1,
			Perpage:       10,
			expectedCount: 2,
			ExpectedErr:   nil,
		},
		{
			Name:          "Failed Get All",
			ResponseAll:   nil,
			Search:        "",
			Sort:          "asc",
			Page:          1,
			Perpage:       10,
			expectedCount: 0,
			ExpectedErr:   errors.New("Database Is Refused"),
		},
	}

	for _, v := range data {
		t.Run(v.Name, func(t *testing.T) {
			// hapus cache testing sebelumnya
			mockmajor.ExpectedCalls = nil
			mockmajor.Calls = nil

			mockmajor.On("GetAllData", v.Search, v.Sort, v.Page, v.Perpage).Return(v.ResponseAll, v.expectedCount, v.ExpectedErr)
			resp, count, err := servicemajor.GetAllMajor(v.Search, v.Sort, v.Page, v.Perpage)
			if v.ExpectedErr != nil {
				assert.Error(t, err)
				assert.Nil(t, resp)
				assert.Equal(t, int64(v.expectedCount), count)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, int64(v.expectedCount), count)
			}
			mockmajor.AssertExpectations(t)
		})
	}
}

func TestCreate(t *testing.T) {
	mockmajor := RepoMajorMock{}
	majorservice := RepoMajorConn(&mockmajor)
	data := []testMajorService{
		{
			Name: "Success Create Major",
			ReqMajor: &request.Majors{
				Major:             "System Information",
				MajorAbbreviation: "SI",
			},
			ExpectedErr: nil,
		},
		{
			Name:        "Failed Create Major",
			ReqMajor:    &request.Majors{},
			ExpectedErr: errors.New("Database Is Refused"),
		},
	}

	for _, v := range data {
		t.Run(v.Name, func(t *testing.T) {
			mockmajor.On("Create", v.ReqMajor).Return(v.ExpectedErr)
			err := majorservice.CreateMajor(v.ReqMajor)
			if v.ExpectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			mockmajor.AssertExpectations(t)
		})
	}
}

func TestShow(t *testing.T) {
	mockmajor := RepoMajorMock{}
	servicemajor := RepoMajorConn(&mockmajor)
	data := []testMajorService{
		{
			Name: "Success Show",
			RespMajor: &response.Majors{
				IdMajor:           1,
				Major:             "System Information",
				MajorAbbreviation: "SI",
				Models: helper.Models{
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
			},
			Id: 1,
			ExpectedErr: nil,
		},
		{
			Name: "Failed Show",
			RespMajor: nil,
			Id: 2,
			ExpectedErr: errors.New("Not Found Id: 2"),
		},
	}

	for _, v := range data {
		t.Run(v.Name, func(t *testing.T) {
			mockmajor.On("ShowById", v.Id).Return(v.RespMajor, v.ExpectedErr)
			resp, err := servicemajor.ShowMajor(v.Id)
			if v.ExpectedErr != nil {
				assert.Error(t, err)
				assert.Nil(t, resp)
			}else{
				assert.NoError(t, err)
				assert.NotNil(t, resp)
			}

			mockmajor.AssertExpectations(t)
		})
	}
}

func TestUpdate(t *testing.T)  {
	mockmajor := RepoMajorMock{}
	servicemajor := RepoMajorConn(&mockmajor)
	data := []testMajorService{
		{
			Name: "Success Update",
			Id: 1,
			ReqMajor: &request.Majors{
				Major: "Computer Science",
			},
			ExpectedErr: nil,
		},
		{
			Name: "Failed Update",
			Id: 2,
			ReqMajor: &request.Majors{},
			ExpectedErr: errors.New("Database Is Refused"),
		},
	}

	for _, v := range data {
		t.Run(v.Name, func(t *testing.T) {
			mockmajor.On("Updates", v.Id, v.ReqMajor).Return(v.ExpectedErr)
			err := servicemajor.UpdatedMajor(v.Id, v.ReqMajor)
			if v.ExpectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			mockmajor.AssertExpectations(t)
		})
	}
}

func TestDelete(t *testing.T)  {
	mockmajor := RepoMajorMock{}
	servicemajor := RepoMajorConn(&mockmajor)
	data := []testMajorService{
		{
			Name: "Success Delete",
			Id: 1,
			ExpectedErr: nil,
		},
		{
			Name: "Failed Delete",
			Id: 2,
			ExpectedErr: errors.New("Database Is Refused"),
		},
	}

	for _, v := range data {
		t.Run(v.Name, func(t *testing.T) {
			mockmajor.On("Deletes", v.Id).Return(v.ExpectedErr)
			err := servicemajor.DeleteMajor(v.Id)
			if v.ExpectedErr !=  nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			mockmajor.AssertExpectations(t)
		})
	}
}
