package service

import (
	"context"
	"flight/modules/user/entity"
	"flight/modules/user/queryparams"
	"flight/modules/user/repo"
	tokenRepo "flight/modules/token/repo"
	"flight/pkg/apperror"
	"flight/pkg/constant"
	"flight/pkg/jwttoken"
	"flight/pkg/pagination"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetUserById(ctx context.Context, idAuth, idParam int) (*entity.User, error)
	GetUsers(ctx context.Context, queryParams queryparams.QueryParams) (*pagination.Pagination, error)
	UpdateUserById(ctx context.Context, user entity.User) error
	DeleteUserById(ctx context.Context, id int) error
	Login(ctx context.Context, userAuth entity.User) (string, error)
	Register(ctx context.Context, user entity.User) error
	UpdatePassword(ctx context.Context, user entity.User) error
}

type UserServiceImpl struct {
	ur repo.UserRepo
	tr tokenRepo.TokenRepo
}

func NewUserService(ur repo.UserRepo, tr tokenRepo.TokenRepo) UserServiceImpl {
	return UserServiceImpl{
		ur: ur,
		tr: tr,
	}
}

func (s UserServiceImpl) GetUserById(ctx context.Context, idAuth, idParam int) (*entity.User, error) {
	isUserExist := s.ur.IsUserExists(ctx, idParam)
	if !isUserExist {
		return nil, apperror.NewErrStatusBadRequest(constant.GET_USER_BY_ID, apperror.ErrUserNotExists, apperror.ErrUserNotExists)
	}

	userAuth, err := s.ur.GetUserById(ctx, idAuth)
	if err != nil {
		return nil, err
	}
	if userAuth.Id != idParam {
		return nil, apperror.NewErrStatusUnauthorized(constant.GET_USER_BY_ID, apperror.ErrUnauthorized, apperror.ErrUnauthorized)
	}

	user, err := s.ur.GetUserById(ctx, idParam)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s UserServiceImpl) GetUsers(ctx context.Context, queryParams queryparams.QueryParams) (*pagination.Pagination, error) {
	totalUser, err := s.ur.GetTotalUser(ctx)
	if err != nil {
		return nil, err
	}

	queryparams.CheckLimit(&queryParams)
	totalPage := totalUser / queryParams.Limit
	if totalUser%queryParams.Limit != 0 {
		totalPage += 1
	}
	queryparams.CheckPage(&queryParams, totalPage)

	users, err := s.ur.GetUsers(ctx, queryParams)
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

func (s UserServiceImpl) UpdateUserById(ctx context.Context, user entity.User) error {
	isUserExists := s.ur.IsUserExists(ctx, user.Id)
	if !isUserExists {
		return apperror.NewErrStatusBadRequest(constant.UPDATE_USER_BY_ID, apperror.ErrUserNotExists, apperror.ErrUserNotExists)
	}

	err := s.ur.UpdateUserById(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (s UserServiceImpl) DeleteUserById(ctx context.Context, id int) error {
	isUserExists := s.ur.IsUserExists(ctx, id)
	if !isUserExists {
		return apperror.NewErrStatusBadRequest(constant.DELETE_USER_BY_ID, apperror.ErrUserNotExists, apperror.ErrUserNotExists)
	}

	err := s.ur.DeleteUserById(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (s UserServiceImpl) Login(ctx context.Context, userAuth entity.User) (string, error) {
	isUserExists := s.ur.IsEmailExists(ctx, userAuth.Email)
	if !isUserExists {
		return "", apperror.NewErrStatusBadRequest(constant.LOGIN, apperror.ErrEmailOrPasswordInvalid, apperror.ErrEmailOrPasswordInvalid)
	}

	user, err := s.ur.GetUserByEmail(ctx, userAuth.Email)
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
	accessToken, err := jwtToken.GenerateAccessTokenForAuth(userIdStr, user.Role)
	if err != nil {
		return "", err
	}

	refreshToken, err := jwtToken.GenerateRefreshToken(userIdStr, user.Role)
	if err != nil {
		return "", err
	}

	isRefreshTokenExists := s.tr.IsRefreshTokenExistsByUserID(ctx, user.Id)
	if !isRefreshTokenExists {
		err = s.tr.AddRefreshToken(ctx, user.Id, refreshToken)
		if err != nil {
			return "", err
		}
	} else {
		err = s.tr.UpdateRefreshToken(ctx, user.Id, refreshToken)
		if err != nil {
			return "", err
		}
	}

	return accessToken, nil
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

	isUserExists := s.ur.IsUserExistsByEmail(ctx, user.Email)
	if isUserExists {
		return apperror.NewErrStatusBadRequest(constant.REGISTER, apperror.ErrUserExists, apperror.ErrUserExists)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return apperror.NewErrInternalServerError(constant.SERVER, apperror.ErrInternalServerError, apperror.ErrInternalServerError)
	}

	user.Password = string(hashedPassword)
	user.Role = constant.USER

	err = s.ur.AddUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (s UserServiceImpl) UpdatePassword(ctx context.Context, user entity.User) error {
	isUserExists := s.ur.IsUserExists(ctx, user.Id)
	if !isUserExists {
		return apperror.NewErrStatusBadRequest(constant.UPDATE_USER_BY_ID, apperror.ErrUserNotExists, apperror.ErrUserNotExists)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return apperror.NewErrInternalServerError(constant.SERVER, apperror.ErrInternalServerError, apperror.ErrInternalServerError)
	}

	user.Password = string(hashedPassword)

	err = s.ur.UpdatePassword(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
