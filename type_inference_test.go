package gogenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Inference[T comparable](i T) T {
	return i
}

func TestInference(t *testing.T) {
	assert.Equal(t, 32, Inference(32))
	assert.Equal(t, 64.0, Inference(64.0))	
	assert.Equal(t, 32, Inference(int64(32)))
	assert.Equal(t, 64.0, Inference(float32(64.0)))	
}