package service

import (
	"context"
	"flight/pkg/apperror"
	"flight/pkg/constant"
	"flight/pkg/jwttoken"
)

type TokenService interface {
	GenerateNewAccessToken(ctx context.Context, refreshToken string) (string, error)
}

type tokenServiceImpl struct{}

func NewTokenService() tokenServiceImpl {
	return tokenServiceImpl{}
}

func (u tokenServiceImpl) GenerateNewAccessToken(ctx context.Context, refreshToken string) (string, error) {
	jwtToken := jwttoken.JwtTokenImpl{}
	tokenClaims := jwtToken.GetJwtTokenClaims(ctx, refreshToken)

	err := jwtToken.CheckJwtTokenForAuth(ctx, refreshToken)
	if err != nil {
		return "", apperror.NewErrStatusUnauthorized(constant.GENERATE_TOKEN, apperror.ErrUnauthorized, err)
	}

	accessToken, err := jwtToken.GenerateAccessTokenForAuth(tokenClaims.UserID, tokenClaims.Role)
	if err != nil {
		return "", apperror.NewErrStatusUnauthorized(constant.GENERATE_TOKEN, apperror.ErrUnauthorized, err)
	}

	return accessToken, nil
}
