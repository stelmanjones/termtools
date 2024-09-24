// Package usure provides a few assertions.
package usure

import (
	"reflect"

	"github.com/charmbracelet/log"
	"github.com/gookit/color"
)

func Assert(comparison func() (success bool), msg string) bool {
	result := comparison()
	if !result {
		panic("[assertion error] " + msg)
	}
	return result
}

// Nil checks if a is nil.
func Nil(a any) bool {
	return a == nil
}

// NotNil checks if a is not nil.
func NotNil(a any) bool {
	return a != nil
}

// IsInstance checks if a is an instance of T.
func IsInstance[T any](a T) bool {
	return reflect.TypeOf(a) == reflect.TypeFor[T]()
}

// Equal checks if a and b are equal.
func Equal[T any](a, b T) bool {
	return reflect.DeepEqual(a, b)
}

// NotEqual checks if a and b are not equal.
func NotEqual[T any](a, b T) bool {
	return !reflect.DeepEqual(a, b)
}

// ExpectEqual checks if a and b are equal and prints a message if the condition is false.
func ExpectEqual[T any](msg string, a, b T) bool {
	if !Equal(a, b) {
		log.Errorf("%s: %v %v is not equal to %v %v",
			color.FgRed.Render(msg),
			color.FgLightBlue.Render(
				reflect.TypeOf(a).Kind()), a,
			color.FgLightBlue.Render(
				reflect.TypeOf(b).Kind()), b)
		return false
	}
	return true
}
