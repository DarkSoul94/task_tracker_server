package permissions

import "errors"

var (
	ErrWrongType = errors.New("Input data isn't map[string]string or []string")
)
