package repo

import (
	"context"
	"flight/modules/user/entity"
	"flight/modules/user/queryparams"
	"flight/pkg/apperror"
	"flight/pkg/constant"

	"github.com/jackc/pgx/v5"
)

type AdminRepo interface {
	GetUsers(ctx context.Context, queryParams queryparams.QueryParams) ([]entity.User, error)
}

type AdminRepoImpl struct {
	db *pgx.Conn
}

func NewAdminRepo(db *pgx.Conn) AdminRepoImpl {
	return AdminRepoImpl{
		db: db,
	}
}

func (r AdminRepoImpl) GetUsers(ctx context.Context, queryParams queryparams.QueryParams) ([]entity.User, error) {
	var users []entity.User

	query := `SELECT id, name, email, password, phone_number, role
				FROM users
				WHERE deleted_at IS NULL`
	query += queryparams.AddPagination(queryParams)

	rows, err := r.db.Query(ctx, query)

	if err != nil {
		return nil, apperror.NewErrInternalServerError(constant.SERVER, apperror.ErrInternalServerError, err)
	}
	defer rows.Close()

	for rows.Next() {
		var user entity.User

		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.PhoneNumber, &user.Role)
		if err != nil {
			return nil, apperror.NewErrInternalServerError(constant.SERVER, apperror.ErrInternalServerError, err)
		}
		users = append(users, user)
	}

	return users, nil
}
