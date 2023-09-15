package types

import "errors"

var (
	ErrRequestBody     = errors.New("failed to parse request body")
	ErrHash            = errors.New("failed to hash password")
	ErrComparePassword = errors.New("invalid password")
	ErrDatabase        = errors.New("database got error")
	ErrInvalidToken    = errors.New("invalid token")
	ErrUserNotFound    = errors.New("user not found")
)
