package global

import "errors"

var (
	ErrInternalError  = errors.New("internal error")
	ErrIncorectParams = errors.New("incorect params")
)
