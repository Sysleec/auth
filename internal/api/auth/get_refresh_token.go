package auth

import (
	"context"

	desc "github.com/Sysleec/auth/pkg/auth_v1"
)

func (i *Implementation) GetRefreshToken(ctx context.Context, req *desc.GetRefreshTokenRequest) (*desc.GetRefreshTokenResponse, error) {
	obj, err := i.authService.GetRefreshToken(ctx, req.GetRefreshToken())
	if err != nil {
		return nil, err
	}

	return &desc.GetRefreshTokenResponse{
		RefreshToken: obj,
	}, nil
}
