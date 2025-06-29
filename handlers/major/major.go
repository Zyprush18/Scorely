package major

import (
	"encoding/json"
	"net/http"

	"github.com/Zyprush18/Scorely/helper"
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
