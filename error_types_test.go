package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrivateErrorAddLogPair(t *testing.T) {
	privateError := PrivateError{}

	privateError.AddLogPair("myKey", "text")

	assert.Equal(t, privateError.logPairs, []interface{}{"myKey", "text"})
}

func TestPrivateErrorAddLogPairMulti(t *testing.T) {
	privateError := PrivateError{}

	privateError.AddLogPair("myKey", "text")
	privateError.AddLogPair("newText", 42)

	assert.Equal(t, privateError.logPairs, []interface{}{"myKey", "text", "newText", 42})
}
