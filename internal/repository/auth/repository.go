package auth

import (
	"context"
	"errors"

	"github.com/Sysleec/auth/internal/client/db"
	"github.com/Sysleec/auth/internal/model"
	"github.com/Sysleec/auth/internal/repository"
	"github.com/Sysleec/auth/internal/utils"

	sq "github.com/Masterminds/squirrel"
)

type repo struct {
	db db.Client
}

const (
	tableName = "users"

	idColumn   = "id"
	nameColumn = "name"
	roleColumn = "role"
	passColumn = "password"
)

func NewRepo(dbClient db.Client) repository.AuthRepository {
	return &repo{db: dbClient}
}

func (r *repo) Login(ctx context.Context, info *model.LoginClaims) (int32, error) {
	builder := sq.Select(roleColumn, passColumn).
		From(tableName).
		Where(sq.Eq{nameColumn: info.Username}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "auth_repository.Login",
		QueryRaw: query,
	}

	var usr model.User
	err = r.db.DB().ScanOneContext(ctx, &usr, q, args...)
	if err != nil {
		return 0, err
	}

	if !utils.VerifyPassword(usr.Password, info.Password) {
		return 0, errors.New("invalid password")
	}

	return int32(usr.Role), nil
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
		Where(sq.Eq{nameColumn: info.Username}).
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

	return role, nil

}
