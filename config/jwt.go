package config

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type CustomClaims struct {
	CodeRole string	`json:"code_role"`
	jwt.RegisteredClaims
}

var secretkey = []byte(os.Getenv("JWT_SECRET_KEY"))

func GenerateToken(user_id uint, code_role string) (string, error) {
	claims := CustomClaims{
		CodeRole: code_role,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:  uuid.New().String(),
			Subject: strconv.Itoa(int(user_id)),
			Issuer: "Scorely-Auth-Service",
			Audience:  []string{"frontend-app"},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(secretkey)

	return ss, err
}