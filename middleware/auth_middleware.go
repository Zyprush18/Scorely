package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/Zyprush18/Scorely/helper"
)

func MiddlewareAuthAdmin(next http.Handler) http.Handler  {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type","application/json")
		auth := r.Header.Get("Authorization")
		if auth == "" {
			w.WriteHeader(helper.Unauthorized)
			json.NewEncoder(w).Encode(helper.Messages{
				Message: "Token Is Missing",
				Errors: "Unauthorized",
			})
			return 
		}

		

	})
}