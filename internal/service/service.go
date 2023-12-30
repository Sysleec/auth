package service

import (
	"context"

	"github.com/Sysleec/auth/internal/model"
)

// UserService is the interface that provides user methods.
type UserService interface {
	Create(ctx context.Context, usr *model.User) (int64, error)
	Get(ctx context.Context, id int64) (*model.User, error)
}
