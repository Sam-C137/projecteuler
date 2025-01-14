package helpers

func Map[T any, K any](slice []T, mapFn func(T) K) []K {
	out := make([]K, len(slice))

	for i := range slice {
		out[i] = mapFn(slice[i])
	}

	return out
}
