package helper

import (
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type CustomClaims struct {
	CodeRole string `json:"code_role"`
	jwt.RegisteredClaims
}

func GenerateToken(secretKey string,user_id uint, code_role string) (string, error) {
	claims := CustomClaims{
		CodeRole: code_role,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        uuid.New().String(),
			Subject:   strconv.Itoa(int(user_id)),
			Issuer:    "Scorely-Auth-Service",
			Audience:  []string{"frontend-app"},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(secretKey))

	return ss, err
}

func ParseTokenJwt(secretKey, tokenJwt string) (*CustomClaims, error) {
	secret := []byte(secretKey)

	token, err := jwt.ParseWithClaims(tokenJwt, &CustomClaims{}, func(t *jwt.Token) (any, error) {
		return secret, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			claims, ok := token.Claims.(*CustomClaims)
			if ok {
				return claims, err
			}
		}
		return nil, err
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return nil, jwt.ErrTokenExpired
	}

	return claims, nil
}
