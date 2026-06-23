package exam

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

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
	if r.Method != helper.Gets {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.MethodNotAllowed,
			Message: "Only Get Method Is Allowed",
			Errors:  "Method Not Allowed",
		})
		return
	}

	page, perpage, sort, search, err := helper.QueryParam(r, 10)
	if err != nil {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.BadRequest,
			Message: "Invalid Query Params Format",
			Errors:  "Bad Request",
		})
		return
	}

	ctx, cancel := helper.Ctxtimeout(r.Context(), 5*time.Second)
	defer cancel()

	resp, count, err := h.service.GetAllExams(ctx, search, sort, page, perpage)
	if err != nil {
		h.logg.Logfile(err.Error())
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.InternalServError,
			Message: "SOmething Went Wrong",
			Errors:  "Internal Server Error",
		})
		return
	}

	helper.ReturnResponse(w, helper.Messages{
		Http_code: helper.Success,
		Message:    "Success",
		Data:       resp,
		Pagination: helper.Paginations(page, perpage, int(count)),
	})
}

func (h *HandlerExam) FindByIdTeacher(w http.ResponseWriter, r *http.Request) {
	if r.Method != helper.Gets {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.MethodNotAllowed,
			Message: "Only Get Method is Allowed",
		})
		return
	}

	// ambil id teacher dari context
	id_teacher := r.Context().Value(helper.KeyUserID).(int)
	page, perpage, sort, search, err := helper.QueryParam(r, 10)
	if err != nil {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.BadRequest,
			Message: "Invalid Query Params Format",
			Errors:  "Bad Request",
		})
		return
	}

	ctx, cancel := helper.Ctxtimeout(r.Context(), 5*time.Second)
	defer cancel()

	resp, count, err := h.service.FindExamsbyIdTeacher(ctx, search, sort, page, perpage, id_teacher)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.ReturnResponse(w, helper.Messages{
				Http_code: helper.Notfound,
				Message: fmt.Sprintf("Not Found user id teacher: %d", id_teacher),
				Errors:  "Not Found",
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
		Http_code: helper.Success,
		Message:    "Success",
		Data:       resp,
		Pagination: helper.Paginations(page, perpage, int(count)),
	})
}

func (h *HandlerExam) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != helper.Post {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.MethodNotAllowed,
			Message: "Only Post Method Is Allowed",
			Errors:  "Method Not Allowed",
		})
		return
	}

	subjectid, err := strconv.Atoi(r.PathValue("subject_id"))
	if err != nil {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.BadRequest,
			Message: "Invalid Id Format",
			Errors: "Bad Request",
		})
		return
	}

	reqexams := new(request.Exams)
	if json.NewDecoder(r.Body).Decode(reqexams) != nil {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.BadRequest,
			Message: "Body Request Is Missing",
			Errors: "Bad Request",
		})
		return
	}

	if err := helper.ValidateForm(reqexams); err != nil {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.UnprocessbleEntity,
			Message: "Failed Validation",
			Errors: err.Error(),
		})
		return
	}

	user_id := r.Context().Value(helper.KeyUserID).(int)
	role := r.Context().Value(helper.KeyCodeRole).(string)


	ctx, cancel := helper.Ctxtimeout(r.Context(), 10*time.Second)
	defer cancel()

	if err := h.service.CreateExams(ctx, reqexams,role,user_id,subjectid);err != nil {
		if errors.Is(err , gorm.ErrRecordNotFound) || errors.Is(err, gorm.ErrCheckConstraintViolated) {
			helper.ReturnResponse(w, helper.Messages{
				Http_code: helper.Notfound,
				Message: "no relation found between teacher and subject",
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

func (h *HandlerExam) Show(w http.ResponseWriter,r *http.Request)  {
	if r.Method != helper.Gets {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.MethodNotAllowed,
			Message: "Only Get Method Is Allowed",
			Errors: "Method Not Allowed",
		})
		return
	}

	id,err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.BadRequest,
			Message: "Invalid Format Param Id",
			Errors: "Bad Request",
		})
		return
	}

	user_id := r.Context().Value(helper.KeyUserID).(int)
	role := r.Context().Value(helper.KeyCodeRole).(string)

	ctx, cancel := helper.Ctxtimeout(r.Context(), 5*time.Second)
	defer cancel()

	resp,err := h.service.ShowExams(ctx, id,user_id,role)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.ReturnResponse(w, helper.Messages{
				Http_code: helper.Notfound,
				Message: "Exam not found",
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

func (h *HandlerExam) Update(w http.ResponseWriter,r *http.Request)  {
	if r.Method != helper.Put {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.MethodNotAllowed,
			Message: "Only Put Method Not Allowed",
			Errors: "Method Not Allowed",
		})
		return
	}

	id,err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.BadRequest,
			Message: "Invalid Params Id Format",
			Errors: "Bad Request",
		})
		return
	}

	updatereq := new(request.Exams)

	if json.NewDecoder(r.Body).Decode(updatereq) != nil {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.BadRequest,
			Message: "Request Body Is Missing",
			Errors: "Bad Request",
		})
		return
	}

	userid := r.Context().Value(helper.KeyUserID).(int)
	role := r.Context().Value(helper.KeyCodeRole).(string)


	ctx, cancel := helper.Ctxtimeout(r.Context(), 10*time.Second)
	defer cancel()

	if err= h.service.UpdateExam(ctx, updatereq,role,id,userid);err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.ReturnResponse(w, helper.Messages{
				Http_code: helper.Notfound,
				Message: "Exam Not Found Or No Relation Found Between Teacher and Subject",
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
		Message: "Success Update exams",
	})

}

func (h *HandlerExam) Delete(w http.ResponseWriter, r *http.Request)  {
	if r.Method != helper.Delete {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.MethodNotAllowed,
			Message: "Only Delete Method Is Allowed",
			Errors: "Method Not Allowed",
		})
		return
	}

	id,err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		helper.ReturnResponse(w, helper.Messages{
			Http_code: helper.BadRequest,
			Message: "Invalid Id Params Format",
			Errors: "Bad Request",
		})
		return
	}

	userid := r.Context().Value(helper.KeyUserID).(int)
	role := r.Context().Value(helper.KeyCodeRole).(string)

	ctx, cancel := helper.Ctxtimeout(r.Context(), 5*time.Second)
	defer cancel()

	if err := h.service.DeleteExam(ctx, id, userid,role);err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.ReturnResponse(w, helper.Messages{
				Http_code: helper.Notfound,
				Message: "Exam Not Found Or No Relation Found Between Teacher and Subject",
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
		Message: "Success Delete",
	})
}