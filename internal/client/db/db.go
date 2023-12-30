package db

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

// Handler - handler for transaction
type Handler func(ctx context.Context) error

// Client for working with DB
type Client interface {
	DB() DB
	Close() error
}

// TxManager a transaction manager that executes a user-specified handler in a transaction
type TxManager interface {
	ReadCommitted(ctx context.Context, f Handler) error
}

// Query wrapper for query having name and raw query
// Name can be used for logging, for example
type Query struct {
	Name     string
	QueryRaw string
}

// Transactor interface for working with transactions
type Transactor interface {
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
}

// SQLExecer combines NamedExecer and QueryExecer
type SQLExecer interface {
	NamedExecer
	QueryExecer
}

// NamedExecer interface for working with named queries
type NamedExecer interface {
	ScanOneContext(ctx context.Context, dest interface{}, q Query, args ...interface{}) error
	ScanAllContext(ctx context.Context, dest interface{}, q Query, args ...interface{}) error
}

// QueryExecer interface for working with usual queries
type QueryExecer interface {
	ExecContext(ctx context.Context, q Query, args ...interface{}) (pgconn.CommandTag, error)
	QueryContext(ctx context.Context, q Query, args ...interface{}) (pgx.Rows, error)
	QueryRowContext(ctx context.Context, q Query, args ...interface{}) pgx.Row
}

// Pinger interface for checking connection to DB
type Pinger interface {
	Ping(ctx context.Context) error
}

// DB interface for working with DB
type DB interface {
	SQLExecer
	Transactor
	Pinger
	Close()
}
