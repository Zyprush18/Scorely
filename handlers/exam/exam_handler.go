package exam

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/service/serviceexam"
	"gorm.io/gorm"
)

type HandlerExam struct {
	service serviceexam.ServiceExams
	logg helper.Loggers
}

func ConnServc(s serviceexam.ServiceExams,l helper.Loggers) HandlerExam {
	return HandlerExam{service: s,logg: l}
}

func (h *HandlerExam) GetALl(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	if r.Method != helper.Gets {
		w.WriteHeader(helper.MethodNotAllowed)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Only Get Method Is Allowed",
			Errors: "Method Not Allowed",
		})
		return
	}

	page,perpage,sort,search,err:=helper.QueryParam(r,10)
	if err != nil {
		w.WriteHeader(helper.BadRequest)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Invalid Query Params Format",
			Errors: "Bad Request",
		})
		return
	}

	resp,count,err := h.service.GetAllExams(search,sort,page,perpage)
	if err != nil {
		h.logg.Logfile(err.Error())
		w.WriteHeader(helper.InternalServError)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "SOmething Went Wrong",
			Errors: "Internal Server Error",
		})
		return
	}

	w.WriteHeader(helper.Success)
	json.NewEncoder(w).Encode(helper.Messages{
		Message: "Success",
		Data: resp,
		Pagination: helper.Paginations(page,perpage,int(count)),
	})
}


func (h *HandlerExam) FindByIdTeacher(w http.ResponseWriter,r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	if r.Method != helper.Gets {
		w.WriteHeader(helper.MethodNotAllowed)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Only Get Method is Allowed",
		})
		return
	}

	// ambil id teacher dari context
	id_teacher := r.Context().Value(helper.KeyTeacherID).(int)
	page,perpage,sort,search,err := helper.QueryParam(r, 10)
	if err != nil {
		w.WriteHeader(helper.BadRequest)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Invalid Query Params Format",
			Errors: "Bad Request",
		})
		return
	}

	resp,count, err := h.service.FindExamsbyIdTeacher(search,sort,page,perpage,id_teacher)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(helper.Notfound)
			json.NewEncoder(w).Encode(helper.Messages{
				Message: fmt.Sprintf("Not Found user id teacher: %d", id_teacher),
				Errors: "Not Found",
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
		Message: "Success",
		Data: resp,
		Pagination: helper.Paginations(page,perpage,int(count)),
	})
}