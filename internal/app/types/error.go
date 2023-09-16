package types

import "errors"

var (
	ErrRequestBody           = errors.New("failed to parse request body")
	ErrHash                  = errors.New("failed to hash password")
	ErrComparePassword       = errors.New("invalid password")
	ErrDatabase              = errors.New("database got error")
	ErrInvalidToken          = errors.New("invalid token")
	ErrUserNotFound          = errors.New("user not found")
	ErrProductNotFound       = errors.New("product not found")
	ErrProductInCartNotFound = errors.New("product in cart not found")
	ErrInsufficientBalance   = errors.New("insufficient balance")
)
