package service

import (
	"context"

	"github.com/Sysleec/auth/internal/model"
	"google.golang.org/protobuf/types/known/emptypb"
)

// UserService is the interface that provides user methods.
type UserService interface {
	Create(ctx context.Context, usr *model.User) (int64, error)
	Get(ctx context.Context, id int64) (*model.User, error)
	Update(ctx context.Context, usr *model.User) (*emptypb.Empty, error)
	Delete(ctx context.Context, id int64) (*emptypb.Empty, error)
}
