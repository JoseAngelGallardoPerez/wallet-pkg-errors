package errors

import (
	"encoding/json"
	"net/http"

	"github.com/Confialink/wallet-pkg-response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/inconshreveable/log15"
)

func ErrorHandler(logger log15.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			processErrors(c, logger)
		}
	}
}

func processErrors(c *gin.Context, logger log15.Logger) {
	res := response.NewResponse()
	currentStatus := c.Writer.Status()

	data, ok := c.Get("data")
	if ok {
		res.Data = data
	}

	for _, e := range c.Errors {
		if validationErrors, ok := e.Err.(validator.ValidationErrors); ok {
			addValidationErrors(res, validationErrors)
			currentStatus = http.StatusBadRequest
			continue
		}
		if unmarshalError, ok := e.Err.(*json.UnmarshalTypeError); ok {
			addUnmarshalError(res, unmarshalError)
			currentStatus = http.StatusBadRequest
			continue
		}
		if typedErr, ok := e.Err.(TypedError); ok {
			typedErr.Process(res, logger)
			continue
		}

		logger.Error("caught error of unhandled type", "error", e.Err)
	}

	status := getStatus(currentStatus, c.Errors)
	c.JSON(status, res)
}

func getStatus(currentStatus int, errors []*gin.Error) int {
	if currentStatus != http.StatusOK {
		return currentStatus
	}

	if err := getFirstNotPrivateError(errors); err != nil {
		return err.GetHttpStatus()
	}

	return http.StatusInternalServerError
}

func addUnmarshalError(res *response.Response, err *json.UnmarshalTypeError) {
	extraErr := &PublicError{
		Code: "UNMARSHAL_JSON",
	}
	addPublicError(res, extraErr, &err.Field, TargetField)
}

func addPublicError(res *response.Response, err *PublicError, source *string, target string) {
	res.AddError(
		err.Title,
		&err.Details,
		&err.Code,
		source,
		target,
		err.Meta,
	)
}

func addValidationErrors(res *response.Response, validationErrors validator.ValidationErrors) {
	var source string
	for _, err := range validationErrors {
		extraErr := validationErrorToPublic(err)
		if len(err.Namespace()) > 0 {
			source = toLowerCamel(err.Namespace())
		} else {
			source = toLowerCamel(err.Field())
		}
		addPublicError(res, extraErr, &source, TargetField)
	}
}

func getFirstNotPrivateError(errors []*gin.Error) TypedError {
	for _, v := range errors {
		if typedErr, ok := v.Err.(TypedError); ok && typedErr.Type() != errorTypePrivate {
			return typedErr
		}
	}
	return nil
}
