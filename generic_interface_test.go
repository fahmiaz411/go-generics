package gogenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type GetterSetter[T any] interface {
	GetValue() T
	SetValue(value T)
}

func ChangeValue[T any](param GetterSetter[T], value T) T {
	param.SetValue(value)
	return param.GetValue()
}

type GenericStruct[T any] struct {
	Value T
}

func (g *GenericStruct[T]) GetValue() T {
	return g.Value
}

func (g *GenericStruct[T]) SetValue(value T) {
	g.Value = value
}

func TestGenericInterface(t *testing.T) {
	data := GenericStruct[string]{}	

	assert.Equal(t, "fahmi", ChangeValue[string](&data, "fahmi"))
}