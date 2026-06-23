package level

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/service/servicelevel"
	"gorm.io/gorm"
)

type HandlerLevel struct {
	service servicelevel.LevelService
	logs	helper.Loggers
}

func ConnectService(s servicelevel.LevelService, l helper.Loggers) HandlerLevel  {
	return HandlerLevel{service: s,logs: l}
}

func (h *HandlerLevel) GetAll(w http.ResponseWriter, r *http.Request)  {
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

	resp, count, err := h.service.GetAllLevel(ctx,search,sort,page,perpage)
	if err != nil {
		h.logs.Logfile(err.Error())
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


func (h *HandlerLevel) Create (w http.ResponseWriter, r *http.Request)  {
	if r.Method != helper.Post {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.MethodNotAllowed,
			Message: "Only Post Method Is Allowed",
			Errors: "Method Not Allowed",
		})
		return
	}

	levelreq := new(request.Levels)
	if err:=json.NewDecoder(r.Body).Decode(levelreq); err!= nil {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.BadRequest,
			Message: "Body Request Is Missing",
			Errors: "Bad Request",
		})

		return
	}

	// validate
	if err:=helper.ValidateForm(levelreq);err != nil {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.UnprocessbleEntity,
			Message: "Validation Error",
			Errors: err.Error(),
		})

		return
	}

	ctx, cancel := helper.Ctxtimeout(r.Context(), 10*time.Second)
	defer cancel()

	if err:=h.service.CreateLevel(ctx, levelreq);err != nil {
		if helper.IsDuplicateEntryError(err) {
			helper.ReturnResponse(w, helper.Messages{
				Http_code: helper.Conflict,
				Message: "Data Is Exists",
				Errors: "Conflict",
			})
			return
		}

		h.logs.Logfile(err.Error())
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.InternalServError,
			Message: "Something Went Wrong",
			Errors: "Internal Server Error",
		})
		return
	}

	helper.ReturnResponse(w, helper.Messages{
		Http_code: helper.Created,
		Message: "Success Created a New Level",
	})
}

func (h *HandlerLevel) Show (w http.ResponseWriter, r *http.Request)  {
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
			Message: "Invalid Level Id Format",
			Errors: "Bad Request",
		})

		return
	} 


	ctx, cancel := helper.Ctxtimeout(r.Context(), 5*time.Second)
	defer cancel()

	resp, err := h.service.ShowLevel(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.ReturnResponse(w, helper.Messages{
				Http_code: helper.Notfound,
				Message: fmt.Sprintf("Not Found Data Id: %d", id),
				Errors: "Not Found",
			})
			return
		}

		h.logs.Logfile(err.Error())
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

func (h *HandlerLevel) Update(w http.ResponseWriter, r *http.Request)  {
	if r.Method != helper.Put {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.MethodNotAllowed,
			Message: "Only Put Method Is Allowed",
			Errors: "Method Not Allowed",
		})
		return
	}

	levelreq := new(request.Levels)
	if err:= json.NewDecoder(r.Body).Decode(levelreq);err != nil {
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
			Message: "Invalid Id Levels Format",
			Errors: "Bad Request",
		})
		return
	} 

	ctx, cancel := helper.Ctxtimeout(r.Context(), 10*time.Second)
	defer cancel()

	if err:= h.service.UpdateLevel(ctx, id, levelreq);err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.ReturnResponse(w, helper.Messages{
				Http_code: helper.Notfound,
				Message: fmt.Sprintf("Not Found Data Id: %d", id),
				Errors: "Not Found",
			})
			return
		}

		if helper.IsDuplicateEntryError(err) {
			helper.ReturnResponse(w, helper.Messages{
				Http_code: helper.Conflict,
				Message: "Data Is Exist",
				Errors: "Conflict",
			})
			return
		}

		h.logs.Logfile(err.Error())
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.InternalServError,
			Message: "Something Wnet Wrong",
			Errors: "internal Server Error",
		})
		return
	}

	helper.ReturnResponse(w, helper.Messages{
		Http_code: helper.Success,
		Message: "Success Update Level",
	})
}


func (h *HandlerLevel) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != helper.Delete {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.MethodNotAllowed,
			Message: "Only Delete Method Is Allowed",
			Errors: "Method Not Allowed",
		})
		return
	}

	id,err:= strconv.Atoi(r.PathValue("id"))
	if err != nil {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.BadRequest,
			Message: "Invalid Id Level Format",
			Errors: "Bad Request",
		})
		return
	}

	ctx, cancel := helper.Ctxtimeout(r.Context(), 5*time.Second)
	defer cancel()

	if err:=h.service.DeleteLevel(ctx, id);err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.ReturnResponse(w, helper.Messages{
				Http_code: helper.Notfound,
				Message: fmt.Sprintf("Not Found id: %d", id),
				Errors: "Not Found",
			})
			return
		}

		h.logs.Logfile(err.Error())
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.InternalServError,
			Message: "Somethind Went Wrong",
			Errors: "Internal Server Error",
		})
		return
	}

	helper.ReturnResponse(w, helper.Messages{
		Http_code: helper.Success,
		Message: "Success Delete Level",
	})
}