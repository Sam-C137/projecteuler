package helpers

import (
	"iter"
	"maps"
)

func NewSet[T comparable]() map[T]struct{} {
	return make(map[T]struct{})
}

func SetAdd[T comparable](set map[T]struct{}, value T) {
	_, ok := set[value]
	if !ok {
		set[value] = struct{}{}
	}
}

func SetDelete[T comparable](set map[T]struct{}, value T) {
	delete(set, value)
}

func SetSize(set map[any]struct{}) int {
	return len(set)
}

func SetValues[T comparable](set map[T]struct{}) iter.Seq[T] {
	return maps.Keys(set)
}
