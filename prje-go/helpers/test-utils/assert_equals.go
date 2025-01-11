package test_utils

import "testing"

type Assert[V comparable] struct {
	T *testing.T
}

func (a *Assert[V]) ExpectEquals(expected V, result V) {
	if expected != result {
		a.T.Errorf("%v != %v; want %v, got %v", result, expected, result, expected)
	}
}
