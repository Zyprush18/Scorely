package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/service"
	"github.com/Zyprush18/Scorely/utils"
)

func AddRoles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != utils.Post {
		w.WriteHeader(utils.MethodNotAllowed)
		json.NewEncoder(w).Encode(utils.Messages{
			Message: "Method Not Allowed",
		})
		return
	}

	roleReq := new(request.Roles)

	//cek body form nya 
	if err := json.NewDecoder(r.Body).Decode(&roleReq); err != nil {
		w.WriteHeader(utils.BadRequest)
		json.NewEncoder(w).Encode(utils.Messages{
			Message: "Bad Request",
		})
		return
	}

	// validasi
	if err := utils.ValidateForm(roleReq);err != nil {
		w.WriteHeader(utils.BadRequest)
		json.NewEncoder(w).Encode(utils.Messages{
			Message: "Validation Failed",
			Errors: err.Error(),
		})

		return
	}

	if err := service.AddRoleLogic(roleReq); err != nil {
		w.WriteHeader(utils.BadRequest)
		json.NewEncoder(w).Encode(utils.Messages{
			Message: "Failed Add Role",
		})
		return
	}

	w.WriteHeader(utils.Created)
	json.NewEncoder(w).Encode(utils.Messages{
		Message: "Success Create a New Role",
	})

}