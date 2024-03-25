package access

import (
	"context"
	"errors"
	"os"

	"github.com/Sysleec/auth/internal/utils"
)

func (s *servAccess) GetName(ctx context.Context) (string, error) {
	accessTokenSecretKey := os.Getenv("accessTokenSecretKey")
	accessToken, err := accessToken(ctx)

	if err != nil {
		return "", errors.New("access token is invalid: " + err.Error())
	}

	claims, err := utils.VerifyToken(accessToken, []byte(accessTokenSecretKey))
	if err != nil {
		return "", errors.New("access token is invalid " + err.Error())
	}
	return claims.Username, nil
}
