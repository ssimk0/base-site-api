package errors

import "errors"

var (
	BadRequest          = errors.New("BAD_REQUEST")
	InternalServerError = errors.New("INTERNAL_SERVER_ERROR")
	NotFound            = errors.New("NOT_FOUND")
)
