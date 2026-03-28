package repo

import (
	"context"
	"flight/modules/user/entity"
	"flight/modules/user/queryparams"
	"flight/pkg/apperror"
	"flight/pkg/common"
	"flight/pkg/constant"

	"gorm.io/gorm"
)

type UserRepo interface {
	GetUserById(ctx context.Context, id int) (*entity.User, error)
	IsUserExists(ctx context.Context, id int) bool
	GetUsers(ctx context.Context, queryParams queryparams.QueryParams) ([]entity.User, int, error)
	UpdateUserById(ctx context.Context, user entity.User) error
	DeleteUserById(ctx context.Context, id int) error
	IsEmailExists(ctx context.Context, email string) bool
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	AddUser(ctx context.Context, user entity.User) error
	IsUserExistsByEmail(ctx context.Context, email string) bool
	UpdatePassword(ctx context.Context, user entity.User) error
}

type userRepoImpl struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) userRepoImpl {
	return userRepoImpl{
		db: db,
	}
}

func (r userRepoImpl) GetUsers(ctx context.Context, queryParams queryparams.QueryParams) ([]entity.User, int, error) {
	var users []entity.User
    var total int64

    query := r.db.WithContext(ctx).Model(&entity.User{}).Where("deleted_at IS NULL")

    if err := query.Count(&total).Error; err != nil {
        return nil, 0, apperror.NewErrInternalServerError(constant.SERVER, apperror.ErrInternalServerError, err)
    }

    err := query.
        Select("id", "name", "email", "phone_number", "role").
        Scopes(common.Paginate(queryParams.Page, queryParams.Limit, int(total))).
        Find(&users).Error
    if err != nil {
        return nil, 0, apperror.NewErrInternalServerError(constant.SERVER, apperror.ErrInternalServerError, err)
    }

    return users, int(total), nil
}


func (r userRepoImpl) GetUserById(ctx context.Context, id int) (*entity.User, error) {
	var user entity.User

	err := r.db.
		Select("id", "name", "email", "phone_number", "role").
		Where(&entity.User{Id: id}).
		First(&user).
		Scan(&user).Error
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
		WHERE id = ? AND deleted_at IS NULL)`
	_ = r.db.Raw(query, id).Scan(&exists)
	return exists
}

func (r userRepoImpl) UpdateUserById(ctx context.Context, user entity.User) error {
	err := r.db.WithContext(ctx).Model(&entity.User{}).
        Where("id = ? AND deleted_at IS NULL", user.Id).
        Updates(entity.User{
            Name:        user.Name,
            Email:       user.Email,
            PhoneNumber: user.PhoneNumber,
        }).Error

    if err != nil {
        return apperror.NewErrInternalServerError(constant.SERVER, apperror.ErrInternalServerError, err)
    }

	return nil
}

func (r userRepoImpl) DeleteUserById(ctx context.Context, id int) error {
	err := r.db.Delete(&entity.User{}, "id = ?", id).Error
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
				WHERE email = ? AND deleted_at IS NULL)`
	_ = r.db.Raw(query, email).Scan(&exists)
	return exists
}

func (r userRepoImpl) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User

	err := r.db.
		Select("id", "name", "email", "password", "phone_number", "role").
		Where(&entity.User{Email: email}).
		First(&user).
		Scan(&user).Error
	if err != nil {
		return nil, apperror.NewErrInternalServerError(constant.SERVER, apperror.ErrInternalServerError, err)
	}

	return &user, nil
}

func (r userRepoImpl) AddUser(ctx context.Context, user entity.User) error {
	query := `INSERT INTO users (name, email, password, phone_number, role)
				VALUES (?, ?, ?, ?, ?)`

	err := r.db.Exec(query, user.Name, user.Email, user.Password, user.PhoneNumber, user.Role)
	if err != nil {
		return apperror.NewErrInternalServerError(constant.SERVER, apperror.ErrInternalServerError, err.Error)
	}

	return nil
}

func (r userRepoImpl) IsUserExistsByEmail(ctx context.Context, email string) bool {
	var exists bool
	query := `SELECT EXISTS(
		SELECT 1 
		FROM users 
		WHERE email = ? AND deleted_at IS NULL)`
	_ = r.db.Raw(query, email).Scan(&exists)
	return exists
}

func (r userRepoImpl) UpdatePassword(ctx context.Context, user entity.User) error {
	err := r.db.WithContext(ctx).Model(&entity.User{}).
        Where("id = ? AND deleted_at IS NULL", user.Id).
        Updates(entity.User{
            Password: user.Password,
        }).Error

    if err != nil {
        return apperror.NewErrInternalServerError(constant.SERVER, apperror.ErrInternalServerError, err)
    }

	return nil
}
