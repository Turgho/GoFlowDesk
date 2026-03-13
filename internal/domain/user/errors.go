package user

import "errors"

var (
	// validation errors
	ErrEmptyName     = errors.New("name cannot be empty")
	ErrEmptyEmail    = errors.New("email cannot be empty")
	ErrEmptyPassword = errors.New("password cannot be empty")

	ErrInvalidEmail    = errors.New("invalid email")
	ErrInvalidPassword = errors.New("invalid password")

	// domain errors
	ErrInvalidID       = errors.New("invalid id")
	ErrInvalidUserRole = errors.New("invalid user role")

	// business errors
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrUserNotFound       = errors.New("user not found")

	// safety errors
	ErrNilUser = errors.New("user is nil")
)
