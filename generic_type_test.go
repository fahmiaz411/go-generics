package gogenerics

import (
	"fmt"
	"testing"
)

type Items[T any] []T

type DataStruct[T any] struct {
	Name string
	Detail *T
}

type DataMap[K, V comparable] map[K]V

func PrintBag[T any](bag Items[T]){
	for _, v := range bag {
		fmt.Println(v)
	}
}

func TestGenericType(t *testing.T) {
	items := []string{
		"Bag1",
		"Bag2",
		"Bag3",
	}
	PrintBag(items)
}