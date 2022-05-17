package gogenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Comparable[T comparable](value1, value2 T) bool {
	return value1 == value2
}

func TestComparable(t *testing.T) {
	assert.Equal(t, true, Comparable(100, 100))
	assert.Equal(t, false, Comparable("eko", "fahmi"))
}