package prompt

import (
	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
)

// ListenForInput listens for input from the user.
func ListenForInput(ch chan keys.Key) error {
	return keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		switch key.Code {
		case keys.RuneKey:
			ch <- key
			return false, nil

		case keys.Enter, keys.CtrlC, keys.CtrlD, keys.Esc:
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
