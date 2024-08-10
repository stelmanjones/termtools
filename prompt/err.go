package prompt

import (
	"errors"
)

var (
	// ErrCanceledPrompt is returned when the user cancels the prompt.
	ErrCanceledPrompt = errors.New("user canceled")
	ErrNoChoices      = errors.New("selection prompt cannot be empty")
)
