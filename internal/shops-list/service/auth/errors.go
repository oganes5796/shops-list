package auth

import "errors"

var (
	ErrUsernameNotFound = errors.New("username not found")
	ErrEmptyUsername    = errors.New("username is empty")
	ErrInvalidRole      = errors.New("invalid role")
)
