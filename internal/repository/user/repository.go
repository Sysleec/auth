package user

import (
	"context"

	"github.com/Sysleec/auth/internal/repository"
	"github.com/Sysleec/auth/internal/repository/user/converter"
	"github.com/Sysleec/auth/internal/repository/user/model"
	desc "github.com/Sysleec/auth/pkg/user_v1"
	"github.com/jackc/pgx/v4/pgxpool"

	sq "github.com/Masterminds/squirrel"
)

const (
	tableName = "users"

	idColumn        = "id"
	nameColumn      = "name"
	emailColumn     = "email"
	passColumn      = "password"
	isAdminColumn   = "role"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

type repo struct {
	db *pgxpool.Pool
}

// NewRepo creates new user repository
func NewRepo(db *pgxpool.Pool) repository.UserRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, usr *desc.CreateRequest) (int64, error) {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(nameColumn, emailColumn, passColumn).
		Values(usr.Name, usr.Email, usr.Password).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	var id int64

	err = r.db.QueryRow(ctx, query, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repo) Get(ctx context.Context, id int64) (*desc.GetResponse, error) {
	builder := sq.Select(idColumn, nameColumn, emailColumn, isAdminColumn, createdAtColumn, updatedAtColumn).
		From(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: id})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var usr model.User

	err = r.db.QueryRow(ctx, query, args...).Scan(&usr.ID, &usr.Name, &usr.Email, &usr.Role, &usr.CreatedAt, &usr.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return converter.ToUserFromRepo(&usr), nil
}
