package gogenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func NestingGeneric[T []E, E any](data T) E {
	return data[0]
}

func TestNestingGeneric(t *testing.T) {
	names := []string{
		"fahmi",
	}

	// inference
	first := NestingGeneric(names)

	assert.Equal(t, "fahmi", first)
}