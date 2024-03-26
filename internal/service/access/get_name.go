package access

import (
	"context"
	"errors"

	"github.com/Sysleec/auth/internal/config/env"
	"github.com/Sysleec/auth/internal/utils"
)

func (s *servAccess) GetName(ctx context.Context) (string, error) {
	jwt, err := env.NewJWTConfig()
	if err != nil {
		return "", err
	}
	accessTokenSecretKey := jwt.AccessTokenSecretKey()
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
