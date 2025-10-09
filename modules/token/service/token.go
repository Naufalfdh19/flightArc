package service

import (
	"context"
	"flight/modules/token/repo"
	"flight/pkg/apperror"
	"flight/pkg/constant"
	"flight/pkg/jwttoken"
	"strconv"
)

type TokenService interface {
	GenerateNewAccessToken(ctx context.Context, token string) (string, error) 
}

type tokenServiceImpl struct {
	r repo.TokenRepo
}

func NewTokenService(r repo.TokenRepo) tokenServiceImpl {
	return tokenServiceImpl{
		r: r,
	}
}


func (u tokenServiceImpl) GenerateNewAccessToken(ctx context.Context, token string) (string, error) {
	jwtToken := jwttoken.JwtTokenImpl{}
	tokenClaims := jwtToken.GetJwtTokenClaims(ctx, token)

	userId, err := strconv.Atoi(tokenClaims.UserID)
	if err != nil {
		return "", apperror.NewErrStatusUnauthorized(constant.GENERATE_TOKEN, apperror.ErrConvertingType, err)
	}

	isRefreshTokenExists := u.r.IsRefreshTokenExistsByUserID(ctx, userId)
	if !isRefreshTokenExists {
		return "", apperror.NewErrStatusUnauthorized(constant.GENERATE_TOKEN, apperror.ErrRefreshTokenNotExists, apperror.ErrRefreshTokenNotExists)
	}

	refreshToken, err := u.r.GetRefreshTokenByUserId(ctx, userId)
	if err != nil {
		return "", err
	}

	err = jwtToken.CheckJwtTokenForAuth(ctx, refreshToken.Value)
	if err != nil {
		return "", apperror.NewErrStatusUnauthorized(constant.GENERATE_TOKEN, apperror.ErrUnauthorized, err)
	}

	accessToken, err := jwtToken.GenerateAccessTokenForAuth(tokenClaims.UserID, tokenClaims.Role)
	if err != nil {
		return "", apperror.NewErrStatusUnauthorized(constant.GENERATE_TOKEN, apperror.ErrUnauthorized, err)
	}

	return accessToken, nil
}
