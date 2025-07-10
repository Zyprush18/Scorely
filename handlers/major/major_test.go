package major

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/models/response"
	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

type testHandlerMajor struct {
	Name              string
	DataRespAll       []response.Majors
	RequestMajor      *request.Majors
	DataResp          *response.Majors
	Method, Target    string
	Code, CountData   int
	Search, Sort      string
	Page, Perpage, Id int
	Mux               *http.ServeMux
	UseMock           bool
	ExpectedErr       error
}

func TestGetAll(t *testing.T) {
	mockmajor := ServiceMajorMock{}
	mocklogmajor := LoggerMock{}
	majorhandler := Handlers(&mockmajor, mocklogmajor)

	data := []testHandlerMajor{
		{
			Name:        "Method Not Allowed",
			DataRespAll: nil,
			Method:      helper.Post,
			Target:      "/major",
			Code:        helper.MethodNotAllowed,
			Mux:         http.NewServeMux(),
			UseMock:     false,
			ExpectedErr: nil,
		},
		{
			Name:        "Invalid Format Query Params",
			DataRespAll: nil,
			Method:      helper.Gets,
			Target:      "/major?page=abc",
			Code:        helper.BadRequest,
			Mux:         http.NewServeMux(),
			UseMock:     false,
			ExpectedErr: nil,
		},
		{
			Name:        "Failed Get All Data",
			DataRespAll: nil,
			Method:      helper.Gets,
			Target:      "/major",
			Code:        helper.InternalServError,
			Search:      "",
			Sort:        "asc",
			Page:        1,
			Perpage:     10,
			CountData:   0,
			Mux:         http.NewServeMux(),
			UseMock:     true,
			ExpectedErr: errors.New("Database Is Refused"),
		},
		{
			Name: "Success Get All Data",
			DataRespAll: []response.Majors{
				{
					IdMajor:           1,
					Major:             "Informatics Engineering",
					MajorAbbreviation: "IE",
					Models: helper.Models{
						CreatedAt: time.Now(),
						UpdatedAt: time.Now(),
					},
				},
				{
					IdMajor:           2,
					Major:             "System Information",
					MajorAbbreviation: "SI",
					Models: helper.Models{
						CreatedAt: time.Now(),
						UpdatedAt: time.Now(),
					},
				},
			},
			Method:      helper.Gets,
			Target:      "/major",
			Code:        helper.InternalServError,
			Search:      "",
			Sort:        "asc",
			Page:        1,
			Perpage:     10,
			CountData:   0,
			Mux:         http.NewServeMux(),
			UseMock:     true,
			ExpectedErr: errors.New("Database Is Refused"),
		},
	}

	for _, v := range data {
		t.Run(v.Name, func(t *testing.T) {
			req := httptest.NewRequest(v.Method, v.Target, nil)
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			if v.UseMock {
				mockmajor.On("GetAllMajor", v.Search, v.Sort, v.Page, v.Perpage).Return(v.DataRespAll, v.CountData, v.ExpectedErr)
			}

			v.Mux.HandleFunc("/major", majorhandler.GetAllData)
			v.Mux.ServeHTTP(w, req)

			assert.Equal(t, v.Code, w.Code)
			mockmajor.AssertExpectations(t)
		})
	}
}

func TestCreate(t *testing.T) {
	mockmajor := ServiceMajorMock{}
	mocklogger := LoggerMock{}
	handlermajor := Handlers(&mockmajor, mocklogger)

	data := []testHandlerMajor{
		{
			Name:         "Method Not Allowed",
			RequestMajor: nil,
			Method:       helper.Gets,
			Target:       "/major/add",
			Code:         helper.MethodNotAllowed,
			Mux:          http.NewServeMux(),
			UseMock:      false,
		},
		{
			Name:         "Body Request Is Missing",
			RequestMajor: nil,
			Method:       helper.Post,
			Target:       "/major/add",
			Code:         helper.BadRequest,
			Mux:          http.NewServeMux(),
			UseMock:      false,
		},
		{
			Name:         "Validation Error",
			RequestMajor: &request.Majors{},
			Method:       helper.Post,
			Target:       "/major/add",
			Code:         helper.UnprocessbleEntity,
			Mux:          http.NewServeMux(),
			UseMock:      false,
		},
		{
			Name: "Failed: Data Is Exists",
			RequestMajor: &request.Majors{
				Major:             "System Information",
				MajorAbbreviation: "SI",
			},
			Method:  helper.Post,
			Target:  "/major/add",
			Code:    helper.Conflict,
			Mux:     http.NewServeMux(),
			UseMock: true,
			ExpectedErr: &mysql.MySQLError{
				Message: "Duplicate Entry",
				Number:  1062,
			},
		},
		{
			Name: "Failed: Server Error",
			RequestMajor: &request.Majors{
				Major:             "Informatics Engineering",
				MajorAbbreviation: "IE",
			},
			Method:      helper.Post,
			Target:      "/major/add",
			Code:        helper.InternalServError,
			Mux:         http.NewServeMux(),
			UseMock:     true,
			ExpectedErr: errors.New("Database Is Refused"),
		},
		{
			Name: "Success",
			RequestMajor: &request.Majors{
				Major:             "System Engineering",
				MajorAbbreviation: "SE",
			},
			Method:      helper.Post,
			Target:      "/major/add",
			Code:        helper.Created,
			Mux:         http.NewServeMux(),
			UseMock:     true,
			ExpectedErr: nil,
		},
	}

	for _, v := range data {
		t.Run(v.Name, func(t *testing.T) {
			var majorreq []byte
			if v.RequestMajor != nil {
				var err error
				majorreq, err = json.Marshal(v.RequestMajor)
				assert.NoError(t, err)
			}
			req := httptest.NewRequest(v.Method, v.Target, bytes.NewReader(majorreq))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			if v.UseMock {
				mockmajor.On("CreateMajor", v.RequestMajor).Return(v.ExpectedErr)
			}
			v.Mux.HandleFunc("/major/add", handlermajor.Create)
			v.Mux.ServeHTTP(w, req)

			assert.Equal(t, v.Code, w.Code)
			mockmajor.AssertExpectations(t)
		})
	}
}

func TestShow(t *testing.T) {
	mockmajor := ServiceMajorMock{}
	mocklogger := LoggerMock{}
	handlermajor := Handlers(&mockmajor, mocklogger)

	data := []testHandlerMajor{
		{
			Name:    "Method Not Allowed",
			Method:  helper.Post,
			Target:  "/major/1",
			Code:    helper.MethodNotAllowed,
			Mux:     http.NewServeMux(),
			UseMock: false,
		},
		{
			Name:    "Invalid Major Id Format",
			Method:  helper.Gets,
			Target:  "/major/abc",
			Code:    helper.BadRequest,
			Mux:     http.NewServeMux(),
			UseMock: false,
		},
		{
			Name:    "Invalid Major Id Format",
			Method:  helper.Gets,
			Target:  "/major/abc",
			Code:    helper.BadRequest,
			Mux:     http.NewServeMux(),
			UseMock: false,
		},
		{
			Name:    "Failed: Not Fund Data",
			DataResp: &response.Majors{
				IdMajor: 1,
				Major: "System Information",
				MajorAbbreviation: "SI",
				Models: helper.Models{
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
			},
			Method:  helper.Gets,
			Id: 1,
			Target:  "/major/1",
			Code:    helper.Notfound,
			Mux:     http.NewServeMux(),
			UseMock: true,
			ExpectedErr: gorm.ErrRecordNotFound,
		},
		{
			Name:    "Failed: Server Error",
			DataResp: &response.Majors{
				IdMajor: 2,
				Major: "Informatics Engineering",
				MajorAbbreviation: "IE",
				Models: helper.Models{
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
			},
			Method:  helper.Gets,
			Id: 2,
			Target:  "/major/2",
			Code:    helper.InternalServError,
			Mux:     http.NewServeMux(),
			UseMock: true,
			ExpectedErr: errors.New("Databse Is Refused"),
		},
		{
			Name:    "Success",
			DataResp: &response.Majors{
				IdMajor: 3,
				Major: "System Engineering",
				MajorAbbreviation: "SE",
				Models: helper.Models{
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
			},
			Method:  helper.Gets,
			Id: 3,
			Target:  "/major/3",
			Code:    helper.Success,
			Mux:     http.NewServeMux(),
			UseMock: true,
			ExpectedErr: nil,
		},
	}

	for _, v := range data {
		t.Run(v.Name, func(t *testing.T) {
			req := httptest.NewRequest(v.Method, v.Target, nil)
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			if v.UseMock {
				mockmajor.On("ShowMajor", v.Id).Return(v.DataResp, v.ExpectedErr)
			}

			v.Mux.HandleFunc("/major/{id}", handlermajor.Show)
			v.Mux.ServeHTTP(w, req)

			assert.Equal(t, v.Code, w.Code)
			mockmajor.AssertExpectations(t)
		})
	}
}

func TestUpdate(t *testing.T)  {
	mockmajor := ServiceMajorMock{}
	mocklogger := LoggerMock{}
	handlermajor := Handlers(&mockmajor,mocklogger)

	data := []testHandlerMajor{
		{
			Name: "Method Not Found",
			RequestMajor: nil,
			Method: helper.Gets,
			Target: "/major/1/update",
			Code: helper.MethodNotAllowed,
			Mux: http.NewServeMux(),
			UseMock: false,
		},
		{
			Name: "Invalid major Id Format",
			RequestMajor: nil,
			Method: helper.Put,
			Target: "/major/abc/update",
			Code: helper.BadRequest,
			Mux: http.NewServeMux(),
			UseMock: false,
		},
		{
			Name: "Body Request Is Missing",
			RequestMajor: nil,
			Method: helper.Put,
			Target: "/major/2/update",
			Code: helper.BadRequest,
			Mux: http.NewServeMux(),
			UseMock: false,
		},
		{
			Name: "Failed: Not Found Data",
			RequestMajor: &request.Majors{
				Major: "System Information",
				MajorAbbreviation: "SI",
			},
			Id: 2,
			Method: helper.Put,
			Target: "/major/2/update",
			Code: helper.Notfound,
			Mux: http.NewServeMux(),
			UseMock: true,
			ExpectedErr: gorm.ErrRecordNotFound,
		},
		{
			Name: "Failed: Duplicate",
			RequestMajor: &request.Majors{
				Major: "System Engineering",
				MajorAbbreviation: "SE",
			},
			Id: 3,
			Method: helper.Put,
			Target: "/major/3/update",
			Code: helper.Conflict,
			Mux: http.NewServeMux(),
			UseMock: true,
			ExpectedErr: &mysql.MySQLError{
				Message: "Duplicate Entry",
				Number: 1062,
			},
		},
		{
			Name: "Failed: Internal Server Error",
			RequestMajor: &request.Majors{
				Major: "Information Engineering",
				MajorAbbreviation: "IE",
			},
			Id: 4,
			Method: helper.Put,
			Target: "/major/4/update",
			Code: helper.InternalServError,
			Mux: http.NewServeMux(),
			UseMock: true,
			ExpectedErr: errors.New("Database Is Refused"),
		},
		{
			Name: "Failed: Internal Server Error",
			RequestMajor: &request.Majors{
				Major: "Informatics Engineering",
				MajorAbbreviation: "IE",
			},
			Id: 1,
			Method: helper.Put,
			Target: "/major/1/update",
			Code: helper.Success,
			Mux: http.NewServeMux(),
			UseMock: true,
			ExpectedErr: nil,
		},
	}

	for _, v := range data {
		t.Run(v.Name,func(t *testing.T) {
			var majorreq []byte
			if v.RequestMajor != nil {
				var err error
				majorreq, err = json.Marshal(v.RequestMajor)
				assert.NoError(t, err)
			}
			req := httptest.NewRequest(v.Method,v.Target, bytes.NewReader(majorreq))
			req.Header.Set("Content-Type", "application/json")
			w:= httptest.NewRecorder()

			if v.UseMock {
				mockmajor.On("UpdatedMajor", v.Id,v.RequestMajor).Return(v.ExpectedErr)
			}

			v.Mux.HandleFunc("/major/{id}/update", handlermajor.Updated)
			v.Mux.ServeHTTP(w, req)

			assert.Equal(t, v.Code,w.Code)
			mockmajor.AssertExpectations(t)
		})
	}
}

func TestDelte(t *testing.T)  {
	mockmajor := ServiceMajorMock{}
	mocklogger := LoggerMock{}
	handlermajor := Handlers(&mockmajor,mocklogger)

	data := []testHandlerMajor{
		{
			Name: "Method Not Allowed",
			Method: helper.Gets,
			Target: "/major/1/delete",
			Code: helper.MethodNotAllowed,
			Mux: http.NewServeMux(),
			UseMock: false,
		},
		{
			Name: "Invalid Major Id Format",
			Method: helper.Delete,
			Target: "/major/abc/delete",
			Code: helper.BadRequest,
			Mux: http.NewServeMux(),
			UseMock: false,
		},
		{
			Name: "Failed: Not Found",
			Method: helper.Delete,
			Id: 3,
			Target: "/major/3/delete",
			Code: helper.Notfound,
			Mux: http.NewServeMux(),
			UseMock: true,
			ExpectedErr: gorm.ErrRecordNotFound,
		},
		{
			Name: "Failed: Server Error",
			Method: helper.Delete,
			Id: 2,
			Target: "/major/2/delete",
			Code: helper.InternalServError,
			Mux: http.NewServeMux(),
			UseMock: true,
			ExpectedErr: errors.New("Database Is Refused"),
		},
		{
			Name: "Success",
			Method: helper.Delete,
			Id: 1,
			Target: "/major/1/delete",
			Code: helper.Success,
			Mux: http.NewServeMux(),
			UseMock: true,
			ExpectedErr: nil,
		},
	}

	for _, v := range data {
		t.Run(v.Name,func(t *testing.T) {
			req := httptest.NewRequest(v.Method,v.Target,nil)
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			if v.UseMock {
				mockmajor.On("DeleteMajor", v.Id).Return(v.ExpectedErr)
			}

			v.Mux.HandleFunc("/major/{id}/delete", handlermajor.Deleted)
			v.Mux.ServeHTTP(w, req)

			assert.Equal(t, v.Code,w.Code)
			mockmajor.AssertExpectations(t)
		})
	}
}