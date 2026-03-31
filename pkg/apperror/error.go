package apperror

import (
	"errors"
	"net/http"
)

type ErrorDetail struct {
	Field   string
	Message string
}

type ErrorStruct struct {
	Field         string
	Message       string
	Status        int
	SpecificError error
}

func (es ErrorStruct) Error() string {
	return es.SpecificError.Error()
}

func NewErrStatusBadRequest(field string, sentinel, err error) *ErrorStruct {
	return &ErrorStruct{
		Field:         field,
		Message:       sentinel.Error(),
		Status:        http.StatusBadRequest,
		SpecificError: err,
	}
}

func NewErrStatusUnauthorized(field string, sentinel, err error) *ErrorStruct {
	return &ErrorStruct{
		Field:         field,
		Message:       sentinel.Error(),
		Status:        http.StatusUnauthorized,
		SpecificError: err,
	}
}

func NewErrStatusNotFound(field string, sentinel, err error) *ErrorStruct {
	return &ErrorStruct{
		Field:         field,
		Message:       sentinel.Error(),
		Status:        http.StatusNotFound,
		SpecificError: err,
	}
}

func NewErrInternalServerError(field string, sentinel, err error) *ErrorStruct {
	return &ErrorStruct{
		Field:         field,
		Message:       sentinel.Error(),
		Status:        http.StatusInternalServerError,
		SpecificError: err,
	}
}

var (
	ErrInternalServerError     = errors.New("internal server error")
	ErrUserNotExists           = errors.New("user not exists")
	ErrRoleNotExists           = errors.New("role not exists")
	ErrAirlineNotExists        = errors.New("airline not exists")
	ErrTransactionFailed       = errors.New("transaction failed")
	ErrBindingRequest          = errors.New("binding request failed")
	ErrConvertingType          = errors.New("converting type error")
	ErrUnexpectedSigningMethod = errors.New("unexpected signing method")
	ErrTokenInvalid            = errors.New("token invalid")
	ErrEmailOrPasswordInvalid  = errors.New("email or password invalid")
	ErrPasswordInvalid         = errors.New("password invalid")
	ErrUsernameInvalid         = errors.New("username invalid")
	ErrUnauthorized            = errors.New("unauthorized")
	ErrUserExists              = errors.New("user exists")
	ErrPlaneExists             = errors.New("plane exists")
	ErrPlaneNotExists          = errors.New("plane not exists")
	ErrWrongRole               = errors.New("wrong role")
	ErrRefreshTokenNotExists   = errors.New("refresh token not exists")
	ErrAccessTokenNotExists    = errors.New("access token not exists")
	ErrGetBookings             = errors.New("get bookings failed")
	ErrAddBookings             = errors.New("add bookings failed")
)
