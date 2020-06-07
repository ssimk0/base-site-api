package errors

import "errors"

var (
	ErrBadRequest          = errors.New("BAD_REQUEST")
	ErrInternalServerError = errors.New("INTERNAL_SERVER_ERROR")
	ErrNotFound            = errors.New("NOT_FOUND")
)
