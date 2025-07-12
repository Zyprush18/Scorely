package major

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

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
	w.Header().Set("Content-Type", "application/json")
	if r.Method != helper.Gets {
		w.WriteHeader(helper.MethodNotAllowed)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Only Get Method is Allowed",
			Errors:  "Method Not Allowed",
		})
		return
	}

	page, perpage, sort, search, err := helper.QueryParam(r, 10)
	if err != nil {
		w.WriteHeader(helper.BadRequest)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Invalid Format Query Params",
			Errors:  "Bad Request",
		})
		return
	}

	resp, count, err := s.service.GetAllMajor(search, sort, page, perpage)
	if err != nil {
		s.log.Logfile(err.Error())
		w.WriteHeader(helper.InternalServError)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Something Went Wrong",
			Errors:  "Internal Server Error",
		})
		return
	}

	w.WriteHeader(helper.Success)
	json.NewEncoder(w).Encode(helper.Messages{
		Message:    "Success",
		Data:       resp,
		Pagination: helper.Paginations(page, perpage, int(count)),
	})
}

func (s *ServiceMajor) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != helper.Post {
		w.WriteHeader(helper.MethodNotAllowed)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Only Post Method Is Allowed",
			Errors:  "Method Not Allowed",
		})
		return
	}

	majorreq := new(request.Majors)
	if err := json.NewDecoder(r.Body).Decode(&majorreq); err != nil {
		w.WriteHeader(helper.BadRequest)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Request Body Is Missing",
			Errors:  "Bad Request",
		})
		return
	}

	if err := helper.ValidateForm(majorreq); err != nil {
		w.WriteHeader(helper.UnprocessbleEntity)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Validation Error",
			Errors:  err.Error(),
		})
		return
	}

	if err := s.service.CreateMajor(majorreq); err != nil {
		if helper.IsDuplicateEntryError(err) {
			w.WriteHeader(helper.Conflict)
			json.NewEncoder(w).Encode(helper.Messages{
				Message: "Data is Exists",
				Errors:  "conflict",
			})
			return
		}

		s.log.Logfile(err.Error())
		w.WriteHeader(helper.InternalServError)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Something Went Wrong",
			Errors:  "Internal Server Error",
		})
		return
	}

	w.WriteHeader(helper.Created)
	json.NewEncoder(w).Encode(helper.Messages{
		Message: "Success Create a New Major",
	})

}

func (s *ServiceMajor) Show(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != helper.Gets {
		w.WriteHeader(helper.MethodNotAllowed)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Only Get Method Is Allowed",
			Errors:  "Method Not Found",
		})
		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(helper.BadRequest)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Invalid Major Id Format",
			Errors:  "Bad Request",
		})
		return
	}

	data, err := s.service.ShowMajor(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(helper.Notfound)
			json.NewEncoder(w).Encode(helper.Messages{
				Message: fmt.Sprintf("Not Found Major Id: %d", id),
				Errors:  "Not Found",
			})
			return
		}

		s.log.Logfile(err.Error())
		w.WriteHeader(helper.InternalServError)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Something Went Wrong",
			Errors:  "Internal Server Error",
		})
		return
	}

	w.WriteHeader(helper.Success)
	json.NewEncoder(w).Encode(helper.Messages{
		Message: "Success",
		Data:    data,
	})
}

func (s *ServiceMajor) Updated(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != helper.Put {
		w.WriteHeader(helper.MethodNotAllowed)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Only Put Method Is Allowed",
			Errors:  "Method Not Allowed",
		})
		return
	}
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(helper.BadRequest)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Inavalid Major Id Format",
			Errors:  "Bad Request",
		})
		return
	}

	majorreq := &request.Majors{}
	if err:= json.NewDecoder(r.Body).Decode(majorreq);err != nil {
		w.WriteHeader(helper.BadRequest)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Body Request Is Missing",
			Errors: "Bad Request",
		})
		return
	}

	if err := s.service.UpdatedMajor(id, majorreq); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(helper.Notfound)
			json.NewEncoder(w).Encode(helper.Messages{
				Message: fmt.Sprintf("Not Found Major Id: %d", id),
				Errors:  "Not Found",
			})
			return
		}

		if helper.IsDuplicateEntryError(err) {
			w.WriteHeader(helper.Conflict)
			json.NewEncoder(w).Encode(helper.Messages{
				Message: "Data Is Exists",
				Errors:  "Conflict",
			})
			return
		}

		s.log.Logfile(err.Error())
		w.WriteHeader(helper.InternalServError)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Something Went Wrong",
			Errors:  "Internal Server Error",
		})
		return
	}

	w.WriteHeader(helper.Success)
	json.NewEncoder(w).Encode(helper.Messages{
		Message: "Success Update Data",
	})

}


func (s *ServiceMajor) Deleted(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != helper.Delete {
		w.WriteHeader(helper.MethodNotAllowed)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Only Delete Method Is Allowed",
			Errors: "Method Not Allowed",
		})
		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(helper.BadRequest)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Invalid Major Id Format",
			Errors: "Bad Request",
		})
		return
	}


	if err:= s.service.DeleteMajor(id);err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(helper.Notfound)
			json.NewEncoder(w).Encode(helper.Messages{
				Message: fmt.Sprintf("Not Found Major Id: %d", id),
				Errors: "Not Found",
			})
			return
		}

		s.log.Logfile(err.Error())
		w.WriteHeader(helper.InternalServError)
			json.NewEncoder(w).Encode(helper.Messages{
				Message: "Something Went Wrong",
				Errors: "Internal Server Error",
			})
			return
	}

	w.WriteHeader(helper.Success)
	json.NewEncoder(w).Encode(helper.Messages{
		Message: "Success Delete Data",
	})
}