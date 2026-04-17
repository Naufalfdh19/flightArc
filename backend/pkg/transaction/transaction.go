package transaction

import (
	"context"

	"gorm.io/gorm"
)

type TransactorRepo interface {
	WithinTransaction(ctx context.Context, fn func(txCtx context.Context) error) error
}

type TransactorRepoImpl struct {
	db *gorm.DB
}

type TxKey struct{}

func NewTransactorRepo(dbConn *gorm.DB) TransactorRepoImpl {
	return TransactorRepoImpl{
		db: dbConn,
	}
}

func (t TransactorRepoImpl) WithinTransaction(ctx context.Context, fn func(txCtx context.Context) error) error {
	return t.db.Transaction(func(tx *gorm.DB) error {
		ctxWithTx := InjectTx(ctx, tx)
		return fn(ctxWithTx)
	})
}

func InjectTx(ctx context.Context, tx *gorm.DB) context.Context {
	return context.WithValue(ctx, TxKey{}, tx)
}

func ExtractTx(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value(TxKey{}).(*gorm.DB); ok {
		return tx
	}
	return nil
}
