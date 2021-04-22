package errors

import "errors"

var (
	ErrorInternalError = errors.New("something went wrong")
	ErrorForbidden     = errors.New("you are not allowed to perform this action")
)
