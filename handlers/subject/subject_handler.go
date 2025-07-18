package subject

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/service/subjectservice"
	"gorm.io/gorm"
)

type HandlerSubject struct {
	service subjectservice.ServiceSubject
	logg helper.Loggers
}

func ConnectService(s subjectservice.ServiceSubject,l helper.Loggers) HandlerSubject  {
	return HandlerSubject{service: s,logg: l}
}

func (h *HandlerSubject) GetAll(w http.ResponseWriter,r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	if r.Method != helper.Gets {
		w.WriteHeader(helper.MethodNotAllowed)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Only Get Method Is Allowed",
			Errors: "Method Not Allowed",
		})
		return
	}

	page,perpage,sort,search,err:= helper.QueryParam(r,10)
	if err != nil {
		w.WriteHeader(helper.BadRequest)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Invalid Query Params Format",
			Errors: "Bad Request",
		})
		return
	}

	resp,count,err:= h.service.GetAllSubject(search,sort,page,perpage)
	if err != nil {
		h.logg.Logfile(err.Error())
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

func (h *HandlerSubject) Create(w http.ResponseWriter,r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	if r.Method != helper.Post {
		w.WriteHeader(helper.MethodNotAllowed)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Only Post Method Is Allowed",
			Errors: "Method Not Allowed",
		})
		return
	}

	subjectreq := new(request.Subjects)
	if err := json.NewDecoder(r.Body).Decode(subjectreq); err != nil {
		w.WriteHeader(helper.BadRequest)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Request Body Is Missing",
			Errors: "Bad Request",
		})
		return
	}

	if err:=helper.ValidateForm(subjectreq);err != nil {
		w.WriteHeader(helper.UnprocessbleEntity)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Failed Validation",
			Errors: err.Error(),
		})
		return
	}

	if err := h.service.CreateSubject(subjectreq);err != nil {
		h.logg.Logfile(err.Error())
		w.WriteHeader(helper.InternalServError)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Something Went Wrong",
			Errors: "Internal Server Error",
		})
		return
	}

	w.WriteHeader(helper.Created)
	json.NewEncoder(w).Encode(helper.Messages{
		Message: "Success Create Subject",
	})
}

func (h *HandlerSubject) Show(w http.ResponseWriter,r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	if r.Method != helper.Gets {
		w.WriteHeader(helper.MethodNotAllowed)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Only Get Method Is Allowed",
			Errors: "Method Not Allowed",
		})
		return
	}

	id,err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(helper.BadRequest)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Invalid Id Params Format",
			Errors: "Bad Request",
		})
		return
	}

	resp,err := h.service.ShowSubject(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(helper.Notfound)
			json.NewEncoder(w).Encode(helper.Messages{
				Message: fmt.Sprintf("Not Found id: %d", id),
				Errors: "Not Found",
			})
			return
		}

		h.logg.Logfile(err.Error())
		w.WriteHeader(helper.InternalServError)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Sonething Went Wrong",
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

func (h *HandlerSubject) Update(w http.ResponseWriter,r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	if r.Method != helper.Put {
		w.WriteHeader(helper.MethodNotAllowed)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Only Put Method Is Allowed",
			Errors: "Method Not Allowed",
		})
		return
	}

	id,err:= strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(helper.BadRequest)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Invalid Params Id Format",
			Errors: "Bad Request",
		})
		return
	}

	subjectreq := new(request.Subjects)
	if err := json.NewDecoder(r.Body).Decode(subjectreq); err != nil {
		w.WriteHeader(helper.BadRequest)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Request Body Is Missing",
			Errors: "Bad Request",
		})
		return
	}

	if err := h.service.UpdateSubject(id, subjectreq); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(helper.Notfound)
			json.NewEncoder(w).Encode(helper.Messages{
				Message: fmt.Sprintf("Not Found id: %d",id),
				Errors: "Not Found",
			})
			return
		}

		h.logg.Logfile(err.Error())
		w.WriteHeader(helper.InternalServError)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Something Went Wrong",
			Errors: "Internal Server Error",
		})
		return
	}

	w.WriteHeader(helper.Success)
	json.NewEncoder(w).Encode(helper.Messages{
		Message: "Success Update Subject",
	})
}

func (h *HandlerSubject) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	if r.Method != helper.Delete {
		w.WriteHeader(helper.MethodNotAllowed)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Only Delete Method Is Allowed",
			Errors: "Method Not Allowed",
		})
		return
	}

	id,err:=strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(helper.BadRequest)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Invalid Params Id Format",
			Errors: "Bad Request",
		})
		return
	}

	if err := h.service.DeleteSubject(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(helper.Notfound)
			json.NewEncoder(w).Encode(helper.Messages{
				Message: fmt.Sprintf("Not Found id:%d",id),
				Errors: "Not Found",
			})
			return
		}

		h.logg.Logfile(err.Error())
		w.WriteHeader(helper.InternalServError)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Something Went Wrong",
			Errors: "INternal Server Error",
		})

		return
	}

	w.WriteHeader(helper.Success)
	json.NewEncoder(w).Encode(helper.Messages{
		Message: "Success Delete Subject",
	})
}