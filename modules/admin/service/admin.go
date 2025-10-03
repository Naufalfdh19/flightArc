package service

import (
	"context"
	"flight/modules/admin/repo"
	"flight/modules/user/entity"
	"flight/modules/user/queryparams"
	userRepo "flight/modules/user/repo"
	"flight/pkg/apperror"
	"flight/pkg/constant"
	"flight/pkg/jwttoken"
	"flight/pkg/pagination"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

type AdminService interface {
	GetUsers(ctx context.Context, queryParams queryparams.QueryParams) (*pagination.Pagination, error)
}

type AdminServiceImpl struct {
	ar repo.AdminRepo
	ur userRepo.UserRepo
}

func NewUserService(adminRepo repo.AdminRepo, userRepo userRepo.UserRepo) AdminServiceImpl {
	return AdminServiceImpl{
		ar: adminRepo,
		ur: userRepo,
	}
}

func (u AdminServiceImpl) GetUsers(ctx context.Context, queryParams queryparams.QueryParams) (*pagination.Pagination, error) {
	totalUser, err := u.ur.GetTotalUser(ctx)
	if err != nil {
		return nil, err
	}

	queryparams.CheckLimit(&queryParams)
	totalPage := totalUser / queryParams.Limit
	if totalUser%queryParams.Limit != 0 {
		totalPage += 1
	}
	queryparams.CheckPage(&queryParams, totalPage)

	users, err := u.ur.GetUsers(ctx, queryParams)
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

func (u AdminServiceImpl) Login(ctx context.Context, userAuth entity.User) (string, error) {
	isUserExists := u.ur.IsEmailExists(ctx, userAuth.Email)
	if !isUserExists {
		return "", apperror.NewErrStatusBadRequest(constant.LOGIN, apperror.ErrEmailOrPasswordInvalid, apperror.ErrEmailOrPasswordInvalid)
	}

	user, err := u.ur.GetUserByEmail(ctx, userAuth.Email)
	if err != nil {
		return "", err
	}
	userIdStr := strconv.Itoa(user.Id)

	if user.Role != constant.ADMIN {
		return "", apperror.NewErrStatusBadRequest(constant.LOGIN, apperror.ErrWrongRole, apperror.ErrWrongRole)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userAuth.Password))
	if err != nil {
		return "", apperror.NewErrStatusBadRequest(constant.LOGIN, apperror.ErrEmailOrPasswordInvalid, apperror.ErrEmailOrPasswordInvalid)
	}

	jwtToken := jwttoken.JwtTokenImpl{}
	token, err := jwtToken.GenerateJwtTokenForAuth(userIdStr, constant.USER)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s AdminServiceImpl) Register(ctx context.Context, user entity.User) error {
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
	user.Role = constant.ADMIN

	err = s.ur.AddUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (u AdminServiceImpl) UpdatePassword(ctx context.Context, user entity.User) error {
	isUserExists := u.ur.IsUserExists(ctx, user.Id)
	if !isUserExists {
		return apperror.NewErrStatusBadRequest(constant.UPDATE_USER_BY_ID, apperror.ErrUserNotExists, apperror.ErrUserNotExists)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return apperror.NewErrInternalServerError(constant.SERVER, apperror.ErrInternalServerError, apperror.ErrInternalServerError)
	}

	user.Password = string(hashedPassword)


	err = u.ur.UpdatePassword(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
