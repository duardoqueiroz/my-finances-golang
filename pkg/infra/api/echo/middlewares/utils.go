package middlewares

import (
	"errors"
	"fmt"
	"strings"

	"github.com/duardoqueiroz/my-finances-golang/pkg/infra/security"
)

func parseToken(token string) (string, string, error) {
	parsedToken, err := security.ParseAccessToken(token)
	if err != nil {
		return "", "", fmt.Errorf("error parsing token: %w", err)
	}
	return parsedToken.Id, parsedToken.Role, nil
}

func getToken(authHeader string) (string, error) {
	if authHeader == "" {
		return "", errors.New("invalid token")
	}
	splitedHeader := strings.Split(authHeader, "Bearer")
	if splitedHeader[1] == "" {
		return "", errors.New("invalid token")
	}

	token := strings.TrimSpace(splitedHeader[1])

	return token, nil
}
