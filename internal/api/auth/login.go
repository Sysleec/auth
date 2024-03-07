package auth

import (
	"context"
	"fmt"

	"github.com/Sysleec/auth/internal/converter"
	desc "github.com/Sysleec/auth/pkg/auth_v1"
)

func (i *Implementation) Login(ctx context.Context, req *desc.LoginRequest) (*desc.LoginResponse, error) {
	obj, err := i.authService.Login(ctx, converter.ToUserClaimsFromLogin(req))
	if err != nil {
		fmt.Println("login err", err)
		return nil, err
	}

	return &desc.LoginResponse{
		RefreshToken: obj,
	}, nil
}
