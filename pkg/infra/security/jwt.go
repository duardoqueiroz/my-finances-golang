package security

import (
	"os"

	"github.com/golang-jwt/jwt"
)

type UserClaims struct {
	Id string `json:"id"`
	Role string `json:"role"`
	jwt.StandardClaims
}

func NewAccessToken(claims UserClaims) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS384, claims)

	return accessToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func ParseAccessToken(token string) (*UserClaims, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil,err
	}

	return parsedToken.Claims.(*UserClaims), nil
}