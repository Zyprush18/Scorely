package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/service/userservice"
	"gorm.io/gorm"
)

type UserService struct {
	service userservice.ServiceUser
	logg    helper.Loggers
}

func NewHandlerUser(service userservice.ServiceUser, log helper.Loggers) UserService {
	return UserService{service: service, logg: log}
}

func (h *UserService) GetAllUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != helper.Gets {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.MethodNotAllowed,
			Message: "Only Get Method Is Allowed",
			Errors:  "Method Not Allowed",
		})
		return
	}

	page,perpage,sort,search, err:=helper.QueryParam(r, 10)
	if err != nil {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.BadRequest,
			Message: "Invalid Query Params Format",
			Errors:  "Bad Request",
		})
		return
	}

	ctx, cancel := helper.Ctxtimeout(r.Context(), 5*time.Second)
	defer cancel()

	resp, count,err := h.service.AllUser(ctx,search,sort,page,perpage)
	if err != nil {
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
		Data:    resp,
		Pagination: helper.Paginations(page,perpage,int(count)),
	})

}

func (h *UserService) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != helper.Post {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.MethodNotAllowed,
			Message: "Only Post Method Is Allowed",
			Errors:  "Method Not Allowed",
		})
		return
	}

	userreq := request.User{}
	// cek body request
	if err := json.NewDecoder(r.Body).Decode(&userreq); err != nil {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.BadRequest,
			Message: "Request Body Is Missing",
			Errors:  "Bad Request",
		})

		return
	}

	// validation
	if err := helper.ValidateForm(&userreq); err != nil {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.UnprocessbleEntity,
			Errors: "Validation Error",
			Fields: err.Error(),
		})

		return
	}

	ctx, cancel := helper.Ctxtimeout(r.Context(), 10*time.Second)
	defer cancel()

	if err := h.service.CreateUser(ctx, &userreq);err != nil {

		// jika terjadi duplicate email
		if helper.IsDuplicateEntryError(err) {
			helper.ReturnResponse(w, helper.Messages{
				Http_code: helper.Conflict,
				Message: "Email Already Exists",
				Errors:  "Conflict",
			})
			return
		}

		// jika terjadi error di database
		h.logg.Logfile(err.Error())
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.InternalServError,
			Message: "Some Thing Wrong",
			Errors:  "Internal Server Error",
		})
		return
	}

	helper.ReturnResponse(w, helper.Messages{
		Http_code: helper.Created,
		Message: "Success Create a New User",
	})
}

func (h *UserService) Show(w http.ResponseWriter, r *http.Request) {
	if r.Method != helper.Gets {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.MethodNotAllowed,
			Message: "Only Get Method Is Allowed",
			Errors:  "Method Not Allowed",
		})
		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.BadRequest,
			Message: "Invalid user ID format",
			Errors:  "Bad Request",
		})
		return
	}

	ctx, cancel := helper.Ctxtimeout(r.Context(), 5*time.Second)
	defer cancel()

	data, err := h.service.ShowUser(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.ReturnResponse(w, helper.Messages{
				Http_code: helper.Notfound,
				Message: fmt.Sprintf("Not Found Id User: %d", id),
				Errors:  "Bad Request",
			})
			return
		}
		h.logg.Logfile(err.Error())
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.InternalServError,
			Message: "Some Thing Wrong",
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

func (h *UserService) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != helper.Put {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.MethodNotAllowed,
			Message: "Only Put Method Is Allowed",
			Errors:  "Method Not Allowed",
		})

		return
	}

	userreq := new(request.User)
	if err := json.NewDecoder(r.Body).Decode(&userreq); err != nil {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.BadRequest,
			Message: "Request Body Is Missing",
			Errors:  "Bad Request",
		})

		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.BadRequest,
			Message: "Invalid user ID format",
			Errors:  "Bad Request",
		})

		return
	}

	ctx, cancel := helper.Ctxtimeout(r.Context(), 10*time.Second)
	defer cancel()

	if err := h.service.UpdateUser(ctx, id, userreq); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.ReturnResponse(w, helper.Messages{
				Http_code: helper.Notfound,
				Message: fmt.Sprintf("Not Found Id: %d", id),
				Errors:  "Not Found",
			})

			return
		}

		if helper.IsDuplicateEntryError(err) {
			helper.ReturnResponse(w, helper.Messages{
				Http_code: helper.Conflict,
				Message: "Email Already Exists",
				Errors:  "Conflict",
			})

			return
		}

		h.logg.Logfile(err.Error())
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

func (h *UserService) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != helper.Delete {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.MethodNotAllowed,
			Message: "Only Method Delete Is Allowed",
			Errors:  "Method Not Allowed",
		})
		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.BadRequest,
			Message: "Invalid User Id format",
			Errors:  "Bad Request",
		})
		return
	}

	ctx, cancel := helper.Ctxtimeout(r.Context(), 5*time.Second)
	defer cancel()

	if err := h.service.DeleteUser(ctx, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.ReturnResponse(w, helper.Messages{
				Http_code: helper.Notfound,
				Message: fmt.Sprintf("Not Found id: %d", id),
				Errors:  "Not Found",
			})
			return
		}

		h.logg.Logfile(err.Error())
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.InternalServError,
			Message: "Something Went Wrong",
			Errors:  "Internal Server Error",
		})
		return
	}

	helper.ReturnResponse(w, helper.Messages{
		Http_code: helper.Success,
		Message: "Success Delete User",
	})
}
