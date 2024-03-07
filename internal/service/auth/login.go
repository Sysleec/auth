package login

import (
	"context"
	"os"
	"time"

	"github.com/pkg/errors"

	"github.com/Sysleec/auth/internal/model"
	"github.com/Sysleec/auth/internal/utils"
)

const (
	refreshTokenExpiration = 60 * time.Minute
	accessTokenExpiration  = 1 * time.Minute
)

func (s *serverAuth) Login(ctx context.Context, info *model.UserClaims) (string, error) {
	refreshTokenSecretKey := os.Getenv("refreshTokenSecretKey")
	r, err := s.loginRepository.GetUserRole(ctx, info)
	if err != nil {
		return "", nil
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
