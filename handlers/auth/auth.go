package auth

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

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
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.MethodNotAllowed,
			Message: "Only Post Method Is Allowed",
			Errors: "Method Not Allowed",
		})
		return
	}

	loginreq := new(request.Login)
	if err := json.NewDecoder(r.Body).Decode(loginreq); err != nil {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.BadRequest,
			Message: "Body Request Is Missing",
			Errors: "Bad Request",
		})
		return
	}

	if err:= helper.ValidateForm(loginreq); err != nil{
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.UnprocessbleEntity,
			Message: "Failed Validation",
			Errors: err.Error(),
		})
		return
	}

	ctx, cancel := helper.Ctxtimeout(r.Context(), 5*time.Second)
	defer cancel()

	token,err := h.servc.Loginuser(ctx, loginreq);
	if  err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.ReturnResponse(w, helper.Messages{
				Http_code: helper.Unauthorized,
				Message: "Incorrect Email or Password",
				Errors: "Unauthorized",
			})
			return
		}

		if errors.Is(err, context.DeadlineExceeded) {
			helper.ReturnResponse(w, helper.Messages{
				Http_code: helper.Timeout,
				Message: "Request Timeout",
				Errors: "Timeout",
			})
		}


		if errors.Is(err, context.Canceled) {
			helper.ReturnResponse(w, helper.Messages{
				Http_code: helper.Cancel,
				Message: "Request Cancelled",
				Errors: "Cancelled",
			})
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
		Message: "Success Login",
		Token: token,
	})
}

func (h *HandlerAuth) Signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != helper.Post {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.MethodNotAllowed,
			Message: "Only Post Method Is Allowed",
			Errors:  "Method Not Allowed",
		})
		return
	}

	register := new(request.Register)
	if err := json.NewDecoder(r.Body).Decode(register); err != nil {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.BadRequest,
			Message: "Body Request Is Missing",
			Errors:  "Bad Request",
		})
		return
	}

	if err := helper.ValidateForm(register); err != nil {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.UnprocessbleEntity,
			Message: "Failed Validation",
			Errors:  err.Error(),
		})
		return
	}

	ctx, cancel := helper.Ctxtimeout(r.Context(), 10*time.Second)
	defer cancel()

	if err := h.servc.Signup(ctx, register); err != nil {
		if helper.IsDuplicateEntryError(err) {
			helper.ReturnResponse(w, helper.Messages{
				Http_code: helper.Conflict,
				Message: "User Already Exists",
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
		Http_code: helper.Created,
		Message: "Success Register",
	})
}

func (h *HandlerAuth) Logout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != helper.Post {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.MethodNotAllowed,
			Message: "Only Post Method Is Allowed",
			Errors:  "Method Not Allowed",
		})
		return
	}

	idteacher, ok := r.Context().Value(helper.KeyUserID).(int)
	if !ok {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.Unauthorized,
			Message: "Unauthorized",
			Errors:  "Unauthorized",
		})
		return
	}

	tokenID, ok := r.Context().Value(helper.KeyTokenID).(string)
	if !ok {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.Unauthorized,
			Message: "Unauthorized",
			Errors:  "Unauthorized",
		})
		return
	}

	ctx, cancel := helper.Ctxtimeout(r.Context(), 5*time.Second)
	defer cancel()

	if err := h.servc.Logout(ctx, idteacher, tokenID); err != nil {
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
		Message: "Success Logout",
	})
}