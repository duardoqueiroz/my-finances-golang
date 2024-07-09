package security

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type UserClaims struct {
	Id   string `json:"id"`
	Role string `json:"role"`
	jwt.StandardClaims
}

func NewAccessToken(claims UserClaims) (string, int64, error) {
	expiresAt := time.Now().Add(24 * time.Hour).Unix()

	claims.StandardClaims.ExpiresAt = expiresAt

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS384, claims)
	signed, err := accessToken.SignedString([]byte(os.Getenv("JWT_SECRET")))

	return signed, expiresAt, err
}

func ParseAccessToken(token string) (*UserClaims, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}

	return parsedToken.Claims.(*UserClaims), nil
}
