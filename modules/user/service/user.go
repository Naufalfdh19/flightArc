package service

import (
	"context"
	"flight/modules/user/entity"
	"flight/modules/user/queryparams"
	"flight/modules/user/repo"
	"flight/pkg/apperror"
	"flight/pkg/constant"
	"flight/pkg/jwttoken"
	"flight/pkg/pagination"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetUserById(ctx context.Context, id int) (*entity.User, error)
	GetUsers(ctx context.Context, queryParams queryparams.QueryParams) (*pagination.Pagination, error)
	UpdateUserById(ctx context.Context, user entity.User) error
	DeleteUserById(ctx context.Context, id int) error
	Login(ctx context.Context, userAuth entity.User) (string, error)
	Register(ctx context.Context, user entity.User) error
	UpdatePassword(ctx context.Context, user entity.User) error 
}

type UserServiceImpl struct {
	r repo.UserRepo
}

func NewUserService(r repo.UserRepo) UserServiceImpl {
	return UserServiceImpl{
		r: r,
	}
}

func (s UserServiceImpl) GetUserById(ctx context.Context, id int) (*entity.User, error) {
	isUserExist := s.r.IsUserExists(ctx, id)
	if !isUserExist {
		return nil, apperror.NewErrStatusBadRequest(constant.GET_USER_BY_ID, apperror.ErrUserNotExists, apperror.ErrUserNotExists)
	}

	user, err := s.r.GetUserById(ctx, id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u UserServiceImpl) GetUsers(ctx context.Context, queryParams queryparams.QueryParams) (*pagination.Pagination, error) {
	totalUser, err := u.r.GetTotalUser(ctx)
	if err != nil {
		return nil, err
	}

	queryparams.CheckLimit(&queryParams)
	totalPage := totalUser / queryParams.Limit
	if totalUser%queryParams.Limit != 0 {
		totalPage += 1
	}
	queryparams.CheckPage(&queryParams, totalPage)

	users, err := u.r.GetUsers(ctx, queryParams)
	if err != nil {
		return nil, err
	}

	pagination := pagination.Pagination{
		Page:         queryParams.Page,
		TotalElement: totalUser,
		TotalPage:    totalPage,
		Data:         users,
	}

	return &pagination, nil
}

func (u UserServiceImpl) UpdateUserById(ctx context.Context, user entity.User) error {
	isUserExists := u.r.IsUserExists(ctx, user.Id)
	if !isUserExists {
		return apperror.NewErrStatusBadRequest(constant.UPDATE_USER_BY_ID, apperror.ErrUserNotExists, apperror.ErrUserNotExists)
	}

	err := u.r.UpdateUserById(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (u UserServiceImpl) DeleteUserById(ctx context.Context, id int) error {
	isUserExists := u.r.IsUserExists(ctx, id)
	if !isUserExists {
		return apperror.NewErrStatusBadRequest(constant.DELETE_USER_BY_ID, apperror.ErrUserNotExists, apperror.ErrUserNotExists)
	}

	err := u.r.DeleteUserById(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (u UserServiceImpl) Login(ctx context.Context, userAuth entity.User) (string, error) {
	isUserExists := u.r.IsEmailExists(ctx, userAuth.Email)
	if !isUserExists {
		return "", apperror.NewErrStatusBadRequest(constant.LOGIN, apperror.ErrEmailOrPasswordInvalid, apperror.ErrEmailOrPasswordInvalid)
	}

	user, err := u.r.GetUserByEmail(ctx, userAuth.Email)
	if err != nil {
		return "", err
	}
	userIdStr := strconv.Itoa(user.Id)

	if user.Role != constant.USER {
		return "", apperror.NewErrStatusBadRequest(constant.LOGIN, apperror.ErrWrongRole, apperror.ErrWrongRole)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userAuth.Password))
	if err != nil {
		return "", apperror.NewErrStatusBadRequest(constant.LOGIN, apperror.ErrEmailOrPasswordInvalid, apperror.ErrEmailOrPasswordInvalid)
	}

	jwtToken := jwttoken.JwtTokenImpl{}
	token, err := jwtToken.GenerateJwtTokenForAuth(userIdStr, user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s UserServiceImpl) Register(ctx context.Context, user entity.User) error {
	isPasswordValid := apperror.IsPasswordValid(user.Password)
	if !isPasswordValid {
		return apperror.NewErrStatusBadRequest(constant.REGISTER, apperror.ErrPasswordInvalid, apperror.ErrPasswordInvalid)
	}
	isUsernameValid := apperror.IsAlphanumeric(user.Name)
	if !isUsernameValid {
		return apperror.NewErrStatusBadRequest(constant.REGISTER, apperror.ErrUsernameInvalid, apperror.ErrUsernameInvalid)
	}

	isUserExists := s.r.IsUserExistsByEmail(ctx, user.Email)
	if isUserExists {
		return apperror.NewErrStatusBadRequest(constant.REGISTER, apperror.ErrUserExists, apperror.ErrUserExists)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return apperror.NewErrInternalServerError(constant.SERVER, apperror.ErrInternalServerError, apperror.ErrInternalServerError)
	}

	user.Password = string(hashedPassword)
	user.Role = constant.USER

	err = s.r.AddUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (u UserServiceImpl) UpdatePassword(ctx context.Context, user entity.User) error {
	isUserExists := u.r.IsUserExists(ctx, user.Id)
	if !isUserExists {
		return apperror.NewErrStatusBadRequest(constant.UPDATE_USER_BY_ID, apperror.ErrUserNotExists, apperror.ErrUserNotExists)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return apperror.NewErrInternalServerError(constant.SERVER, apperror.ErrInternalServerError, apperror.ErrInternalServerError)
	}

	user.Password = string(hashedPassword)


	err = u.r.UpdatePassword(ctx, user)
	if err != nil {
		return err
	}

	return nil
}