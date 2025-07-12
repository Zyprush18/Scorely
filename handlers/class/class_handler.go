package class

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/service/classservice"
	"gorm.io/gorm"
)

type HandlerClass struct {
	service classservice.ServiceClass
	logg 	helper.Loggers
}

func NewHandlerClass(s classservice.ServiceClass, l helper.Loggers) HandlerClass  {
	return HandlerClass{service: s,logg: l}
}

func (h *HandlerClass) GetAll(w http.ResponseWriter,r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	if r.Method != helper.Gets {
		w.WriteHeader(helper.MethodNotAllowed)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Only Get Method Is Allowed",
			Errors: "Method Not Allowed",
		})
		return
	}

	page,perpage,sort,search, err :=helper.QueryParam(r, 10)
	if err != nil {
		w.WriteHeader(helper.BadRequest)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Invalid Query Params",
			Errors: "Bad Request",
		})
		return
	}

	resp,count,err := h.service.AllData(search,sort,page,perpage)
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

func (h *HandlerClass) Create(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != helper.Post {
		w.WriteHeader(helper.MethodNotAllowed)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Only Post Method Is Allowed",
			Errors: "Method Not Allowed",
		})
		return
	}

	classreq := new(request.Class)
	if err := json.NewDecoder(r.Body).Decode(classreq); err != nil {
		w.WriteHeader(helper.BadRequest)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Body Request Is Missing",
			Errors: "Bad Request",
		})
		return
	}

	if err:= helper.ValidateForm(classreq); err != nil {
		w.WriteHeader(helper.UnprocessbleEntity)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Validation Failed",
			Errors: err.Error(),
		})
		return
	}

	if err := h.service.CreateClass(classreq); err != nil {
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
		Message: "Success Create a New Class",
	})
}

func (h *HandlerClass) Show(w http.ResponseWriter, r *http.Request)  {
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
			Message: "Invalid Params Id Level Format",
			Errors: "Bad Request",
		})
		return
	}

	resp, err := h.service.ShowClass(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(helper.Notfound)
			json.NewEncoder(w).Encode(helper.Messages{
				Message: fmt.Sprintf("Not Found Id: %d", id),
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
		Message: "Success",
		Data: resp,
	})
}

func (h *HandlerClass) Update(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	if r.Method != helper.Put {
		w.WriteHeader(helper.MethodNotAllowed)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Only Put Method Not Allowed",
			Errors: "Method Not Allowed",
		})

		return
	}

	classreq := new(request.Class)
	if err := json.NewDecoder(r.Body).Decode(&classreq); err != nil {
		w.WriteHeader(helper.BadRequest)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Body Request Is Missing",
			Errors: "Bad Request",
		})
		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(helper.BadRequest)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Invalid Params Id Class Format",
			Errors: "Bad Request",
		})
		return
	}

	if err := h.service.UpdateClass(id,classreq); err != nil {
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
			Message: "Something Went Wrong",
			Errors: "Internal Server Error",
		})
		return
	}

	w.WriteHeader(helper.Success)
	json.NewEncoder(w).Encode(helper.Messages{
		Message: "Success Updated Class",
	})
}

func (h *HandlerClass) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	if r.Method != helper.Delete {
		w.WriteHeader(helper.MethodNotAllowed)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Only Delete Method Not Allowed",
			Errors: "Method Not Allowed",
		})

		return
	}

	id, err:= strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(helper.BadRequest)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Invalid Id Class Format",
			Errors: "Bad Request",
		})
		return
	}

	if err := h.service.DeleteClass(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(helper.Notfound)
			json.NewEncoder(w).Encode(helper.Messages{
				Message: fmt.Sprintf("Not Found Id: %d",id),
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
		Message: "Success Delete Class",
	})
}