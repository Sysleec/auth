package repository

import (
	"context"

	"github.com/Sysleec/auth/internal/model"
	"google.golang.org/protobuf/types/known/emptypb"
)

// UserRepository is an interface for user repository
type UserRepository interface {
	Create(ctx context.Context, usr *model.User) (int64, error)
	Get(ctx context.Context, id int64) (*model.User, error)
	Update(ctx context.Context, usr *model.User) (*emptypb.Empty, error)
	Delete(ctx context.Context, id int64) (*emptypb.Empty, error)
}

type AccessRepository interface {
	GetAccessibleRoles(ctx context.Context) (map[string]int32, error)
}

type AuthRepository interface {
	Login(ctx context.Context, info *model.UserClaims) (string, error)
	GetAccessToken(ctx context.Context, token string) (string, error)
	GetRefreshToken(ctx context.Context, token string) (string, error)
	GetUserRole(ctx context.Context, info *model.UserClaims) (int32, error)
}
