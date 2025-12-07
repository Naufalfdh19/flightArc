package service

import (
	"context"
	"flight/modules/user/entity"
	"flight/pkg/apperror"
	"flight/pkg/constant"
	"flight/pkg/jwttoken"
	"flight/setup/redis"
	"strconv"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (s UserServiceImpl) Login(ctx context.Context, userAuth entity.User) (string, error) {
	rdb := redis.NewRedisClient()

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

	refreshToken := uuid.New().String()
	userIdStr = strconv.Itoa(user.Id)
	rdb.Set(ctx, "refreshToken:userId:"+userIdStr, refreshToken, time.Hour*24*7)

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

func (s UserServiceImpl) GenerateNewAccessToken(ctx context.Context, userId int) (string, error) {
	userIdStr := strconv.Itoa(userId)

	rdb := redis.NewRedisClient()
	_, err := rdb.Get(ctx, "refreshToken:userId"+userIdStr).Result()
	if err != nil {
		return "", apperror.NewErrStatusNotFound(constant.GENERATE_NEW_ACCESS_TOKEN, apperror.ErrRefreshTokenNotExists, err)
	}

	user, err := s.ur.GetUserById(ctx, userId)
	if err != nil {
		return "", apperror.NewErrStatusNotFound(constant.GENERATE_NEW_ACCESS_TOKEN, apperror.ErrUserNotExists, err)
	}

	jwtToken := jwttoken.JwtTokenImpl{}
	accessToken, err := jwtToken.GenerateAccessTokenForAuth(userIdStr, user.Role)
	if err != nil {
		return "", err
	}



	return accessToken, nil
}
