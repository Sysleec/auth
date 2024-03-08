package login

import (
	"context"

	"github.com/pkg/errors"

	"github.com/Sysleec/auth/internal/config/env"
	"github.com/Sysleec/auth/internal/model"
	"github.com/Sysleec/auth/internal/utils"
)

func (s *serverAuth) Login(ctx context.Context, info *model.LoginClaims) (string, error) {
	jwt, err := env.NewJWTConfig()
	if err != nil {
		return "", err
	}

	refreshTokenSecretKey := jwt.RefreshTokenSecretKey()
	refreshTokenExpiration := jwt.RefreshTokenExpiration()

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
