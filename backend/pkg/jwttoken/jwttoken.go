package jwttoken

import (
	"context"
	"flight/pkg/apperror"
	"flight/pkg/constant"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtTokenImpl struct{}

func NewJWT() *JwtTokenImpl {
	return &JwtTokenImpl{}
}

type JwtTokenClaims struct {
	Role   string
	UserID string
}

func NewJWTTokenClaims(userID, role string) *JwtTokenClaims {
	return &JwtTokenClaims{
		Role:   role,
		UserID: userID,
	}
}

type customClaims struct {
	Role string `json:"role"`
	jwt.RegisteredClaims
}

func (j JwtTokenImpl) GenerateAccessTokenForAuth(userID, role string) (string, error) {
	now := time.Now()

	registeredClaims := customClaims{
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:   "flightApp",
			Subject:  userID,
			IssuedAt: jwt.NewNumericDate(now),
			ExpiresAt: &jwt.NumericDate{
				Time: now.Add(24 * time.Hour),
			},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, registeredClaims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", apperror.NewErrStatusUnauthorized(constant.CHECK_AUTH, apperror.ErrUnexpectedSigningMethod, err)
	}
	return tokenString, nil
}

func (j JwtTokenImpl) GenerateRefreshToken(userID, role string) (string, error) {
	now := time.Now()

	registeredClaims := customClaims{
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:   "flightApp",
			Subject:  userID,
			IssuedAt: jwt.NewNumericDate(now),
			ExpiresAt: &jwt.NumericDate{
				Time: now.Add(24 * 7 * time.Hour),
			},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, registeredClaims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", apperror.NewErrStatusUnauthorized(constant.CHECK_AUTH, apperror.ErrUnexpectedSigningMethod, err)
	}
	return tokenString, nil
}

func (j JwtTokenImpl) CheckJwtTokenForAuth(ctx context.Context, tokenString string) error {
	_, err := jwt.Parse(
		tokenString,
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, apperror.ErrTokenInvalid
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		},
		jwt.WithIssuedAt(),
		jwt.WithIssuer("flightApp"),
		jwt.WithExpirationRequired(),
	)
	if err != nil {
		return err
	}

	return nil
}

func (j JwtTokenImpl) GetJwtTokenClaims(ctx context.Context, tokenString string) *JwtTokenClaims {
	token, _ := jwt.Parse(
		tokenString,
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, apperror.ErrTokenInvalid
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		},
		jwt.WithIssuedAt(),
		jwt.WithIssuer("flightApp"),
	)

	claims, exists := token.Claims.(jwt.MapClaims)
	if !exists {
		return nil
	}

	userID, err := claims.GetSubject()
	if err != nil {
		return nil
	}

	role := claims["role"].(string)

	jwtTokenClaims := NewJWTTokenClaims(userID, role)

	return jwtTokenClaims
}
