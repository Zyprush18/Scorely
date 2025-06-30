package major

import (
	"encoding/json"
	"net/http"

	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/service/majorservice"
)

type ServiceMajor struct {
	service majorservice.MajorService
	log helper.Loggers
}

func Handlers(s majorservice.MajorService, l helper.Loggers) ServiceMajor  {
	return ServiceMajor{service: s, log: l}
}

func (s *ServiceMajor) GetAllData(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != helper.Gets {
		w.WriteHeader(helper.MethodNotAllowed)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Method Get is Allowed",
			Errors: "Method Not Allowed",
		})
		return
	}

	page,perpage,sort,search,err:= helper.QueryParam(r, 10)
	if err != nil {
		w.WriteHeader(helper.BadRequest)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Invalid Format Query Params",
			Errors: "Bad Request",
		})
		return
	}

	resp, count, err := s.service.GetAllMajor(search,sort,page,perpage)
	if err != nil {
		s.log.Logfile(err.Error())
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


func (s *ServiceMajor) Create(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != helper.Post {
		w.WriteHeader(helper.MethodNotAllowed)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Method Post Is Allowed",
			Errors: "Method Not Allowed",
		})
		return
	}

	majorreq := new(request.Majors)
	if err := json.NewDecoder(r.Body).Decode(&majorreq); err != nil {
		w.WriteHeader(helper.BadRequest)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Request Body Is Missing",
			Errors: "Bad Request",
		})
		return
	}

	if err := helper.ValidateForm(majorreq);err != nil {
		w.WriteHeader(helper.UnprocessbleEntity)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Validation Error",
			Errors: err.Error(),
		})
		return
	}

	if err:= s.service.CreateMajor(majorreq);err != nil {
		if helper.IsDuplicateEntryError(err) {
			w.WriteHeader(helper.Conflict)
			json.NewEncoder(w).Encode(helper.Messages{
				Message: "Data is Exists",
				Errors: "conflict",
			})
			return
		}

		s.log.Logfile(err.Error())
		w.WriteHeader(helper.InternalServError)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Something Went Wrong",
			Errors: "Internal Server Error",
		})
		return
	}

	w.WriteHeader(helper.Created)
	json.NewEncoder(w).Encode(helper.Messages{
		Message: "Success Create a New Major",
	})

}