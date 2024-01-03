package user

import (
	"context"

	"github.com/Sysleec/auth/internal/client/db"
	"github.com/Sysleec/auth/internal/model"
	"github.com/Sysleec/auth/internal/repository"
	"github.com/Sysleec/auth/internal/repository/user/converter"
	modelRepo "github.com/Sysleec/auth/internal/repository/user/model"
	"google.golang.org/protobuf/types/known/emptypb"

	sq "github.com/Masterminds/squirrel"
)

const (
	tableName = "users"

	idColumn        = "id"
	nameColumn      = "name"
	emailColumn     = "email"
	passColumn      = "password"
	roleColumn      = "role"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

type repo struct {
	db db.Client
}

// NewRepo creates new user repository
func NewRepo(db db.Client) repository.UserRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, usr *model.User) (int64, error) {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(nameColumn, emailColumn, passColumn).
		Values(usr.Name, usr.Email, usr.Password).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "user_repositoy.Create",
		QueryRaw: query,
	}

	var id int64

	err = r.db.DB().ScanOneContext(ctx, &id, q, args...)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repo) Get(ctx context.Context, id int64) (*model.User, error) {
	builder := sq.Select(idColumn, nameColumn, emailColumn, roleColumn, createdAtColumn, updatedAtColumn).
		From(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: id})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "user_repositoy.Get",
		QueryRaw: query,
	}

	var usr modelRepo.User
	err = r.db.DB().ScanOneContext(ctx, &usr, q, args...)
	if err != nil {
		return nil, err
	}

	return converter.ToUserFromRepo(&usr), nil
}

func (r *repo) Update(ctx context.Context, usr *model.User) (*emptypb.Empty, error) {
	var builder sq.UpdateBuilder

	if usr.Name == "" && usr.Email == "" {
		return &emptypb.Empty{}, nil

	} else if usr.Name == "" {
		builder = sq.Update(tableName).
			PlaceholderFormat(sq.Dollar).
			Set(emailColumn, usr.Email).
			Where(sq.Eq{idColumn: usr.ID})

	} else if usr.Email == "" {
		builder = sq.Update(tableName).
			PlaceholderFormat(sq.Dollar).
			Set(nameColumn, usr.Name).
			Where(sq.Eq{idColumn: usr.ID})

	} else {
		builder = sq.Update(tableName).
			PlaceholderFormat(sq.Dollar).
			Set(nameColumn, usr.Name).
			Set(emailColumn, usr.Email).
			Where(sq.Eq{idColumn: usr.ID})
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return &emptypb.Empty{}, err
	}

	q := db.Query{
		Name:     "user_repositoy.Update",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}

func (r *repo) Delete(ctx context.Context, id int64) (*emptypb.Empty, error) {
	builder := sq.Delete(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: id})

	query, args, err := builder.ToSql()
	if err != nil {
		return &emptypb.Empty{}, err
	}

	q := db.Query{
		Name:     "user_repositoy.Delete",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}
