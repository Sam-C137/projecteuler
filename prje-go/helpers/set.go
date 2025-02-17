package helpers

import (
	"iter"
	"maps"
)

type Set[T comparable] map[T]struct{}

func NewSet[T comparable]() Set[T] {
	return make(map[T]struct{})
}

func (set Set[T]) Add(value T) {
	_, ok := set[value]
	if !ok {
		set[value] = struct{}{}
	}
}

func (set Set[T]) Remove(value T) {
	delete(set, value)
}

func (set Set[T]) Size() int {
	return len(set)
}

func (set Set[T]) Values() iter.Seq[T] {
	return maps.Keys(set)
}
