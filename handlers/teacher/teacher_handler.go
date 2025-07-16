package teacher

import (
	"encoding/json"
	"errors"
	"net/http"

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
	w.Header().Set("Content-Type","application/json")
	if r.Method != helper.Gets {
		w.WriteHeader(helper.MethodNotAllowed)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Only Get Method Is Allowed",
			Errors: "Method Not Found",
		})
		return
	}

	page,perpage,sort,search,err:=helper.QueryParam(r, 10)
	if err != nil {
		w.WriteHeader(helper.BadRequest)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Invalid Format Query Params",
			Errors: "Bad Request",
		})
		return
	}

	resp,count, err := h.serviice.GetAllTeacher(search,sort,page,perpage)
	if err != nil {
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
		Message: "Success",
		Data: resp,
		Pagination: helper.Paginations(page,perpage, int(count)),
	})

}


func (h *HandlerTeacher) Create(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	if r.Method != helper.Post {
		w.WriteHeader(helper.MethodNotAllowed)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Only Post Method Is Allowed",
			Errors: "Method Not Allowed",
		})
		return
	}


	datareq := new(request.Teachers)
	if err := json.NewDecoder(r.Body).Decode(datareq);err != nil {
		w.WriteHeader(helper.BadRequest)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Body request Is Missing",
			Errors: "Bad Request",
		})
		return
	}

	if err := helper.ValidateForm(datareq); err != nil {
		w.WriteHeader(helper.UnprocessbleEntity)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Failed Validation",
			Errors: err.Error(),
		})
		return
	}

	if err := h.serviice.CreateTeacher(datareq);err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(helper.Notfound)
			json.NewEncoder(w).Encode(helper.Messages{
				Message: "Subject Not Found",
				Errors: "Not Found",
			})
			return
		}

		if helper.IsDuplicateEntryError(err) {
			w.WriteHeader(helper.Conflict)
			json.NewEncoder(w).Encode(helper.Messages{
				Message: "Nip Is Already Exists",
				Errors: "Conflict",
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


	w.WriteHeader(helper.Created)
	json.NewEncoder(w).Encode(helper.Messages{
		Message: "Success Create new Teacher",
	})
}