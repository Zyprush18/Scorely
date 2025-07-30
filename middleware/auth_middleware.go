package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/Zyprush18/Scorely/config"
	"github.com/Zyprush18/Scorely/helper"
	"github.com/golang-jwt/jwt/v5"
)

func MiddlewareAuth(next http.Handler, roles ...string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		auth := strings.Split(r.Header.Get("Authorization"), " ")

		// cek apakah di authorization nya ada token atau nggak
		if len(auth) != 2 || strings.TrimSpace(auth[1]) == "" {
			w.WriteHeader(helper.Unauthorized)
			json.NewEncoder(w).Encode(helper.Messages{
				Message: "Token Is Missing",
				Errors:  "Unauthorized",
			})
			return
		}

		// mengecek apakah tokennya valid atau nggak
		token, err := config.ParseTokenJwt(strings.TrimSpace(auth[1]))
		if err != nil {
			// mengecek apakah token expired atau nggak
			if errors.Is(err, jwt.ErrTokenExpired) {
				w.WriteHeader(helper.Unauthorized)
				json.NewEncoder(w).Encode(helper.Messages{
					Message: "Token Is Expired",
					Errors:  "Unauthorized",
				})
				return
			}

			// mengecek apakah token auth nya benar atau nggak
			w.WriteHeader(helper.Unauthorized)
			json.NewEncoder(w).Encode(helper.Messages{
				Message: "Invalid Auth Token",
				Errors:  "Unauthorization",
			})
			return
		}

		// mengecek role yang login
		if !checkRole(token.CodeRole,roles...) {
				w.WriteHeader(helper.Forbidden)
				json.NewEncoder(w).Encode(helper.Messages{
					Message: "Your role does not have access to this endpoint.",
					Errors:  "Forbidden",
				})
				return
			}

		idteacher, err := strconv.Atoi(token.Subject)
		if err != nil {
			w.WriteHeader(helper.Unauthorized)
			json.NewEncoder(w).Encode(helper.Messages{
				Message: "Invalid Token Subject",
				Errors:  "Token subject must be numeric",
			})
			return
		}

		// kirim dalam bentuk context
		ctx := context.WithValue(r.Context(), helper.KeyUserID, idteacher)
		ctx = context.WithValue(ctx, helper.KeyCodeRole, token.CodeRole)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}


func checkRole(coderole string, role ...string) bool  {
	for _, v := range role {
		if coderole != "" && coderole == v {
			return true
		}
	}

	return false
}