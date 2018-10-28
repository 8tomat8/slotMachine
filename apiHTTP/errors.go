package apiHTTP

import (
	"errors"
)

var (
	ErrInvalidJWT      = errors.New("invalid JWT")
	ErrInvalidUserData = errors.New("user data must contain not empty uid, bet, chips fields")
)
