package role

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/service/servicerole"
)

type HandlerRole struct {
	services servicerole.ServiceRole
	logg helper.Loggers
}

func RoleHandler(s servicerole.ServiceRole, l helper.Loggers) *HandlerRole  {
	return &HandlerRole{services: s, logg: l}
}


func (h *HandlerRole) GetRole(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != helper.Gets {
		w.WriteHeader(helper.MethodNotAllowed)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Method Not Allowed",
		})
		return
	}

	resp, err := h.services.GetAllData()
	if err != nil {
		h.logg.Logfile(err.Error())
		w.WriteHeader(helper.BadRequest)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Failed Get All Data Role",
		})
		return
	}

	w.WriteHeader(helper.Success)
	json.NewEncoder(w).Encode(helper.Messages{
		Message: "Success",
		Data: resp,
	})
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


func (h *HandlerRole) Show(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != helper.Gets {
		w.WriteHeader(helper.MethodNotAllowed)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Method Not Allowed",
		})
		return
	}

	// ambil id di path url
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		h.logg.Logfile(err.Error())
		w.WriteHeader(helper.InternalServError)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Internal Server Error",
		})

		return
	}

	// ambil data by id
	resp,err:= h.services.ShowRoleById(id);
	if err != nil {
		h.logg.Logfile(err.Error())
		w.WriteHeader(helper.Notfound)
		msg := fmt.Sprintf("Not Found data by id: %v", id)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: msg,
		})
		return
	}

	w.WriteHeader(helper.Success)
	json.NewEncoder(w).Encode(helper.Messages{
		Message: "Success",
		Data: resp,
	})
	
}

func (h *HandlerRole) Update(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != helper.Put {
		w.WriteHeader(helper.MethodNotAllowed)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Method Not Allowed",
		})
		return
	}
	user := &request.Roles{}

	// cek body nya kosong atau tidak
	if err:=json.NewDecoder(r.Body).Decode(user);err != nil {
		w.WriteHeader(helper.BadRequest)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Bad Request",
		})
		return
	}

	// ambil id di path url
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		h.logg.Logfile(err.Error())
		w.WriteHeader(helper.InternalServError)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Internal Server Error",
		})

		return
	}

	if err:=h.services.UpdateRole(id,user);err != nil {
		h.logg.Logfile(err.Error())
		w.WriteHeader(helper.BadRequest)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Failed Update Role",
		})
		return
	}

	w.WriteHeader(helper.Success)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Success Update Role",
		})
}

func (h *HandlerRole) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","applivation/json")
	if r.Method != helper.Delete {
		w.WriteHeader(helper.MethodNotAllowed)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Method Not Allowed",
		})
		return 
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(helper.InternalServError)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Internal Server Error",
		})
		return 
	}

	if err:= h.services.DeleteRole(id);err != nil {
		w.WriteHeader(helper.Notfound)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Failed Delete Role",
		})
		return
	}

	w.WriteHeader(helper.Success)
	json.NewEncoder(w).Encode(helper.Messages{
		Message: "Success Delete Role",
	})

}