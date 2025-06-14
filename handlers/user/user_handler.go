package user

import (
	"encoding/json"
	"net/http"

	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/service/userservice"
)

type UserService struct {
	service userservice.ServiceUser
	logg helper.Loggers
}

func NewHandlerUser(service userservice.ServiceUser, log helper.Loggers) UserService  {
	return UserService{service: service, logg: log}
}

func (h *UserService) Create(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	if r.Method != helper.Post {
		w.WriteHeader(helper.MethodNotAllowed)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Only Post Method Is Allowed",
			Errors: "Method Not Allowed",
		})
		return
	}

	userreq := request.User{}
	// cek body request
	if err:= json.NewDecoder(r.Body).Decode(&userreq);err != nil {
		w.WriteHeader(helper.BadRequest)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Request Body Is Missing",
			Errors: "Bad Request",
		})

		return
	}

	// validation
	if err:= helper.ValidateForm(&userreq);err != nil {
		w.WriteHeader(helper.UnprocessbleEntity)
		json.NewEncoder(w).Encode(helper.Messages{
			Errors: "Validation Error",
			Fields: err.Error(),
		})

		return
	}

	// create new user
	if err:= h.service.CreateUser(&userreq);err != nil {
		h.logg.Logfile(err.Error())
		w.WriteHeader(helper.BadRequest)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Failed Create New User",
			Errors: "Bad Request",
		})

		return
	}

	w.WriteHeader(helper.Created)
	json.NewEncoder(w).Encode(helper.Messages{
		Message: "Success Create a New User",
	})
}