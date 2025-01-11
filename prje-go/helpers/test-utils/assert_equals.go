package testing

import "testing"

type Assert[V comparable] struct {
	t testing.T
}

func (a *Assert[V]) Equals(expected V, result V) {
	if expected != result {
		a.Errorf("%v != %v; want %v, got %v", result, expected, result, expected)
	}
}
