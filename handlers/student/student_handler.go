package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/service/servicestudent"
	"gorm.io/gorm"
)

type HandlerStudent struct {
	service servicestudent.ServiceStudent
	log helper.Loggers
}

func NewHandlerStudent(s servicestudent.ServiceStudent,l helper.Loggers) HandlerStudent  {
	return  HandlerStudent{service: s, log: l}
}

func (h *HandlerStudent) GetAll(w http.ResponseWriter,r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	if r.Method != helper.Gets {
		w.WriteHeader(helper.MethodNotAllowed)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Only Get Method Is Allowed",
			Errors: "Method Not Allowed",
		})
		return
	}

	page,perpage,sort,search,err:=helper.QueryParam(r,10)
	if err != nil {
		w.WriteHeader(helper.BadRequest)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Invalid Format Query Params",
			Errors: "Bad Request",
		})
		return
	}

	resp, count, err:= h.service.GetAllStudent(search,sort,page,perpage)
	if err != nil {
		h.log.Logfile(err.Error())
		w.WriteHeader(helper.InternalServError)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Something Went Wrong",
			Errors: "Internal Server Error",
		})
		return
	}

	w.WriteHeader(helper.Success)
	json.NewEncoder(w).Encode(helper.Messages{
		Message: "Success",
		Data: resp,
		Pagination: helper.Paginations(page,perpage,int(count)),
	})
}

func (h *HandlerStudent) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	if r.Method != helper.Post {
		w.WriteHeader(helper.MethodNotAllowed)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Only Post Method Allowed",
			Errors: "Method Not Allowed",
		})
		return
	}

	studentreq := new(request.Students)
	if err := json.NewDecoder(r.Body).Decode(&studentreq); err != nil {
		w.WriteHeader(helper.BadRequest)
		h.log.Logfile(err.Error())
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Body Request Is Missing",
			Errors: "Bad Request",
		})
		return
	}

	if err:= helper.ValidateForm(studentreq); err != nil {
		w.WriteHeader(helper.UnprocessbleEntity)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Failed Validation",
			Errors: err.Error(),
		})
		return
	}


	if err := h.service.CreateStudent(studentreq); err != nil {
		if helper.IsDuplicateEntryError(err) {
			w.WriteHeader(helper.Conflict)
			json.NewEncoder(w).Encode(helper.Messages{
				Message: "Nisn Or Phone Is Exist",
				Errors: "Conflict",
			})
			return
		}

		h.log.Logfile(err.Error())
		w.WriteHeader(helper.InternalServError)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Something Went Wrong",
			Errors: "Internal Server Error",
		})
		return
	}

	w.WriteHeader(helper.Success)
	json.NewEncoder(w).Encode(helper.Messages{
		Message: "Success Create New Student",
	})
}

func (h *HandlerStudent) Show(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	if r.Method != helper.Gets {
		w.WriteHeader(helper.MethodNotAllowed)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Only Get Method Is Allowed",
			Errors: "Method Not Allowed",
		})
		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(helper.BadRequest)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Invalid Params Id Format",
			Errors: "Bad Request",
		})
		return
	}

	resp, err := h.service.ShowStudent(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(helper.Notfound)
			json.NewEncoder(w).Encode(helper.Messages{
				Message: fmt.Sprintf("Not Found Data Id: %d", id),
				Errors: "Not Found",
			})
			return
		}

		h.log.Logfile(err.Error())
		w.WriteHeader(helper.InternalServError)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Something Went Wrong",
			Errors: "Internal Server Error",
		})
		return
	}

	w.WriteHeader(helper.Success)
	json.NewEncoder(w).Encode(helper.Messages{
		Message: "Success",
		Data: resp,
	})
}

func (h *HandlerStudent) Update(w http.ResponseWriter,r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	if r.Method != helper.Put {
		w.WriteHeader(helper.MethodNotAllowed)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Only Put Method Is Allowed",
			Errors: "Method Not Allowed",
		})
		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(helper.BadRequest)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Invalid Format Id Params",
			Errors: "Bad Request",
		})
		return
	}

	datareq := new(request.Students)
	if err := json.NewDecoder(r.Body).Decode(&datareq); err != nil {
		w.WriteHeader(helper.BadRequest)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Body Request Is Missing",
			Errors: "Bad Request",
		})

		return
	}

	if err:= h.service.UpdateStudent(id, datareq);err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(helper.Notfound)
			json.NewEncoder(w).Encode(helper.Messages{
				Message: fmt.Sprintf("Not Found Id: %d", id),
				Errors: "Not Found",
			})
			return
		}

		if helper.IsDuplicateEntryError(err) {
			w.WriteHeader(helper.Conflict)
			json.NewEncoder(w).Encode(helper.Messages{
				Message: "Nisn Or Phone Is Exists",
				Errors: "Conflict",
			})
			return
		}

		h.log.Logfile(err.Error())
		w.WriteHeader(helper.InternalServError)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Something Went Wrong",
			Errors: "Internal Server Error",
		})

		return
	}

	w.WriteHeader(helper.Success)
	json.NewEncoder(w).Encode(helper.Messages{
		Message: "Success Update Students",
	})
}

func (h *HandlerStudent) Delete(w http.ResponseWriter,r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
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
			Message: "Invalid Format Id Params",
			Errors: "Bad Request",
		})
		return
	}

	if err := h.service.DeleteStudent(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(helper.Notfound)
			json.NewEncoder(w).Encode(helper.Messages{
				Message: fmt.Sprintf("Not Found Id: %d", id),
				Errors: "Not Found",
			})
			return
		}

		h.log.Logfile(err.Error())
		w.WriteHeader(helper.InternalServError)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Something Went Wrong",
			Errors: "Internal Server Error",
		})

		return
	}


	w.WriteHeader(helper.Success)
	json.NewEncoder(w).Encode(helper.Messages{
		Message: "Success Delete Student",
	})
}