package pkg

import "errors"

var ErrRepeatPasswordNotMatch = errors.New("repeat password and password does not match")
var ErrEmailAlreadyExists = errors.New("email already exists")
var ErrUsernameAlreadyExists = errors.New("username already exists")
var ErrRoleNotFound = errors.New("role not found")
var ErrUserNotFound = errors.New("user not found")
var ErrUnauthorized = errors.New("unauthorized")

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
	default:
		return NewAppError(500, "internal server error")
	}
}
