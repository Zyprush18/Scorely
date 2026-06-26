package examquestion

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

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
	if r.Method != helper.Gets {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.MethodNotAllowed,
			Message: "Only Get Method Is Allowed",
			Errors: "Methid Not Allowed",
		})
		return
	}

	examid,err := strconv.Atoi(r.PathValue("id_exam"))
	if err != nil {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.BadRequest,
			Message: "Invalid Id Exam Format",
			Errors: "Bad Request",
		})
		return
	}

	userid := r.Context().Value(helper.KeyUserID).(int)
	coderole := r.Context().Value(helper.KeyCodeRole).(string)

	page,perpage,sort,search,err:= helper.QueryParam(r,10)
	if err != nil {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.BadRequest,
			Message: "Invalid Query Param Format",
			Errors: "Bad Request",
		})
		return
	}


	ctx, cancel := helper.Ctxtimeout(r.Context(), 5*time.Second)
	defer cancel()

	resp,count,err := h.service.GetAllExamQuest(ctx,search,sort,coderole,page,perpage,userid,examid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.ReturnResponse(w, helper.Messages{
				Http_code: helper.Notfound,
				Message: "No Relation Found Exam",
				Errors: "Not Found",
			})
			return
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
		Message: "Sucesss",
		Data: resp,
		Pagination: helper.Paginations(page,perpage,int(count)),
	})
}

func (h *HandlerExamQuest) Create(w http.ResponseWriter, r *http.Request)  {
	if r.Method != helper.Post {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.MethodNotAllowed,
			Message: "Only Post Method Is Allowed",
			Errors: "Method Not Allowed",
		})
		return
	}

	idexam,err := strconv.Atoi(r.PathValue("id_exam"))
	if err != nil {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.BadRequest,
			Message: "Invalid Id Exam Format",
			Errors: "Bad Request",
		})
		return
	}

	userid := r.Context().Value(helper.KeyUserID).(int)
	coderole := r.Context().Value(helper.KeyCodeRole).(string)

	reqexamquest := new(request.Exam_Questions)

	if json.NewDecoder(r.Body).Decode(reqexamquest) != nil {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.BadRequest,
			Message: "Body Request Is Missing",
			Errors: "Bad Request",
		})
		return
	}

	if err := helper.ValidateForm(reqexamquest);err != nil {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.UnprocessbleEntity,
			Message: "Failed Validation",
			Errors: err.Error(),
		})
		return
	}


	ctx, cancel := helper.Ctxtimeout(r.Context(), 10*time.Second)
	defer cancel()

	if err := h.service.CreateExamQuest(ctx, reqexamquest,userid,idexam,coderole);err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.ReturnResponse(w, helper.Messages{
				Http_code: helper.Notfound,
				Message: "No Relation Found Exam",
				Errors: "Not Found",
			})
			return
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
		Http_code: helper.Created,
		Message: "Success Created",
	})
}

func (h *HandlerExamQuest) Show(w http.ResponseWriter,r *http.Request)  {
	if r.Method != helper.Gets {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.MethodNotAllowed,
			Message: "Only Get Method Is Allowed",
			Errors: "Method Not Allowed",
		})
		return
	}

	id,errs := strconv.Atoi(r.PathValue("id"))
	id_exam,err := strconv.Atoi(r.PathValue("id_exam"))
	if errs != nil||err != nil {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.BadRequest,
			Message: "Invalid Id Format",
			Errors: "Bad Request",
		})
		return
	}

	user_id := r.Context().Value(helper.KeyUserID).(int)  
	coderole := r.Context().Value(helper.KeyCodeRole).(string)
	
	ctx, cancel := helper.Ctxtimeout(r.Context(), 5*time.Second)
	defer cancel()

	resp,err := h.service.ShowExamQuest(ctx, id,user_id,id_exam,coderole)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.ReturnResponse(w, helper.Messages{
				Http_code: helper.Notfound,
				Message: "Not Found Exam Question or No Relation Found Exam",
				Errors: "Not Found",
			})
			return
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
		Message: "Success",
		Data: resp,
	})

}