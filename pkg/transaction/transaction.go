package transaction

import (
	"context"
	"database/sql"
	"flight/pkg/apperror"
	"flight/pkg/constant"
)

func Transaction(ctx context.Context, db *sql.DB) (error) {
	tx, err := db.BeginTx(ctx, nil)
    if err != nil {
        return apperror.NewErrInternalServerError(constant.TRANSACTION, apperror.ErrInternalServerError, apperror.ErrInternalServerError)
    }

    // Defer a rollback in case anything fails.
    defer tx.Rollback()

	if err = tx.Commit(); err != nil {
        return apperror.NewErrInternalServerError(constant.TRANSACTION, apperror.ErrInternalServerError, apperror.ErrInternalServerError)
    }

	return nil
}