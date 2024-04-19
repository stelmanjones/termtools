package usure

import (
	"reflect"

	"github.com/charmbracelet/log"
	"github.com/gookit/color"
)

func Nil(a any) bool {
	return a == nil
}

func NotNil(a any) bool {
	return a != nil
}

func IsInstance(a any, b any) bool {
	return reflect.TypeOf(a) == reflect.TypeOf(b)
}

func Equal[T any](a, b T) bool {
	return reflect.DeepEqual(a, b)
}

func NotEqual[T any](a, b T) bool {
	return !reflect.DeepEqual(a, b)
}

func ExpectEqual[T any](msg string, a, b T) {
	if !Equal(a, b) {
		log.Errorf("%s: %v %v is not equal to %v %v",
			color.FgRed.Render(msg),
			color.FgLightBlue.Render(
				reflect.TypeOf(a).Kind()), a,
			color.FgLightBlue.Render(
				reflect.TypeOf(b).Kind()), b)
	}
}

// HACK: Dont know if this function makes any sense
func ExpectError(msg string, r ...any) any {
	if r == nil {
		return nil
	}
	for _, v := range r {
		switch v.(type) {
		case error:
			{
				log.Errorf("%s:  %v", color.FgRed.Render(msg), v)
				return nil
			}
		default:
		}

	}
	return r
}
