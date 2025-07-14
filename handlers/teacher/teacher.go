package teacher

import (
	"encoding/json"
	"net/http"

	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/service/serviceteacher"
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