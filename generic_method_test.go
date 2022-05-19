package gogenerics

import (
	"fmt"
	"testing"
)

type StructMethod[T any] struct {
	First  T
	Second T
}

func (s *StructMethod[_]) Hello(name string) string {
	return name
}

func (s *StructMethod[T]) ChangeFirst(first T) T {
	s.First = first
	return s.First
}

func TestGenericMethod(t *testing.T) {
	data := StructMethod[string]{
		First: "fahmi",		
	}

	data.ChangeFirst("eko")
	fmt.Println(data.First)
}