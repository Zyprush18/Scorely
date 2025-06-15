package role

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

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
	w.Header().Set("Content-Type", "application/json")
	if r.Method != helper.Gets {
		w.WriteHeader(helper.MethodNotAllowed)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Only Get Method Is Allowed",
			Errors:  "Method Not Allowed",
		})
		return
	}

	resp, err := h.services.GetAllData()
	if err != nil {
		h.logg.Logfile(err.Error())
		w.WriteHeader(helper.InternalServError)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Some thing Wrong",
			Errors:  "Internal Server Error",
		})
		return
	}

	w.WriteHeader(helper.Success)
	json.NewEncoder(w).Encode(helper.Messages{
		Message: "Success",
		Data:    resp,
	})
}

func (h *HandlerRole) AddRoles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != helper.Post {
		w.WriteHeader(helper.MethodNotAllowed)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Only Post Method Is Allowed",
			Errors:  "Method Not Allowed",
		})
		return
	}

	roleReq := new(request.Roles)

	//cek body form nya
	if err := json.NewDecoder(r.Body).Decode(&roleReq); err != nil {
		w.WriteHeader(helper.BadRequest)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Request Body Is Missing",
			Errors:  "Bad Request",
		})
		return
	}

	// validasi
	if err := helper.ValidateForm(roleReq); err != nil {
		w.WriteHeader(helper.UnprocessbleEntity)
		json.NewEncoder(w).Encode(helper.Messages{
			Errors: "Validation Failed",
			Fields: err.Error(),
		})

		return
	}

	if err := h.services.Create(roleReq); err != nil {
		if helper.IsDuplicateEntryError(err) {
			w.WriteHeader(helper.Conflict)
			json.NewEncoder(w).Encode(helper.Messages{
				Message: "Name Role is exists",
				Errors:  "Conflict",
			})

			return
		}
		h.logg.Logfile(err.Error())
		w.WriteHeader(helper.InternalServError)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Some Thing Wrong",
			Errors:  "Internal Server Error",
		})
		return
	}

	w.WriteHeader(helper.Created)
	json.NewEncoder(w).Encode(helper.Messages{
		Message: "Success Create a New Role",
	})

}

func (h *HandlerRole) Show(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != helper.Gets {
		w.WriteHeader(helper.MethodNotAllowed)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Only Get Method Is Allowed",
			Errors:  "Method Not Allowed",
		})
		return
	}

	// ambil id di path url
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		h.logg.Logfile(err.Error())
		w.WriteHeader(helper.BadRequest)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Invalid role ID format",
			Errors:  "Bad Request",
		})

		return
	}

	// ambil data by id
	resp, err := h.services.ShowRoleById(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		w.WriteHeader(helper.Notfound)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: fmt.Sprintf("Not Found data by id: %v", id),
			Errors:  "Not Found",
		})
		return
	}
	if err != nil {
		h.logg.Logfile(err.Error())
		w.WriteHeader(helper.InternalServError)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Some Thing Wrong",
			Errors:  "Internal Server Error",
		})
		return
	}

	w.WriteHeader(helper.Success)
	json.NewEncoder(w).Encode(helper.Messages{
		Message: "Success",
		Data:    resp,
	})

}

func (h *HandlerRole) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != helper.Put {
		w.WriteHeader(helper.MethodNotAllowed)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Only Put Method Is Allowed",
			Errors:  "Method Not Allowed",
		})
		return
	}
	user := &request.Roles{}

	// cek body nya kosong atau tidak
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		w.WriteHeader(helper.BadRequest)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Bad Request",
		})
		return
	}

	// ambil id di path url
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		h.logg.Logfile(err.Error())
		w.WriteHeader(helper.BadRequest)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Invalid role ID format",
			Errors:  "Bad Request",
		})

		return
	}

	if err := h.services.UpdateRole(id, user); err != nil {
		// not found id
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(helper.Notfound)
			json.NewEncoder(w).Encode(helper.Messages{
				Message: fmt.Sprintf("Not Found Id User: %d", id),
				Errors:  "Bad Request",
			})
			return
		}

		// name role is exist
		if helper.IsDuplicateEntryError(err) {
			w.WriteHeader(helper.Conflict)
			json.NewEncoder(w).Encode(helper.Messages{
				Message: "Name Role is Exist",
				Errors:  "Conflict",
			})
			return
		}

		h.logg.Logfile(err.Error())
		w.WriteHeader(helper.InternalServError)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Something went wrong",
			Errors:  "Internal Server Error",
		})
		return
	}

	w.WriteHeader(helper.Success)
	json.NewEncoder(w).Encode(helper.Messages{
		Message: "Success Update Role",
	})
}

func (h *HandlerRole) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applivation/json")
	if r.Method != helper.Delete {
		w.WriteHeader(helper.MethodNotAllowed)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Only Delete Method Is Allowed",
			Errors:  "Method Not Allowed",
		})
		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(helper.BadRequest)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Invalid role ID format",
			Errors:  "Bad Request",
		})
		return
	}

	if err := h.services.DeleteRole(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(helper.Notfound)
			json.NewEncoder(w).Encode(helper.Messages{
				Message: fmt.Sprintf("Not Found id role: %d", id),
				Errors: "Not Found",
			})
			return
		}
		w.WriteHeader(helper.InternalServError)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Something Went Wrong",
			Errors: "Internal Server Error",
		})
		return
	}

	w.WriteHeader(helper.Success)
	json.NewEncoder(w).Encode(helper.Messages{
		Message: "Success Delete Role",
	})

}
