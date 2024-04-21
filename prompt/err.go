package prompt

import (
	"errors"
)

var (
	// ErrCanceledPrompt is returned when the user cancels the prompt.
	ErrCanceledPrompt = errors.New("user canceled")
)
