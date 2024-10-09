package internal

import (
	"bytes"
	"net/http"
	"runtime/debug"
	"sync"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
	"github.com/charmbracelet/log"
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

// FireAndForget executes `f()` in a new go routine and auto recovers if panic.
//
// **Note:** Use this only if you are not interested in the result of `f()`
// and don't want to block the parent go routine.
func FireAndForget(f func(), wg ...*sync.WaitGroup) {
	if len(wg) > 0 && wg[0] != nil {
		wg[0].Add(1)
	}

	go func() {
		if len(wg) > 0 && wg[0] != nil {
			defer wg[0].Done()
		}

		defer func() {
			if err := recover(); err != nil {
				log.Errorf("RECOVERED FROM PANIC (safe to ignore): %v", err)
				log.Debug(string(debug.Stack()))
			}
		}()

		f()
	}()
}

// Fetch fetches the given url and returns the response body as a string.
func Fetch(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	var buf bytes.Buffer
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
