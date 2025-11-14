package pkg

import (
	"errors"

	"github.com/gofiber/fiber/v2/log"
)

var ErrRepeatPasswordNotMatch = errors.New("repeat password and password does not match")
var ErrEmailAlreadyExists = errors.New("email already exists")
var ErrUsernameAlreadyExists = errors.New("username already exists")
var ErrRoleNotFound = errors.New("role not found")
var ErrUserNotFound = errors.New("user not found")
var ErrUnauthorized = errors.New("unauthorized")
var ErrBearerTokenNotFound = errors.New("bearer token not found")
var ErrInvalidBearerToken = errors.New("invalid bearer token")
var ErrInsufficientPermissions = errors.New("insufficient permissions")

type AppError struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

func NewAppError(statusCode int, message string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		Message:    message,
	}
}

func MapError(err error) *AppError {
	switch err {
	case ErrRepeatPasswordNotMatch:
		return NewAppError(400, err.Error())
	case ErrEmailAlreadyExists:
		return NewAppError(409, err.Error())
	case ErrUsernameAlreadyExists:
		return NewAppError(409, err.Error())
	case ErrRoleNotFound:
		return NewAppError(404, err.Error())
	case ErrUserNotFound:
		return NewAppError(404, err.Error())
	case ErrUnauthorized:
		return NewAppError(401, err.Error())
	case ErrBearerTokenNotFound:
		return NewAppError(401, err.Error())
	case ErrInvalidBearerToken:
		return NewAppError(401, err.Error())
	case ErrInsufficientPermissions:
		return NewAppError(401, err.Error())
	default:
		log.Error(err)
		return NewAppError(500, "internal server error")
	}
}
