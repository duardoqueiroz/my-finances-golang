package middlewares

import (
	"fmt"

	"github.com/duardoqueiroz/my-finances-golang/pkg/infra/security"
)

func parseToken(token string) (string, string, error) {
	parsedToken, err := security.ParseAccessToken(token)
	if err != nil {
		return "", "", fmt.Errorf("error parsing token: %w", err)
	}
	return parsedToken.Id, parsedToken.Role, nil
}
