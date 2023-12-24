package repository

import (
	"context"

	desc "github.com/Sysleec/auth/pkg/user_v1"
)

// UserRepository is an interface for user repository
type UserRepository interface {
	Create(ctx context.Context, usr *desc.CreateRequest) (int64, error)
	Get(ctx context.Context, id int64) (*desc.GetResponse, error)
}
