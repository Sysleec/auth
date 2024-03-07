package transaction

import (
	"context"

	"github.com/Sysleec/auth/internal/client/db"
	"github.com/Sysleec/auth/internal/client/db/pg"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

type manager struct {
	db db.Transactor
}

// NewTransactionManager creates a new transaction manager that provides the db.TxManager interface
func NewTransactionManager(db db.Transactor) db.TxManager {
	return &manager{
		db: db,
	}
}

// transaction is the main function that executes a user-specified handler in a transaction
func (m *manager) transaction(ctx context.Context, opts pgx.TxOptions, fn db.Handler) (err error) {
	// If this is a nested transaction, we skip the initiation of a new transaction and execute the handler
	tx, ok := ctx.Value(pg.TxKey).(pgx.Tx)
	if ok {
		return fn(ctx)
	}

	// Start a new transaction
	tx, err = m.db.BeginTx(ctx, opts)
	if err != nil {
		return errors.Wrap(err, "can't begin transaction")
	}

	// Put the transaction into the context
	ctx = pg.MakeContextTx(ctx, tx)

	// We configure the deferred function to rollback or commit the transaction
	defer func() {
		// Recover after panic
		if r := recover(); r != nil {
			err = errors.Errorf("panic recovered: %v", r)
		}

		// If there was an error, rollback the transaction
		if err != nil {
			if errRollback := tx.Rollback(ctx); errRollback != nil {
				err = errors.Wrapf(err, "errRollback: %v", errRollback)
			}

			return
		}

		// If there was no error, commit the transaction
		if nil == err {
			err = tx.Commit(ctx)
			if err != nil {
				err = errors.Wrap(err, "tx commit failed")
			}
		}
	}()

	// Run the user-specified handler
	// If the handler returns an error, the transaction is rolled back,
	// Otherwise, the transaction is committed
	if err = fn(ctx); err != nil {
		err = errors.Wrap(err, "failed executing code inside transaction")
	}

	return err
}

func (m *manager) ReadCommitted(ctx context.Context, f db.Handler) error {
	txOpts := pgx.TxOptions{IsoLevel: pgx.ReadCommitted}
	return m.transaction(ctx, txOpts, f)
}
