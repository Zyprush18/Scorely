package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/Zyprush18/Scorely/helper"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

func MiddlewareAuth(rdb *redis.Client, next http.Handler, roles ...string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		auth := strings.Split(r.Header.Get("Authorization"), " ")

		if len(auth) != 2 || strings.TrimSpace(auth[1]) == "" {
			helper.ReturnResponse(w, helper.Messages{
				Http_code: helper.Unauthorized,
				Message:   "Token Is Missing",
				Errors:    "Unauthorized",
			})
			return
		}

		token, err := helper.ParseTokenJwt(os.Getenv("JWT_SECRET_KEY"), strings.TrimSpace(auth[1]))
		if err != nil {
			if errors.Is(err, jwt.ErrTokenExpired) && token != nil {
				subject := token.Subject

				refreshToken, rerr := rdb.Get(r.Context(), "refresh_token:"+subject).Result()
				if rerr != nil {
					helper.ReturnResponse(w, helper.Messages{
						Http_code: helper.Unauthorized,
						Message:   "Token Is Expired",
						Errors:    "Unauthorized",
					})
					return
				}

				refreshClaims, rerr := helper.ParseTokenJwt(os.Getenv("REFRESH_SECRET_KEY"), refreshToken)
				if rerr != nil {
					helper.ReturnResponse(w, helper.Messages{
						Http_code: helper.Unauthorized,
						Message:   "Token Is Expired",
						Errors:    "Unauthorized",
					})
					return
				}

				userID, _ := strconv.Atoi(refreshClaims.Subject)
				newToken, terr := helper.GenerateToken(os.Getenv("JWT_SECRET_KEY"), uint(userID), refreshClaims.CodeRole)
				if terr != nil {
					helper.ReturnResponse(w, helper.Messages{
						Http_code: helper.InternalServError,
						Message:   "Something Went Wrong",
						Errors:    "Internal Server Error",
					})
					return
				}

				w.Header().Set("X-New-Token", newToken)

				idteacher, aerr := strconv.Atoi(refreshClaims.Subject)
				if aerr != nil {
					helper.ReturnResponse(w, helper.Messages{
						Http_code: helper.Unauthorized,
						Message:   "Invalid Token Subject",
						Errors:    "Token subject must be numeric",
					})
					return
				}

				ctx := context.WithValue(r.Context(), helper.KeyUserID, idteacher)
				ctx = context.WithValue(ctx, helper.KeyCodeRole, refreshClaims.CodeRole)
				ctx = context.WithValue(ctx, helper.KeyTokenID, refreshClaims.ID)

				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}

			helper.ReturnResponse(w, helper.Messages{
				Http_code: helper.Unauthorized,
				Message:   "Invalid Auth Token",
				Errors:    "Unauthorization",
			})
			return
		}

		blacklisted, berr := rdb.Exists(r.Context(), "blacklist_token:"+token.ID).Result()
		if berr == nil && blacklisted > 0 {
			w.WriteHeader(helper.Unauthorized)
			json.NewEncoder(w).Encode(helper.Messages{
				Message: "Token Is Expired",
				Errors:  "Unauthorized",
			})
			return
		}

		if !checkRole(strings.ToLower(token.CodeRole), roles...) {
			w.WriteHeader(helper.Forbidden)
			json.NewEncoder(w).Encode(helper.Messages{
				Message: "Your role does not have access to this endpoint.",
				Errors:  "Forbidden",
			})
			return
		}

		idteacher, err := strconv.Atoi(token.Subject)
		if err != nil {
			helper.ReturnResponse(w, helper.Messages{
				Http_code: helper.Unauthorized,
				Message:   "Invalid Token Subject",
				Errors:    "Token subject must be numeric",
			})
			return
		}

		ctx := context.WithValue(r.Context(), helper.KeyUserID, idteacher)
		ctx = context.WithValue(ctx, helper.KeyCodeRole, token.CodeRole)
		ctx = context.WithValue(ctx, helper.KeyTokenID, token.ID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func checkRole(coderole string, role ...string) bool {
	for _, v := range role {
		if coderole != "" && coderole == v {
			return true
		}
	}

	return false
}
