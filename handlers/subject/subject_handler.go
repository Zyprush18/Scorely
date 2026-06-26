package subject

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

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
	if r.Method != helper.Gets {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.MethodNotAllowed,
			Message: "Only Get Method Is Allowed",
			Errors: "Method Not Allowed",
		})
		return
	}

	page,perpage,sort,search,err:= helper.QueryParam(r,10)
	if err != nil {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.BadRequest,
			Message: "Invalid Query Params Format",
			Errors: "Bad Request",
		})
		return
	}

	ctx, cancel := helper.Ctxtimeout(r.Context(), 5*time.Second)
	defer cancel()

	resp,count,err:= h.service.GetAllSubject(ctx,search,sort,page,perpage)
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

func (h *HandlerSubject) Create(w http.ResponseWriter,r *http.Request)  {
	if r.Method != helper.Post {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.MethodNotAllowed,
			Message: "Only Post Method Is Allowed",
			Errors: "Method Not Allowed",
		})
		return
	}

	subjectreq := new(request.Subjects)
	if err := json.NewDecoder(r.Body).Decode(subjectreq); err != nil {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.BadRequest,
			Message: "Request Body Is Missing",
			Errors: "Bad Request",
		})
		return
	}

	if err:=helper.ValidateForm(subjectreq);err != nil {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.UnprocessbleEntity,
			Message: "Failed Validation",
			Errors: err.Error(),
		})
		return
	}

	ctx, cancel := helper.Ctxtimeout(r.Context(), 10*time.Second)
	defer cancel()

	if err := h.service.CreateSubject(ctx, subjectreq);err != nil {
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
		Message: "Success Create Subject",
	})
}

func (h *HandlerSubject) Show(w http.ResponseWriter,r *http.Request) {
	if r.Method != helper.Gets {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.MethodNotAllowed,
			Message: "Only Get Method Is Allowed",
			Errors: "Method Not Allowed",
		})
		return
	}

	id,err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.BadRequest,
			Message: "Invalid Id Params Format",
			Errors: "Bad Request",
		})
		return
	}

	ctx, cancel := helper.Ctxtimeout(r.Context(), 5*time.Second)
	defer cancel()

	resp,err := h.service.ShowSubject(ctx, id)
	if err != nil {
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
			Message: "Sonething Went Wrong",
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

func (h *HandlerSubject) Update(w http.ResponseWriter,r *http.Request) {
	if r.Method != helper.Put {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.MethodNotAllowed,
			Message: "Only Put Method Is Allowed",
			Errors: "Method Not Allowed",
		})
		return
	}

	id,err:= strconv.Atoi(r.PathValue("id"))
	if err != nil {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.BadRequest,
			Message: "Invalid Params Id Format",
			Errors: "Bad Request",
		})
		return
	}

	subjectreq := new(request.Subjects)
	if err := json.NewDecoder(r.Body).Decode(subjectreq); err != nil {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.BadRequest,
			Message: "Request Body Is Missing",
			Errors: "Bad Request",
		})
		return
	}

	ctx, cancel := helper.Ctxtimeout(r.Context(), 10*time.Second)
	defer cancel()

	if err := h.service.UpdateSubject(ctx, id, subjectreq); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.ReturnResponse(w, helper.Messages{
				Http_code: helper.Notfound,
				Message: fmt.Sprintf("Not Found id: %d",id),
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
		Message: "Success Update Subject",
	})
}

func (h *HandlerSubject) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != helper.Delete {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.MethodNotAllowed,
			Message: "Only Delete Method Is Allowed",
			Errors: "Method Not Allowed",
		})
		return
	}

	id,err:=strconv.Atoi(r.PathValue("id"))
	if err != nil {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.BadRequest,
			Message: "Invalid Params Id Format",
			Errors: "Bad Request",
		})
		return
	}

	ctx, cancel := helper.Ctxtimeout(r.Context(), 5*time.Second)
	defer cancel()

	if err := h.service.DeleteSubject(ctx, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.ReturnResponse(w, helper.Messages{
				Http_code: helper.Notfound,
				Message: fmt.Sprintf("Not Found id:%d",id),
				Errors: "Not Found",
			})
			return
		}

		h.logg.Logfile(err.Error())
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.InternalServError,
			Message: "Something Went Wrong",
			Errors: "INternal Server Error",
		})

		return
	}

	helper.ReturnResponse(w, helper.Messages{
		Http_code: helper.Success,
		Message: "Success Delete Subject",
	})
}