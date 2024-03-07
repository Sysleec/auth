package auth

import (
	"context"
	"fmt"

	"github.com/Sysleec/auth/internal/client/db"
	"github.com/Sysleec/auth/internal/model"
	"github.com/Sysleec/auth/internal/repository"

	sq "github.com/Masterminds/squirrel"
)

type repo struct {
	db db.Client
}

const (
	tableName = "users"

	idColumn    = "id"
	nameColumn  = "name"
	emailColumn = "email"
	roleColumn  = "role"
)

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
func (r *repo) GetUserRole(ctx context.Context, info *model.UserClaims) (int32, error) {
	builder := sq.Select(roleColumn).
		From(tableName).
		Where(sq.Eq{emailColumn: info.Username}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "auth_repository.GetUserRole",
		QueryRaw: query,
	}

	var role int32
	err = r.db.DB().ScanOneContext(ctx, &role, q, args...)
	if err != nil {
		return 0, err
	}

	fmt.Println("got role: ", role)

	return role, nil // 1 = admin
}
