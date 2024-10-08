package usure

type expectedValue[T comparable] struct {
	value T
}

func Expect[T comparable](value T) *expectedValue[T] {
	return &expectedValue[T]{
		value,
	}
}

func (e *expectedValue[T]) ToEqual(value T) bool {
	return Equal(e.value, value)
}

func (e *expectedValue[T]) ToNotEqual(value T) bool {
	return NotEqual(e.value, value)
}
