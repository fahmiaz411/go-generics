package gogenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Comparable[T comparable](value1, value2 T) bool {
	return value1 == value2
}

func SumNumbers[K comparable, V Number](m map[K]V) V {
    var s V
    for _, v := range m {
        s += v
    }
    return s
}

// Inline type Set int64 | float64 or interface { int64 | float64 }
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
    var s V
    for _, v := range m {
        s += v
    }
    return s
}

func TestComparable(t *testing.T) {
	assert.Equal(t, true, Comparable(100, 100))
	assert.Equal(t, false, Comparable("eko", "fahmi"))

	assert.Equal(t, int64(10), SumIntsOrFloats(map[string]int64{
		"1":5,
		"2":5,
	}))
	assert.Equal(t, 11.2, SumIntsOrFloats(map[string]float64{
		"1": 5.6,
		"2": 5.6,
	}))
}