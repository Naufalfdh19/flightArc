package repo

import (
	"context"
	"database/sql"
	"flight/modules/token/entity"
	"flight/pkg/apperror"
	"flight/pkg/constant"
)

type TokenRepo interface {
	GetRefreshTokenByUserId(ctx context.Context, userID int) (*entity.Token, error)
	UpdateRefreshToken(ctx context.Context, userId int, token string) error
	AddRefreshToken(ctx context.Context, userId int, token string) error
	IsRefreshTokenExistsByUserID(ctx context.Context, userId int) bool
}

type tokenRepoImpl struct {
	db *sql.DB
}

func NewTokenRepo(db *sql.DB) tokenRepoImpl {
	return tokenRepoImpl{
		db: db,
	}
}

func (r tokenRepoImpl) GetRefreshTokenByUserId(ctx context.Context, userID int) (*entity.Token, error) {
	var token entity.Token

	query := `SELECT id, name, value
				FROM tokens
				WHERE user_id = $1 AND name = $2 AND deleted_at IS NULL`

	err := r.db.QueryRow(query, userID, constant.REFRESH_TOKEN).Scan(
		&token.Id,
		&token.Name,
		&token.Value,
	)
	if err != nil {
		return nil, apperror.NewErrInternalServerError(constant.SERVER, apperror.ErrInternalServerError, err)
	}

	return &token, nil
}

func (r tokenRepoImpl) UpdateRefreshToken(ctx context.Context, userId int, token string) error {
	query := `UPDATE tokens  
				SET name = $2, value = $3, 
					updated_at = NOW()
				WHERE user_id = $1 AND deleted_at IS NULL`

	_, err := r.db.Exec(query, userId, constant.REFRESH_TOKEN, token)
	if err != nil {
		return apperror.NewErrInternalServerError(constant.SERVER, apperror.ErrInternalServerError, err)
	}

	return nil
}

func (r tokenRepoImpl) AddRefreshToken(ctx context.Context, userId int, token string) error {
	query := `INSERT INTO tokens (name, value, user_id)
				VALUES ($1, $2, $3)`

	_, err := r.db.Exec(query, constant.REFRESH_TOKEN, token, userId)
	if err != nil {
		return apperror.NewErrInternalServerError(constant.SERVER, apperror.ErrInternalServerError, err)
	}

	return nil
}

func (r tokenRepoImpl) IsRefreshTokenExistsByUserID(ctx context.Context, userId int) bool {
	var exists bool
	query := `SELECT EXISTS(
		SELECT 1 
		FROM tokens
		WHERE user_id = $1 AND name = $2 AND deleted_at IS NULL)`
	_ = r.db.QueryRow(query, userId, constant.REFRESH_TOKEN).Scan(&exists)
	return exists
}
