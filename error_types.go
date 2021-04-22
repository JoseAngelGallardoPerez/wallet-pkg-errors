package errors

import (
	"github.com/Confialink/wallet-pkg-response"
	"github.com/gin-gonic/gin"
	"github.com/inconshreveable/log15"
)

const (
	TargetField  = "field"
	TargetCommon = "common"
)

const (
	errorTypePrivate       ErrorType = "private"
	errorTypePublic        ErrorType = "public"
	errorTypeValidation    ErrorType = "validation"
	errorTypeUnmarshalJSON ErrorType = "unmarshalJSON"
)

type ErrorType string

type TypedError interface {
	Process(response *response.Response, logger log15.Logger)
	Type() ErrorType
	Error() string
	GetHttpStatus() int
}

func AddErrors(c *gin.Context, errs ...TypedError) {
	for _, e := range errs {
		c.Error(e)
	}
}

func AddShouldBindError(c *gin.Context, err error) {
	c.Error(ShouldBindToTyped(err))
}
