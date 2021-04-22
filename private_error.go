package errors

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/Confialink/wallet-pkg-response"
	"github.com/inconshreveable/log15"
)

type PrivateError struct {
	Message       string
	OriginalError error
	logPairs      []interface{}
}

func (e *PrivateError) Process(response *response.Response, logger log15.Logger) {
	e.logPairs = append(e.logPairs, "originalError", e.OriginalError)
	logger.Error(e.Message, e.logPairs...)
}

func (e *PrivateError) Type() ErrorType {
	return errorTypePrivate
}

func (e *PrivateError) Error() string {
	buf := bytes.NewBuffer([]byte{})

	buf.WriteString(fmt.Sprintf("Message: %s.", e.Message))
	if e.OriginalError != nil {
		buf.WriteString(fmt.Sprintf(" Original error: %s.", e.OriginalError.Error()))
	}

	for _, e := range e.logPairs {
		buf.WriteString(fmt.Sprintf("%v.", e))
	}
	return buf.String()
}

func (e *PrivateError) AddLogPair(key string, value interface{}) {
	e.logPairs = append(e.logPairs, key, value)
}

func (e *PrivateError) GetHttpStatus() int {
	return http.StatusInternalServerError
}
