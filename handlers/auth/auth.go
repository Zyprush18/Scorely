package auth

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/service/serviceauth"
	"gorm.io/gorm"
)

type HandlerAuth struct {
	servc serviceauth.AuthService
	logg helper.Loggers
}

func ConnectService(s serviceauth.AuthService,l helper.Loggers) HandlerAuth  {
	return HandlerAuth{servc: s,logg: l}
}

func (h *HandlerAuth) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	if r.Method != helper.Post {
		w.WriteHeader(helper.MethodNotAllowed)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Only Post Method Is Allowed",
			Errors: "Method Not Allowed",
		})
		return
	}

	loginreq := new(request.Login)
	if err := json.NewDecoder(r.Body).Decode(loginreq); err != nil {
		w.WriteHeader(helper.BadRequest)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Body Request Is Missing",
			Errors: "Bad Request",
		})
		return
	}

	if err:= helper.ValidateForm(loginreq); err != nil{
		w.WriteHeader(helper.UnprocessbleEntity)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Failed Validation",
			Errors: err.Error(),
		})
		return
	}

	token,err := h.servc.Loginuser(loginreq);
	if  err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(helper.Notfound)
			json.NewEncoder(w).Encode(helper.Messages{
				Message: "Email Not Found",
				Errors: "Not Found",
			})
			return
		}

		if err.Error() == "invalid_pw" {
			w.WriteHeader(helper.Unauthorized)
			json.NewEncoder(w).Encode(helper.Messages{
				Message: "Incorrect Email or Password",
				Errors: "Unauthorized",
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
		Message: "Success Login",
		Token: token,
	})
}