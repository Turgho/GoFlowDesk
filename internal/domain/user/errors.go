package user

import "errors"

var (
	ErrEmptyName          = errors.New("name cannot be empty")
	ErrEmptyEmail         = errors.New("email cannot be empty")
	ErrEmptyPassword      = errors.New("password cannot be empty")
	ErrInvalidUserRole    = errors.New("invalid user role")
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidID          = errors.New("invalid id")
	ErrNilUser            = errors.New("user is nil")
)
