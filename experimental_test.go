package gogenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/maps"
)

func ExperimentalMin[T constraints.Ordered](first, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func TestExperimentalMaps(t *testing.T) {
	first := map[string]string{
		"Name":"fahmi",
	}
	second := map[string]string{
		"Name":"fahmi",
	}

	assert.True(t, maps.Equal(first, second))
}