package teacher

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/service/serviceteacher"
	"gorm.io/gorm"
)

type HandlerTeacher struct {
	serviice serviceteacher.ServiceTeacher
	logg helper.Loggers
}


func ConnectService(s serviceteacher.ServiceTeacher, l helper.Loggers) HandlerTeacher  {
	return  HandlerTeacher{serviice: s,logg: l}
}

func (h *HandlerTeacher) GetAll(w http.ResponseWriter,r *http.Request) {
	if r.Method != helper.Gets {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.MethodNotAllowed,
			Message: "Only Get Method Is Allowed",
			Errors: "Method Not Found",
		})
		return
	}

	page,perpage,sort,search,err:=helper.QueryParam(r, 10)
	if err != nil {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.BadRequest,
			Message: "Invalid Format Query Params",
			Errors: "Bad Request",
		})
		return
	}

	ctx, cancel := helper.Ctxtimeout(r.Context(), 5*time.Second)
	defer cancel()

	resp,count, err := h.serviice.GetAllTeacher(ctx,search,sort,page,perpage)
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
		Pagination: helper.Paginations(page,perpage, int(count)),
	})

}


func (h *HandlerTeacher) Create(w http.ResponseWriter, r *http.Request)  {
	if r.Method != helper.Post {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.MethodNotAllowed,
			Message: "Only Post Method Is Allowed",
			Errors: "Method Not Allowed",
		})
		return
	}


	datareq := new(request.Teachers)
	if err := json.NewDecoder(r.Body).Decode(datareq);err != nil {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.BadRequest,
			Message: "Body request Is Missing",
			Errors: "Bad Request",
		})
		return
	}

	if err := helper.ValidateForm(datareq); err != nil {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.UnprocessbleEntity,
			Message: "Failed Validation",
			Errors: err.Error(),
		})
		return
	}

	ctx, cancel := helper.Ctxtimeout(r.Context(), 10*time.Second)
	defer cancel()

	if err := h.serviice.CreateTeacher(ctx, datareq);err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.ReturnResponse(w, helper.Messages{
				Http_code: helper.Notfound,
				Message: "Subject Not Found",
				Errors: "Not Found",
			})
			return
		}

		if helper.IsDuplicateEntryError(err) {
			helper.ReturnResponse(w, helper.Messages{
				Http_code: helper.Conflict,
				Message: "Nip Or Phone Is Already Exists",
				Errors: "Conflict",
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
		Http_code: helper.Created,
		Message: "Success Create new Teacher",
	})
}


func (h *HandlerTeacher) Show(w http.ResponseWriter,r *http.Request) {
	if r.Method != helper.Gets {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.MethodNotAllowed,
			Message: "Only Get Method Is Allowed",
			Errors: "Method Not Allowed",
		})
		return
	}


	id,err:= strconv.Atoi(r.PathValue("id"))
	if err != nil {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.BadRequest,
			Message: "Invalid Parms Id Format",
			Errors: "Bad Request",
		})
		return
	}

	ctx, cancel := helper.Ctxtimeout(r.Context(), 5*time.Second)
	defer cancel()

	resp,err := h.serviice.ShowTeacher(ctx, id)
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
			Message: "Somethin Went Wrong",
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


func (h *HandlerTeacher) Update(w http.ResponseWriter,r *http.Request) {
	if r.Method != helper.Put {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.MethodNotAllowed,
			Message: "Only Put Method Is Allowed",
			Errors: "Method Not Allowed",
		})
		return
	}

	datareq := new(request.Teachers)
	if err := json.NewDecoder(r.Body).Decode(&datareq);err != nil {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.BadRequest,
			Message: "Request Body Is Missing",
			Errors: "Bad Request",
		})
		return
	}

	id,err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.BadRequest,
			Message: "Invalid Params Id Format",
			Errors: "Bad Request",
		})

		return
	}

	ctx, cancel := helper.Ctxtimeout(r.Context(), 10*time.Second)
	defer cancel()

	if err:= h.serviice.UpdateTeacher(ctx, id, datareq); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.ReturnResponse(w, helper.Messages{
				Http_code: helper.Notfound,
				Message: fmt.Sprintf("Not Found Id:%d", id),
				Errors: "Not Found",
			})
			return
		}

		if helper.IsDuplicateEntryError(err) {
			helper.ReturnResponse(w, helper.Messages{
				Http_code: helper.Conflict,
				Message: "Nip Or Phone Is Already Exists",
				Errors: "Conflict",
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
		Message: "SUccess Update Teacher",
	})
}

func (h *HandlerTeacher) Delete(w http.ResponseWriter, r *http.Request) {
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
			Message: "Invalid Params Id Format",
			Errors: "Bad Request",
		})
		return
	}

	ctx, cancel := helper.Ctxtimeout(r.Context(), 5*time.Second)
	defer cancel()

	if err := h.serviice.DeleteTeacher(ctx, id); err!= nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.ReturnResponse(w, helper.Messages{
				Http_code: helper.Notfound,
				Message: fmt.Sprintf("Not Found id:%d", id),
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
		Message: "Success Delete Teacher",
	})
}