package major

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/service/majorservice"
	"gorm.io/gorm"
)

type ServiceMajor struct {
	service majorservice.MajorService
	log     helper.Loggers
}

func Handlers(s majorservice.MajorService, l helper.Loggers) ServiceMajor {
	return ServiceMajor{service: s, log: l}
}

func (s *ServiceMajor) GetAllData(w http.ResponseWriter, r *http.Request) {
	if r.Method != helper.Gets {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.MethodNotAllowed,
			Message: "Only Get Method is Allowed",
			Errors:  "Method Not Allowed",
		})
		return
	}

	page, perpage, sort, search, err := helper.QueryParam(r, 10)
	if err != nil {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.BadRequest,
			Message: "Invalid Format Query Params",
			Errors:  "Bad Request",
		})
		return
	}


	ctx, cancel := helper.Ctxtimeout(r.Context(), 5 * time.Second)
	defer cancel()

	resp, count, err := s.service.GetAllMajor(ctx,search, sort, page, perpage)
	if err != nil {
		s.log.Logfile(err.Error())
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.InternalServError,
			Message: "Something Went Wrong",
			Errors:  "Internal Server Error",
		})
		return
	}

	helper.ReturnResponse(w, helper.Messages{
		Http_code: helper.Success,
		Message:    "Success",
		Data:       resp,
		Pagination: helper.Paginations(page, perpage, int(count)),
	})
}

func (s *ServiceMajor) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != helper.Post {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.MethodNotAllowed,
			Message: "Only Post Method Is Allowed",
			Errors:  "Method Not Allowed",
		})
		return
	}

	majorreq := new(request.Majors)
	if err := json.NewDecoder(r.Body).Decode(&majorreq); err != nil {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.BadRequest,
			Message: "Request Body Is Missing",
			Errors:  "Bad Request",
		})
		return
	}

	if err := helper.ValidateForm(majorreq); err != nil {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.UnprocessbleEntity,
			Message: "Validation Error",
			Errors:  err.Error(),
		})
		return
	}

	ctx, cancel := helper.Ctxtimeout(r.Context(), 10 * time.Second)
	defer cancel()

	if err := s.service.CreateMajor(ctx,majorreq); err != nil {
		if helper.IsDuplicateEntryError(err) {
			helper.ReturnResponse(w, helper.Messages{
				Http_code: helper.Conflict,
				Message: "Data is Exists",
				Errors:  "conflict",
			})
			return
		}

		s.log.Logfile(err.Error())
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.InternalServError,
			Message: "Something Went Wrong",
			Errors:  "Internal Server Error",
		})
		return
	}

	helper.ReturnResponse(w, helper.Messages{
		Http_code: helper.Created,
		Message: "Success Create a New Major",
	})

}

func (s *ServiceMajor) Show(w http.ResponseWriter, r *http.Request) {
	if r.Method != helper.Gets {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.MethodNotAllowed,
			Message: "Only Get Method Is Allowed",
			Errors:  "Method Not Found",
		})
		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.BadRequest,
			Message: "Invalid Major Id Format",
			Errors:  "Bad Request",
		})
		return
	}

	ctx, cancel := helper.Ctxtimeout(r.Context(), 5 * time.Second)
	defer cancel()

	data, err := s.service.ShowMajor(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.ReturnResponse(w, helper.Messages{
				Http_code: helper.Notfound,
				Message: fmt.Sprintf("Not Found Major Id: %d", id),
				Errors:  "Not Found",
			})
			return
		}

		s.log.Logfile(err.Error())
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.InternalServError,
			Message: "Something Went Wrong",
			Errors:  "Internal Server Error",
		})
		return
	}

	helper.ReturnResponse(w, helper.Messages{
		Http_code: helper.Success,
		Message: "Success",
		Data:    data,
	})
}

func (s *ServiceMajor) Updated(w http.ResponseWriter, r *http.Request) {
	if r.Method != helper.Put {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.MethodNotAllowed,
			Message: "Only Put Method Is Allowed",
			Errors:  "Method Not Allowed",
		})
		return
	}
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.BadRequest,
			Message: "Inavalid Major Id Format",
			Errors:  "Bad Request",
		})
		return
	}

	majorreq := &request.Majors{}
	if err:= json.NewDecoder(r.Body).Decode(majorreq);err != nil {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.BadRequest,
			Message: "Body Request Is Missing",
			Errors: "Bad Request",
		})
		return
	}

	ctx, cancel := helper.Ctxtimeout(r.Context(), 10 * time.Second)
	defer cancel()

	if err := s.service.UpdatedMajor(ctx,id, majorreq); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.ReturnResponse(w, helper.Messages{
				Http_code: helper.Notfound,
				Message: fmt.Sprintf("Not Found Major Id: %d", id),
				Errors:  "Not Found",
			})
			return
		}

		if helper.IsDuplicateEntryError(err) {
			helper.ReturnResponse(w, helper.Messages{
				Http_code: helper.Conflict,
				Message: "Data Is Exists",
				Errors:  "Conflict",
			})
			return
		}

		s.log.Logfile(err.Error())
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.InternalServError,
			Message: "Something Went Wrong",
			Errors:  "Internal Server Error",
		})
		return
	}

	helper.ReturnResponse(w, helper.Messages{
		Http_code: helper.Success,
		Message: "Success Update Data",
	})

}


func (s *ServiceMajor) Deleted(w http.ResponseWriter, r *http.Request) {
	if r.Method != helper.Delete {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.MethodNotAllowed,
			Message: "Only Delete Method Is Allowed",
			Errors: "Method Not Allowed",
		})
		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.BadRequest,
			Message: "Invalid Major Id Format",
			Errors: "Bad Request",
		})
		return
	}

	ctx, cancel := helper.Ctxtimeout(r.Context(), 5 * time.Second)
	defer cancel()

	if err:= s.service.DeleteMajor(ctx, id);err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.ReturnResponse(w, helper.Messages{
				Http_code: helper.Notfound,
				Message: fmt.Sprintf("Not Found Major Id: %d", id),
				Errors: "Not Found",
			})
			return
		}

		s.log.Logfile(err.Error())
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.InternalServError,
				Message: "Something Went Wrong",
				Errors: "Internal Server Error",
			})
			return
	}

	helper.ReturnResponse(w, helper.Messages{
		Http_code: helper.Success,
		Message: "Success Delete Data",
	})
}