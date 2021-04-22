package errors

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Confialink/wallet-pkg-response"
	"github.com/inconshreveable/log15"
)

type ValidationErrors struct {
	Errors []ValidationError
}

type ValidationError struct {
	Code   string
	Source string
	Title  string
	Meta   interface{}
}

func (e *ValidationErrors) Process(response *response.Response, logger log15.Logger) {
	for _, v := range e.Errors {
		func(v ValidationError) {
			v.process(response, logger)
		}(v)
	}
}

func (e *ValidationError) process(response *response.Response, logger log15.Logger) {
	response.AddError(
		e.Title,
		nil,
		&e.Code,
		&e.Source,
		TargetField,
		e.Meta,
	)
}

func (e *ValidationErrors) Type() ErrorType {
	return errorTypeValidation
}

func (e *ValidationErrors) Error() string {
	result := make([]string, len(e.Errors))
	for i, v := range e.Errors {
		result[i] = fmt.Sprintf("%s %s", v.Source, v.Code)
	}
	return strings.Join(result, "\n")
}

func (e *ValidationErrors) GetHttpStatus() int {
	return http.StatusBadRequest
}
