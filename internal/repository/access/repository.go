package access

import (
	"context"

	"github.com/Sysleec/auth/internal/client/db"
	"github.com/Sysleec/auth/internal/model"
	"github.com/Sysleec/auth/internal/repository"
)

type repo struct {
	db db.Client
}

func NewRepo(db db.Client) repository.AccessRepository {
	return &repo{db: db}
}

func (r *repo) GetAccessibleRoles(ctx context.Context) (map[string]int32, error) {
	accessibleRoles := make(map[string]int32)
	accessibleRoles[model.ExamplePath] = 1 // 1 = admin
	return accessibleRoles, nil
}
