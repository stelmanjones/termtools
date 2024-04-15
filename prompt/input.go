package prompt

import (
	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
)

func ListenForInput(ch chan keys.Key) error {
	return keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		switch key.Code {
		case keys.Enter:
			{
				ch <- key
				return true, nil
			}
		default:
			{
				ch <- key
				return false, nil
			}
		}
	})
}
