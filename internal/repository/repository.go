package repository

import (
	"context"

	"github.com/Sysleec/auth/internal/model"
)

// UserRepository is an interface for user repository
type UserRepository interface {
	Create(ctx context.Context, usr *model.User) (int64, error)
	Get(ctx context.Context, id int64) (*model.User, error)
}
