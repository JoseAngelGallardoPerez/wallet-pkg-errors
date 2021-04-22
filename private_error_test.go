package errors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrivateErrorError(t *testing.T) {
	privateError := PrivateError{Message: "Test message", OriginalError: errors.New("original error")}

	privateError.AddLogPair("myKey", "text")

	assert.Equal(t, privateError.Error(), "Message: Test message. Original error: original error.myKey.text.")
}
