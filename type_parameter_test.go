package gogenerics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Length[T any](param T) T {
	fmt.Println(param)
	return param
}

func TestSample(t *testing.T) {
	var resultString string = Length("Fahmi")
	assert.Equal(t, "Fahmi", resultString)

	var resultInt int = Length(100)
	assert.Equal(t, 100, resultInt)
}