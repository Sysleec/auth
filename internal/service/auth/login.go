package login

import (
	"context"
	"time"

	"github.com/pkg/errors"

	"github.com/Sysleec/auth/internal/config/env"
	"github.com/Sysleec/auth/internal/model"
	"github.com/Sysleec/auth/internal/utils"
)

const (
	refreshTokenExpiration = 60 * time.Minute
	accessTokenExpiration  = 1 * time.Minute
)

func (s *serverAuth) Login(ctx context.Context, info *model.LoginClaims) (string, error) {
	jwt, err := env.NewJWTConfig()
	if err != nil {
		return "", err
	}

	refreshTokenSecretKey := jwt.RefreshTokenSecretKey()

	r, err := s.loginRepository.Login(ctx, info)
	if err != nil {
		return "", err
	}

	refreshToken, err := utils.GenerateToken(model.UserInfo{
		Username: info.Username,
		Role:     r,
	},
		[]byte(refreshTokenSecretKey),
		refreshTokenExpiration,
	)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return refreshToken, nil
}
