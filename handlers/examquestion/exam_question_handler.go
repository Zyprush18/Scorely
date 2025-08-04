package examquestion

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/service/serviceexamquest"
	"gorm.io/gorm"
)

type HandlerExamQuest struct {
	service serviceexamquest.ServiceExamQuest
	logg helper.Loggers
}

func ConnectService(s serviceexamquest.ServiceExamQuest,l helper.Loggers) HandlerExamQuest  {
	return HandlerExamQuest{service: s,logg: l}
}

func (h *HandlerExamQuest) GetAll(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	if r.Method != helper.Gets {
		w.WriteHeader(helper.MethodNotAllowed)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Only Get Method Is Allowed",
			Errors: "Methid Not Allowed",
		})
		return
	}

	examid,err := strconv.Atoi(r.PathValue("id_exam"))
	if err != nil {
		w.WriteHeader(helper.BadRequest)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Invalid Id Exam Format",
			Errors: "Bad Request",
		})
		return
	}

	userid := r.Context().Value(helper.KeyUserID).(int)
	coderole := r.Context().Value(helper.KeyCodeRole).(string)

	page,perpage,sort,search,err:= helper.QueryParam(r,10)
	if err != nil {
		w.WriteHeader(helper.BadRequest)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Invalid Query Param Format",
			Errors: "Bad Request",
		})
		return
	}


	resp,count,err := h.service.GetAllExamQuest(search,sort,coderole,page,perpage,userid,examid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(helper.Notfound)
			json.NewEncoder(w).Encode(helper.Messages{
				Message: "No Relation Found Exam",
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
		Message: "Sucesss",
		Data: resp,
		Pagination: helper.Paginations(page,perpage,int(count)),
	})
}

func (h *HandlerExamQuest) Create(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	if r.Method != helper.Post {
		w.WriteHeader(helper.MethodNotAllowed)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Only Post Method Is Allowed",
			Errors: "Method Not Allowed",
		})
		return
	}

	idexam,err := strconv.Atoi(r.PathValue("id_exam"))
	if err != nil {
		w.WriteHeader(helper.BadRequest)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Invalid Id Exam Format",
			Errors: "Bad Request",
		})
		return
	}

	userid := r.Context().Value(helper.KeyUserID).(int)
	coderole := r.Context().Value(helper.KeyCodeRole).(string)

	reqexamquest := new(request.Exam_Questions)

	if json.NewDecoder(r.Body).Decode(reqexamquest) != nil {
		w.WriteHeader(helper.BadRequest)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Body Request Is Missing",
			Errors: "Bad Request",
		})
		return
	}

	if err := helper.ValidateForm(reqexamquest);err != nil {
		w.WriteHeader(helper.UnprocessbleEntity)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Failed Validation",
			Errors: err.Error(),
		})
		return
	}


	if err := h.service.CreateExamQuest(reqexamquest,userid,idexam,coderole);err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(helper.Notfound)
			json.NewEncoder(w).Encode(helper.Messages{
				Message: "No Relation Found Exam",
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