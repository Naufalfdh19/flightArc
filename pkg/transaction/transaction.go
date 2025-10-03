package transaction

import (
	"context"
	"database/sql"
)

type TransactorRepo interface {
	WithinTransaction(context.Context, func(context.Context) error) error
}

type TransactorRepoImpl struct {
	db *sql.DB
}

type TxKey struct{}

func NewTransactorRepo(dbConn *sql.DB) TransactorRepoImpl {
	return TransactorRepoImpl{
		db: dbConn,
	}
}

func (t *TransactorRepoImpl) WithinTransaction(ctx context.Context, fn func(txCtx context.Context) error) error {
	tx, err := t.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = fn(injectTx(ctx, tx))
	if err != nil {
		if errRollback := tx.Rollback(); errRollback != nil {
			return err
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func injectTx(ctx context.Context, tx *sql.Tx) context.Context {
	return context.WithValue(ctx, TxKey{}, tx)
}

func ExtractTx(ctx context.Context) *sql.Tx {
	if tx, ok := ctx.Value(TxKey{}).(*sql.Tx); ok {
		return tx
	}
	return nil
}