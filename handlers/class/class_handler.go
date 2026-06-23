package class

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

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
	if r.Method != helper.Gets {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.MethodNotAllowed,
			Message: "Only Get Method Is Allowed",
			Errors: "Method Not Allowed",
		})
		return
	}

	page,perpage,sort,search, err :=helper.QueryParam(r, 10)
	if err != nil {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.BadRequest,
			Message: "Invalid Query Params",
			Errors: "Bad Request",
		})
		return
	}

	ctx, cancel := helper.Ctxtimeout(r.Context(), 5*time.Second)
	defer cancel()

	resp,count,err := h.service.AllData(ctx,search,sort,page,perpage)
	if err != nil {
		h.logg.Logfile(err.Error())
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.InternalServError,
			Message: "Something Went Wrong",
			Errors: "Internal Server Error",
		})
		return
	}

	helper.ReturnResponse(w, helper.Messages{
		Http_code: helper.Success,
		Message: "Success",
		Data: resp,
		Pagination: helper.Paginations(page,perpage,int(count)),
	})
}

func (h *HandlerClass) Create(w http.ResponseWriter, r *http.Request)  {
	if r.Method != helper.Post {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.MethodNotAllowed,
			Message: "Only Post Method Is Allowed",
			Errors: "Method Not Allowed",
		})
		return
	}

	classreq := new(request.Class)
	if err := json.NewDecoder(r.Body).Decode(classreq); err != nil {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.BadRequest,
			Message: "Body Request Is Missing",
			Errors: "Bad Request",
		})
		return
	}

	if err:= helper.ValidateForm(classreq); err != nil {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.UnprocessbleEntity,
			Message: "Validation Failed",
			Errors: err.Error(),
		})
		return
	}

	ctx, cancel := helper.Ctxtimeout(r.Context(), 10*time.Second)
	defer cancel()

	if err := h.service.CreateClass(ctx, classreq); err != nil {
		h.logg.Logfile(err.Error())
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.InternalServError,
			Message: "Something Went Wrong",
			Errors: "Internal Server Error",
		})
		return
	}

	helper.ReturnResponse(w, helper.Messages{
		Http_code: helper.Created,
		Message: "Success Create a New Class",
	})
}

func (h *HandlerClass) Show(w http.ResponseWriter, r *http.Request)  {
	if r.Method != helper.Gets {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.MethodNotAllowed,
			Message: "Only Get Method Is Allowed",
			Errors: "Method Not Allowed",
		})
		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.BadRequest,
			Message: "Invalid Params Id Level Format",
			Errors: "Bad Request",
		})
		return
	}

	ctx, cancel := helper.Ctxtimeout(r.Context(), 5*time.Second)
	defer cancel()

	resp, err := h.service.ShowClass(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.ReturnResponse(w, helper.Messages{
				Http_code: helper.Notfound,
				Message: fmt.Sprintf("Not Found Id: %d", id),
				Errors: "Not Found",
			})
			return
		}

		h.logg.Logfile(err.Error())
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.InternalServError,
			Message: "Something Went Wrong",
			Errors: "Internal Server Error",
		})
		return
	}

	helper.ReturnResponse(w, helper.Messages{
		Http_code: helper.Success,
		Message: "Success",
		Data: resp,
	})
}

func (h *HandlerClass) Update(w http.ResponseWriter, r *http.Request)  {
	if r.Method != helper.Put {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.MethodNotAllowed,
			Message: "Only Put Method Not Allowed",
			Errors: "Method Not Allowed",
		})

		return
	}

	classreq := new(request.Class)
	if err := json.NewDecoder(r.Body).Decode(&classreq); err != nil {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.BadRequest,
			Message: "Body Request Is Missing",
			Errors: "Bad Request",
		})
		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.BadRequest,
			Message: "Invalid Params Id Class Format",
			Errors: "Bad Request",
		})
		return
	}

	ctx, cancel := helper.Ctxtimeout(r.Context(), 10*time.Second)
	defer cancel()

	if err := h.service.UpdateClass(ctx, id, classreq); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.ReturnResponse(w, helper.Messages{
				Http_code: helper.Notfound,
				Message: fmt.Sprintf("Not Found id: %d", id),
				Errors: "Not Found",
			})
			return
		}

		h.logg.Logfile(err.Error())
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.InternalServError,
			Message: "Something Went Wrong",
			Errors: "Internal Server Error",
		})
		return
	}

	helper.ReturnResponse(w, helper.Messages{
		Http_code: helper.Success,
		Message: "Success Updated Class",
	})
}

func (h *HandlerClass) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != helper.Delete {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.MethodNotAllowed,
			Message: "Only Delete Method Not Allowed",
			Errors: "Method Not Allowed",
		})

		return
	}

	id, err:= strconv.Atoi(r.PathValue("id"))
	if err != nil {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.BadRequest,
			Message: "Invalid Id Class Format",
			Errors: "Bad Request",
		})
		return
	}

	ctx, cancel := helper.Ctxtimeout(r.Context(), 5*time.Second)
	defer cancel()

	if err := h.service.DeleteClass(ctx, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.ReturnResponse(w, helper.Messages{
				Http_code: helper.Notfound,
				Message: fmt.Sprintf("Not Found Id: %d",id),
				Errors: "Not Found",
			})
			return
		}
		
		h.logg.Logfile(err.Error())
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.InternalServError,
			Message: "Something Went Wrong",
			Errors: "Internal Server Error",
		})
		return
	}

	helper.ReturnResponse(w, helper.Messages{
		Http_code: helper.Success,
		Message: "Success Delete Class",
	})
}