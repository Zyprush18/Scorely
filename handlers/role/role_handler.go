package role

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/service/servicerole"
	"gorm.io/gorm"
)

type HandlerRole struct {
	services servicerole.ServiceRole
	logg     helper.Loggers
}

func RoleHandler(s servicerole.ServiceRole, l helper.Loggers) *HandlerRole {
	return &HandlerRole{services: s, logg: l}
}

func (h *HandlerRole) GetRole(w http.ResponseWriter, r *http.Request) {
	if r.Method != helper.Gets {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.MethodNotAllowed,
			Message: "Only Get Method Is Allowed",
			Errors:  "Method Not Allowed",
		})
		return
	}

	page, perpage, sort, search, err := helper.QueryParam(r, 10)
	if err != nil {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.BadRequest,
			Message: "Invalid page format",
			Errors:  "Bad Request",
		})

		return
	}

	ctx, cancel := helper.Ctxtimeout(r.Context(), 5*time.Second)
	defer cancel()

	resp, count, err := h.services.GetAllData(ctx, search, sort, page, perpage)
	if err != nil {
		h.logg.Logfile(err.Error())
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.InternalServError,
			Message: "Some thing Wrong",
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

func (h *HandlerRole) AddRoles(w http.ResponseWriter, r *http.Request) {
	if r.Method != helper.Post {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.MethodNotAllowed,
			Message: "Only Post Method Is Allowed",
			Errors:  "Method Not Allowed",
		})
		return
	}

	roleReq := new(request.Roles)

	//cek body form nya
	if err := json.NewDecoder(r.Body).Decode(&roleReq); err != nil {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.BadRequest,
			Message: "Request Body Is Missing",
			Errors:  "Bad Request",
		})
		return
	}

	// validasi
	if err := helper.ValidateForm(roleReq); err != nil {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.UnprocessbleEntity,
			Errors: "Validation Failed",
			Fields: err.Error(),
		})

		return
	}

	ctx, cancel := helper.Ctxtimeout(r.Context(), 10*time.Second)
	defer cancel()

	if err := h.services.Create(ctx, roleReq); err != nil {
		if helper.IsDuplicateEntryError(err) {
			helper.ReturnResponse(w, helper.Messages{
				Http_code: helper.Conflict,
				Message: "Name Role is exists",
				Errors:  "Conflict",
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
		Http_code: helper.Created,
		Message: "Success Create a New Role",
	})

}

func (h *HandlerRole) Show(w http.ResponseWriter, r *http.Request) {
	if r.Method != helper.Gets {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.MethodNotAllowed,
			Message: "Only Get Method Is Allowed",
			Errors:  "Method Not Allowed",
		})
		return
	}

	// ambil id di path url
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		h.logg.Logfile(err.Error())
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.BadRequest,
			Message: "Invalid role ID format",
			Errors:  "Bad Request",
		})

		return
	}

	// ambil data by id
	ctx, cancel := helper.Ctxtimeout(r.Context(), 5*time.Second)
	defer cancel()

	resp, err := h.services.ShowRoleById(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.ReturnResponse(w, helper.Messages{
				Http_code: helper.Notfound,
				Message: fmt.Sprintf("Not Found data by id: %v", id),
				Errors:  "Not Found",
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
		Data:    resp,
	})

}

func (h *HandlerRole) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != helper.Put {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.MethodNotAllowed,
			Message: "Only Put Method Is Allowed",
			Errors:  "Method Not Allowed",
		})
		return
	}
	user := &request.Roles{}

	// cek body nya kosong atau tidak
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.BadRequest,
			Message: "Bad Request",
		})
		return
	}

	// ambil id di path url
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		h.logg.Logfile(err.Error())
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.BadRequest,
			Message: "Invalid role ID format",
			Errors:  "Bad Request",
		})

		return
	}

	ctx, cancel := helper.Ctxtimeout(r.Context(), 10*time.Second)
	defer cancel()

	if err := h.services.UpdateRole(ctx, id, user); err != nil {
		// not found id
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.ReturnResponse(w, helper.Messages{
				Http_code: helper.Notfound,
				Message: fmt.Sprintf("Not Found Id Role: %d", id),
				Errors:  "Bad Request",
			})
			return
		}

		// name role is exist
		if helper.IsDuplicateEntryError(err) {
			helper.ReturnResponse(w, helper.Messages{
				Http_code: helper.Conflict,
				Message: "Name Role is Exist",
				Errors:  "Conflict",
			})
			return
		}

		h.logg.Logfile(err.Error())
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.InternalServError,
			Message: "Something went wrong",
			Errors:  "Internal Server Error",
		})
		return
	}

	helper.ReturnResponse(w, helper.Messages{
		Http_code: helper.Success,
		Message: "Success Update Role",
	})
}

func (h *HandlerRole) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != helper.Delete {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.MethodNotAllowed,
			Message: "Only Delete Method Is Allowed",
			Errors:  "Method Not Allowed",
		})
		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.BadRequest,
			Message: "Invalid role ID format",
			Errors:  "Bad Request",
		})
		return
	}

	ctx, cancel := helper.Ctxtimeout(r.Context(), 5*time.Second)
	defer cancel()

	if err := h.services.DeleteRole(ctx, id); err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.ReturnResponse(w, helper.Messages{
				Http_code: helper.Notfound,
				Message: fmt.Sprintf("Not Found id role: %d", id),
				Errors:  "Not Found",
			})
			return
		}

		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.InternalServError,
			Message: "Something Went Wrong",
			Errors:  "Internal Server Error",
		})
		return
	}

	helper.ReturnResponse(w, helper.Messages{
		Http_code: helper.Success,
		Message: "Success Delete Role",
	})

}
