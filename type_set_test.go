package gogenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// type Set
type Number interface {
	int | int8 | int16 | int64 | float32 | float64
}

func Min[T Number](first, second T) T {
	return first - second
}

func TestSet(t *testing.T) {
	// float32
	assert.Equal(t, float32(1.3) - float32(1.2), Min[float32](1.3, 1.2))
	// int64
	assert.Equal(t, int64(-1), Min[int64](1, 2))
}