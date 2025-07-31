package exam

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/service/serviceexam"
	"gorm.io/gorm"
)

type HandlerExam struct {
	service serviceexam.ServiceExams
	logg    helper.Loggers
}

func ConnServc(s serviceexam.ServiceExams, l helper.Loggers) HandlerExam {
	return HandlerExam{service: s, logg: l}
}

func (h *HandlerExam) GetALl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != helper.Gets {
		w.WriteHeader(helper.MethodNotAllowed)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Only Get Method Is Allowed",
			Errors:  "Method Not Allowed",
		})
		return
	}

	page, perpage, sort, search, err := helper.QueryParam(r, 10)
	if err != nil {
		w.WriteHeader(helper.BadRequest)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Invalid Query Params Format",
			Errors:  "Bad Request",
		})
		return
	}

	resp, count, err := h.service.GetAllExams(search, sort, page, perpage)
	if err != nil {
		h.logg.Logfile(err.Error())
		w.WriteHeader(helper.InternalServError)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "SOmething Went Wrong",
			Errors:  "Internal Server Error",
		})
		return
	}

	w.WriteHeader(helper.Success)
	json.NewEncoder(w).Encode(helper.Messages{
		Message:    "Success",
		Data:       resp,
		Pagination: helper.Paginations(page, perpage, int(count)),
	})
}

func (h *HandlerExam) FindByIdTeacher(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != helper.Gets {
		w.WriteHeader(helper.MethodNotAllowed)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Only Get Method is Allowed",
		})
		return
	}

	// ambil id teacher dari context
	id_teacher := r.Context().Value(helper.KeyUserID).(int)
	page, perpage, sort, search, err := helper.QueryParam(r, 10)
	if err != nil {
		w.WriteHeader(helper.BadRequest)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Invalid Query Params Format",
			Errors:  "Bad Request",
		})
		return
	}

	resp, count, err := h.service.FindExamsbyIdTeacher(search, sort, page, perpage, id_teacher)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(helper.Notfound)
			json.NewEncoder(w).Encode(helper.Messages{
				Message: fmt.Sprintf("Not Found user id teacher: %d", id_teacher),
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
		Message:    "Success",
		Data:       resp,
		Pagination: helper.Paginations(page, perpage, int(count)),
	})
}

func (h *HandlerExam) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != helper.Post {
		w.WriteHeader(helper.MethodNotAllowed)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Only Post Method Is Allowed",
			Errors:  "Method Not Allowed",
		})
		return
	}

	subjectid, err := strconv.Atoi(r.PathValue("subject_id"))
	if err != nil {
		w.WriteHeader(helper.BadRequest)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Invalid Id Format",
			Errors: "Bad Request",
		})
		return
	}

	reqexams := new(request.Exams)
	if json.NewDecoder(r.Body).Decode(reqexams) != nil {
		w.WriteHeader(helper.BadRequest)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Body Request Is Missing",
			Errors: "Bad Request",
		})
		return
	}

	if err := helper.ValidateForm(reqexams); err != nil {
		w.WriteHeader(helper.UnprocessbleEntity)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Failed Validation",
			Errors: err.Error(),
		})
		return
	}

	user_id := r.Context().Value(helper.KeyUserID).(int)
	role := r.Context().Value(helper.KeyCodeRole).(string)


	if err := h.service.CreateExams(reqexams,role,user_id,subjectid);err != nil {
		if errors.Is(err , gorm.ErrRecordNotFound) || errors.Is(err, gorm.ErrCheckConstraintViolated) {
			w.WriteHeader(helper.Notfound)
			json.NewEncoder(w).Encode(helper.Messages{
				Message: "no relation found between teacher and subject",
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

	w.WriteHeader(helper.Created)
	json.NewEncoder(w).Encode(helper.Messages{
		Message: "Success Created",
	})
}

func (h *HandlerExam) Show(w http.ResponseWriter,r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	if r.Method != helper.Gets {
		w.WriteHeader(helper.MethodNotAllowed)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Only Get Method Is Allowed",
			Errors: "Method Not Allowed",
		})
		return
	}

	id,err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(helper.BadRequest)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Invalid Format Param Id",
			Errors: "Bad Request",
		})
		return
	}

	user_id := r.Context().Value(helper.KeyUserID).(int)
	role := r.Context().Value(helper.KeyCodeRole).(string)

	resp,err := h.service.ShowExams(id,user_id,role)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(helper.Notfound)
			json.NewEncoder(w).Encode(helper.Messages{
				Message: "Exam not found",
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
	})
}

func (h *HandlerExam) Update(w http.ResponseWriter,r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	if r.Method != helper.Put {
		w.WriteHeader(helper.MethodNotAllowed)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Only Put Method Not Allowed",
			Errors: "Method Not Allowed",
		})
		return
	}

	id,err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(helper.BadRequest)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Invalid Params Id Format",
			Errors: "Bad Request",
		})
		return
	}

	updatereq := new(request.Exams)

	if json.NewDecoder(r.Body).Decode(updatereq) != nil {
		w.WriteHeader(helper.BadRequest)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Request Body Is Missing",
			Errors: "Bad Request",
		})
		return
	}

	userid := r.Context().Value(helper.KeyUserID).(int)
	role := r.Context().Value(helper.KeyCodeRole).(string)


	if err= h.service.UpdateExam(updatereq,role,id,userid);err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(helper.Notfound)
			json.NewEncoder(w).Encode(helper.Messages{
				Message: "Exam Not Found Or No Relation Found Between Teacher and Subject",
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
		Message: "Success Update exams",
	})

}