package repo

import (
	"context"
	"database/sql"
	"flight/modules/user/entity"
	"flight/modules/user/queryparams"
	"flight/pkg/apperror"
	"flight/pkg/constant"
)

type UserRepo interface {
	GetUserById(ctx context.Context, id int) (*entity.User, error)
	IsUserExists(ctx context.Context, id int) bool
	GetUsers(ctx context.Context, queryParams queryparams.QueryParams) ([]entity.User, error)
	GetTotalUser(ctx context.Context) (int, error)
	UpdateUserById(ctx context.Context, user entity.User) error
	DeleteUserById(ctx context.Context, id int) error
	IsEmailExists(ctx context.Context, email string) bool
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	AddUser(ctx context.Context, user entity.User) error
	IsUserExistsByEmail(ctx context.Context, email string) bool
	UpdatePassword(ctx context.Context, user entity.User) error
}

type userRepoImpl struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) userRepoImpl {
	return userRepoImpl{
		db: db,
	}
}

func (r userRepoImpl) GetUsers(ctx context.Context, queryParams queryparams.QueryParams) ([]entity.User, error) {
	var users []entity.User

	query := `SELECT id, name, email, phone_number, role
				FROM users
				WHERE deleted_at IS NULL`
	query += queryparams.AddPagination(queryParams)

	rows, err := r.db.Query(query)

	if err != nil {
		return nil, apperror.NewErrInternalServerError(constant.SERVER, apperror.ErrInternalServerError, err)
	}
	defer rows.Close()

	for rows.Next() {
		var user entity.User

		err := rows.Scan(
			&user.Id,
			&user.Name,
			&user.Email,
			&user.PhoneNumber,
			&user.Role)
		if err != nil {
			return nil, apperror.NewErrInternalServerError(constant.SERVER, apperror.ErrInternalServerError, err)
		}
		users = append(users, user)
	}

	return users, nil
}

func (r userRepoImpl) GetTotalUser(ctx context.Context) (int, error) {
	var totalUser int
	query := `SELECT COUNT(*) 
				FROM users
				WHERE deleted_at IS NULL`

	err := r.db.QueryRow(query).Scan(&totalUser)
	if err != nil {
		return 0, apperror.NewErrInternalServerError(constant.SERVER, apperror.ErrInternalServerError, err)
	}

	return totalUser, nil
}

func (r userRepoImpl) GetUserById(ctx context.Context, id int) (*entity.User, error) {
	var user entity.User

	query := `SELECT id, name, email, phone_number, role
			FROM users 
			WHERE id = $1`

	err := r.db.QueryRow(query, id).Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.PhoneNumber,
		&user.Role,
	)
	if err != nil {
		return nil, apperror.NewErrInternalServerError(constant.SERVER, apperror.ErrInternalServerError, err)
	}

	return &user, nil
}

func (r userRepoImpl) IsUserExists(ctx context.Context, id int) bool {
	var exists bool
	query := `SELECT EXISTS(
		SELECT 1 
		FROM users 
		WHERE id = $1 AND deleted_at IS NULL)`
	_ = r.db.QueryRow(query, id).Scan(&exists)
	return exists
}

func (r userRepoImpl) UpdateUserById(ctx context.Context, user entity.User) error {
	query := `UPDATE users  
				SET name = $2, email = $3, phone_number = $4, 
					updated_at = NOW()
				WHERE id = $1 AND deleted_at IS NULL`

	_, err := r.db.Exec(query,
		user.Id,
		user.Name,
		user.Email,
		user.PhoneNumber)
	if err != nil {
		return apperror.NewErrInternalServerError(constant.SERVER, apperror.ErrInternalServerError, err)
	}

	return nil
}

func (r userRepoImpl) DeleteUserById(ctx context.Context, id int) error {
	query := `UPDATE flight.users 
				SET deleted_at = NOW(), updated_at = NOW()
				WHERE id = $1 AND deleted_at IS NULL`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r userRepoImpl) IsEmailExists(ctx context.Context, email string) bool {
	var exists bool
	query := `SELECT EXISTS(
				SELECT 1 
				FROM users 
				WHERE email = $1 AND deleted_at IS NULL)`
	_ = r.db.QueryRow(query, email).Scan(&exists)
	return exists
}

func (r userRepoImpl) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User

	query := `SELECT id, name, email, password, phone_number, role
			FROM users 
			WHERE email = $1`

	err := r.db.QueryRow(query, email).Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.PhoneNumber,
		&user.Role,
	)
	if err != nil {
		return nil, apperror.NewErrInternalServerError(constant.SERVER, apperror.ErrInternalServerError, err)
	}

	return &user, nil
}

func (r userRepoImpl) AddUser(ctx context.Context, user entity.User) error {
	query := `INSERT INTO users (name, email, password, phone_number, role)
				VALUES ($1, $2, $3, $4, $5)`

	_, err := r.db.Exec(query, user.Name, user.Email, user.Password, user.PhoneNumber, user.Role)
	if err != nil {
		return apperror.NewErrInternalServerError(constant.SERVER, apperror.ErrInternalServerError, err)
	}

	return nil
}

func (r userRepoImpl) IsUserExistsByEmail(ctx context.Context, email string) bool {
	var exists bool
	query := `SELECT EXISTS(
		SELECT 1 
		FROM users 
		WHERE email = $1 AND deleted_at IS NULL)`
	_ = r.db.QueryRow(query, email).Scan(&exists)
	return exists
}

func (r userRepoImpl) UpdatePassword(ctx context.Context, user entity.User) error {
	query := `UPDATE users  
				SET password = $2, 
					updated_at = NOW()
				WHERE id = $1 AND deleted_at IS NULL`

	_, err := r.db.Exec(query,
		user.Id,
		user.Password)
	if err != nil {
		return apperror.NewErrInternalServerError(constant.SERVER, apperror.ErrInternalServerError, err)
	}

	return nil
}

