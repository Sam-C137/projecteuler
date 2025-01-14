package helpers

func DefaultArg[T any](args []T, _default T) T {
	if len(args) < 1 {
		return _default
	} else {
		return args[0]
	}
}
