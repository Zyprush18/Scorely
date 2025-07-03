package level

import (
	"encoding/json"
	"net/http"

	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/service/servicelevel"
)

type HandlerLevel struct {
	service servicelevel.LevelService
	logs	helper.Loggers
}

func ConnectService(s servicelevel.LevelService, l helper.Loggers) HandlerLevel  {
	return HandlerLevel{service: s,logs: l}
}

func (h *HandlerLevel) GetAll(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != helper.Gets {
		w.WriteHeader(helper.MethodNotAllowed)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Method Get Is Allowed",
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


	resp, count, err := h.service.GetAllLevel(search,sort,page,perpage)
	if err != nil {
		h.logs.Logfile(err.Error())
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