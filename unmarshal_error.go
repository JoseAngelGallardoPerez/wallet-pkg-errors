package errors

import (
	"encoding/json"
	"net/http"

	"github.com/Confialink/wallet-pkg-response"
	"github.com/inconshreveable/log15"
)

type UnmarshalError struct {
	Err *json.UnmarshalTypeError
}

func (e *UnmarshalError) Process(response *response.Response, logger log15.Logger) {
	extraErr := &PublicError{
		Code: CodeUnmarshalJSON,
	}
	addPublicError(response, extraErr, &e.Err.Field, TargetField)
}

func (e *UnmarshalError) Type() ErrorType {
	return errorTypeUnmarshalJSON
}

func (e *UnmarshalError) Error() string {
	return CodeUnmarshalJSON
}

func (e *UnmarshalError) GetHttpStatus() int {
	return http.StatusBadRequest
}
