package prompt

import (
	"fmt"
	"strings"

	"atomicgo.dev/keyboard/keys"
	"github.com/muesli/termenv"
)

type QuestionPrompt struct {
	PromptBase[string]
	removeWhenDone bool
}

func NewQuestionPrompt(label string) *QuestionPrompt {
	p := &QuestionPrompt{

		PromptBase: PromptBase[string]{
			label: label,
		},
	}
	return p

}

func (p *QuestionPrompt) SetLabel(label string) {
	p.label = label
}

func (p *QuestionPrompt) RemoveWhenDone() *QuestionPrompt {
	p.removeWhenDone = true
	return p
}

func (p *QuestionPrompt) NoRemoveWhenDone() *QuestionPrompt {
	p.removeWhenDone = false
	return p
}

func (p *QuestionPrompt) render(out *termenv.Output) {
	if p.label != "" {

		_, err := out.WriteString(p.label + " ")
		if err != nil {
			fmt.Println(err)
		}
	}
}

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
