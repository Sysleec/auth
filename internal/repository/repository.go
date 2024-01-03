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
