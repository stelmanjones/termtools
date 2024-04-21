package prompt

import (
	"fmt"
	"strings"

	"atomicgo.dev/keyboard/keys"
	"github.com/muesli/termenv"
)

// QuestionPrompt struct represents a prompt for a question.
type QuestionPrompt struct {
	Base[string]
	removeWhenDone bool
}

// NewQuestionPrompt creates a new QuestionPrompt with the provided label.
func NewQuestionPrompt(label string) *QuestionPrompt {
	p := &QuestionPrompt{
		Base: Base[string]{
			label: label,
		},
	}
	return p
}

// SetLabel sets the label of the QuestionPrompt.
func (p *QuestionPrompt) SetLabel(label string) {
	p.label = label
}

// RemoveWhenDone clears the prompt when done.
func (p *QuestionPrompt) RemoveWhenDone() *QuestionPrompt {
	p.removeWhenDone = true
	return p
}

// NoRemoveWhenDone sets the removeWhenDone flag to false.
func (p *QuestionPrompt) NoRemoveWhenDone() *QuestionPrompt {
	p.removeWhenDone = false
	return p
}

// render writes the label of the QuestionPrompt to the provided Output.
func (p *QuestionPrompt) render(out *termenv.Output) {
	if p.label != "" {
		_, err := out.WriteString(p.label + " ")
		if err != nil {
			fmt.Println(err)
		}
	}
}

// Run starts the QuestionPrompt and returns the user's input as a string.
func (p *QuestionPrompt) Run() (string, error) {
	out := termenv.DefaultOutput()
	var input strings.Builder
	p.render(out)
	ch := make(chan keys.Key)
	defer close(ch)
	go ListenForInput(ch)

outer:
	for key := range ch {
		switch key.Code {
		case keys.CtrlC, keys.CtrlD, keys.Esc:
			out.ClearLine()
			return "", ErrCanceledPrompt

		case keys.Enter:
			break outer
		case keys.RuneKey:
			out.WriteString(key.String())
			input.WriteString(key.String())
		default:
			out.WriteString(string(key.Runes))
			input.WriteString(string(key.Runes))
			continue

		}
	}
	out.ClearLine()
	out.MoveCursor(0, 0)
	return input.String(), nil
}
