package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

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
	w.Header().Set("Content-Type", "application/json")
	if r.Method != helper.Gets {
		w.WriteHeader(helper.MethodNotAllowed)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Only Get Method Is Allowed",
			Errors:  "Method Not Allowed",
		})
		return
	}

	resp, err := h.service.AllUser()
	if err != nil {
		w.WriteHeader(helper.InternalServError)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Something Went Wrong",
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

func (h *UserService) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != helper.Post {
		w.WriteHeader(helper.MethodNotAllowed)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Only Post Method Is Allowed",
			Errors:  "Method Not Allowed",
		})
		return
	}

	userreq := request.User{}
	// cek body request
	if err := json.NewDecoder(r.Body).Decode(&userreq); err != nil {
		w.WriteHeader(helper.BadRequest)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Request Body Is Missing",
			Errors:  "Bad Request",
		})

		return
	}

	// validation
	if err := helper.ValidateForm(&userreq); err != nil {
		w.WriteHeader(helper.UnprocessbleEntity)
		json.NewEncoder(w).Encode(helper.Messages{
			Errors: "Validation Error",
			Fields: err.Error(),
		})

		return
	}

	// userreq.Models.CreatedAt = time.Now()
	// create new user
	err := h.service.CreateUser(&userreq)
	// pengecekan error dari create
	if err != nil {

		// jika terjadi duplicate email
		if helper.IsDuplicateEntryError(err) {
			w.WriteHeader(helper.Conflict)
			json.NewEncoder(w).Encode(helper.Messages{
				Message: "Email Already Exists",
				Errors:  "Conflict",
			})
			return
		}

		// jika terjadi error di database
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
		Message: "Success Create a New User",
	})
}

func (h *UserService) Show(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != helper.Gets {
		w.WriteHeader(helper.MethodNotAllowed)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Only Get Method Is Allowed",
			Errors:  "Method Not Allowed",
		})
		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(helper.BadRequest)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Invalid user ID format",
			Errors:  "Bad Request",
		})
		return
	}

	data, err := h.service.ShowUser(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(helper.Notfound)
			json.NewEncoder(w).Encode(helper.Messages{
				Message: fmt.Sprintf("Not Found Id User: %d", id),
				Errors:  "Bad Request",
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

	w.WriteHeader(helper.Success)
	json.NewEncoder(w).Encode(helper.Messages{
		Message: "Success",
		Data:    data,
	})
}

func (h *UserService) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != helper.Put {
		w.WriteHeader(helper.MethodNotAllowed)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Only Put Method Is Allowed",
			Errors:  "Method Not Allowed",
		})

		return
	}

	userreq := new(request.User)
	if err := json.NewDecoder(r.Body).Decode(&userreq); err != nil {
		w.WriteHeader(helper.BadRequest)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Request Body Is Missing",
			Errors:  "Bad Request",
		})

		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(helper.BadRequest)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Invalid user ID format",
			Errors:  "Bad Request",
		})

		return
	}

	if err := h.service.UpdateUser(id, userreq); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(helper.Notfound)
			json.NewEncoder(w).Encode(helper.Messages{
				Message: fmt.Sprintf("Not Found Id: %d", id),
				Errors:  "Not Found",
			})

			return
		}

		if helper.IsDuplicateEntryError(err) {
			w.WriteHeader(helper.Conflict)
			json.NewEncoder(w).Encode(helper.Messages{
				Message: "Email Already Exists",
				Errors:  "Conflict",
			})

			return
		}

		h.logg.Logfile(err.Error())
		w.WriteHeader(helper.InternalServError)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Something Went Wrong",
			Errors:  "Internal Server Error",
		})
		return
	}

	w.WriteHeader(helper.Success)
	json.NewEncoder(w).Encode(helper.Messages{
		Message: "Success Update Data",
	})
}

func (h *UserService) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != helper.Delete {
		w.WriteHeader(helper.MethodNotAllowed)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Only Method Delete Is Allowed",
			Errors:  "Method Not Allowed",
		})
		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(helper.BadRequest)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Invalid User Id format",
			Errors:  "Bad Request",
		})
		return
	}

	if err := h.service.DeleteUser(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(helper.Notfound)
			json.NewEncoder(w).Encode(helper.Messages{
				Message: fmt.Sprintf("Not Found id: %d", id),
				Errors:  "Not Found",
			})
			return
		}

		h.logg.Logfile(err.Error())
		w.WriteHeader(helper.InternalServError)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Something Went Wrong",
			Errors:  "Internal Server Error",
		})
		return
	}

	w.WriteHeader(helper.Success)
	json.NewEncoder(w).Encode(helper.Messages{
		Message: "Success Delete User",
	})
}
