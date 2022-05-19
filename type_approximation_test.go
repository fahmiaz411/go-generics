package gogenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type A int

type N interface {
	~int
}

func B[T N](a T) T {
	return a
}

func TestApproximation(t *testing.T) {
	a := A(1)
	assert.Equal(t, A(1), B(a))
}