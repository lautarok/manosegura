package pkg

import "errors"

var ErrRepeatPasswordNotMatch = errors.New("repeat password and password does not match")
var ErrEmailAlreadyExists = errors.New("email already exists")
var ErrUsernameAlreadyExists = errors.New("username already exists")
var ErrRoleNotFound = errors.New("role not found")
var ErrUserNotFound = errors.New("user not found")
