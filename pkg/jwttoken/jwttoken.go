package jwttoken

import (
	"context"
	"flight/pkg/apperror"
	"flight/pkg/constant"
	"log"
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

func (j JwtTokenImpl) GenerateJwtTokenForAuth(userID, role string) (string, error) {

	now := time.Now()

	registeredClaims := customClaims{
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:   "flight",
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

func (j JwtTokenImpl) ParseJwtTokenForAuth(ctx context.Context, tokenString string) (*JwtTokenClaims, error) {
	token, err := jwt.Parse(
		tokenString,
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, apperror.ErrTokenInvalid
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		},
		jwt.WithIssuedAt(),
		jwt.WithIssuer("flight"),
		jwt.WithExpirationRequired(),
	)
	if err != nil {
		return nil, err
	}

	claims, exists := token.Claims.(jwt.MapClaims)
	if !exists {
		return nil, err
	}

	userID, err := claims.GetSubject()
	if err != nil {
		return nil, err
	}

	log.Println(userID)
	log.Println(claims["role"])
	role := claims["role"].(string)

	jwtTokenClaims := NewJWTTokenClaims(userID, role)

	return jwtTokenClaims, nil
}
