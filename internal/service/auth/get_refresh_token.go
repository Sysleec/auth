package login

import (
	"context"
	"os"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Sysleec/auth/internal/model"
	"github.com/Sysleec/auth/internal/utils"
)

func (s *serverAuth) GetRefreshToken(ctx context.Context, token string) (string, error) {
	refreshTokenSecretKey := os.Getenv("refreshTokenSecretKey")
	claims, err := utils.VerifyToken(token, []byte(refreshTokenSecretKey))
	if err != nil {
		return "", status.Errorf(codes.Aborted, "invalid refresh token")
	}
	r, err := s.loginRepository.GetUserRole(ctx, claims)
	if err != nil {
		return "", nil
	}
	refreshToken, err := utils.GenerateToken(model.UserInfo{
		Username: claims.Username,
		Role:     r,
	},
		[]byte(refreshTokenSecretKey),
		refreshTokenExpiration,
	)
	if err != nil {
		return "", err
	}

	return refreshToken, nil
}
