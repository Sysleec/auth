package auth

import (
	"context"

	"github.com/Sysleec/auth/internal/client/db"
	"github.com/Sysleec/auth/internal/model"
	"github.com/Sysleec/auth/internal/repository"
)

type repo struct {
	db db.Client
}

func NewRepo(dbClient db.Client) repository.AuthRepository {
	return &repo{db: dbClient}
}

func (r *repo) Login(ctx context.Context, info *model.UserClaims) (string, error) {
	return "", nil
}
func (r *repo) GetAccessToken(ctx context.Context, token string) (string, error) {
	return "", nil
}
func (r *repo) GetRefreshToken(ctx context.Context, token string) (string, error) {
	return "", nil
}
func (r *repo) GetUserRole(ctx context.Context) (int32, error) {
	return 1, nil // 1 = admin
}
