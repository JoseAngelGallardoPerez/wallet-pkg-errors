package errors

import (
	"github.com/Confialink/wallet-pkg-response"
	"github.com/inconshreveable/log15"
)

type PublicError struct {
	Title      string
	Details    string
	Code       string
	HttpStatus int
	Meta       interface{}
}

func (e *PublicError) Process(response *response.Response, logger log15.Logger) {
	addPublicError(response, e, nil, TargetCommon)
}

func (e *PublicError) Type() ErrorType {
	return errorTypePublic
}

func (e *PublicError) Error() string {
	return e.Code
}

func (e *PublicError) GetHttpStatus() int {
	return e.HttpStatus
}
