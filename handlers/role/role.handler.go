package role

import (
	"encoding/json"
	"net/http"

	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/service/servicerole"
)

type HandlerRole struct {
	services servicerole.ServiceRole
	logg helper.Logger
}

func RoleHandler(s servicerole.ServiceRole, l helper.Logger) *HandlerRole  {
	return &HandlerRole{services: s, logg: l}
}
func (h *HandlerRole) AddRoles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != helper.Post {
		w.WriteHeader(helper.MethodNotAllowed)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Method Not Allowed",
		})
		return
	}

	roleReq := new(request.Roles)

	//cek body form nya
	if err := json.NewDecoder(r.Body).Decode(&roleReq); err != nil {
		w.WriteHeader(helper.BadRequest)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Bad Request",
		})
		return
	}

	// validasi
	if err := helper.ValidateForm(roleReq); err != nil {
		w.WriteHeader(helper.BadRequest)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Validation Failed",
			Errors:  err.Error(),
		})

		return
	}

	if err := h.services.Create(roleReq); err != nil {
		h.logg.Logfile(err.Error())
		w.WriteHeader(helper.BadRequest)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Failed Add Role",
		})
		return
	}

	w.WriteHeader(helper.Created)
	json.NewEncoder(w).Encode(helper.Messages{
		Message: "Success Create a New Role",
	})

}
