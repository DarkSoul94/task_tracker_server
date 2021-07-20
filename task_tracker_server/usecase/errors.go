package usecase

import "errors"

var (
	ErrLoginFailed = errors.New("Invalid user name or password")
)
