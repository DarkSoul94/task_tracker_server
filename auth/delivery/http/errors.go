package http

import "errors"

var (
	ErrInvalidAccessToken = errors.New("Invalid access token")
	ErrUserNameIsBlank    = errors.New("Username is blank")
	ErrPassIsBlank        = errors.New("Password is blank")
)
