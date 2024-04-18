package prompt

import (
	"fmt"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
	"github.com/charmbracelet/log"
	"github.com/muesli/termenv"
	"github.com/stelmanjones/termtools/styles"
)

type ConfirmationPrompt struct {
	PromptBase[string]
}

func NewConfirmationPrompt(label string) *ConfirmationPrompt {
	p := &ConfirmationPrompt{

		PromptBase: PromptBase[string]{
			label: label,
		},
	}
	return p

}

func (p *ConfirmationPrompt) SetLabel(label string) *ConfirmationPrompt {
	p.label = label
	return p
}

func (p *ConfirmationPrompt) render(out *termenv.Output) {
	if p.label != "" {

		_, err := out.WriteString(p.label + styles.Dimmed.Styled(" (y/n)"))
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (p *ConfirmationPrompt) Run() (bool, error) {
	out := termenv.DefaultOutput()
	p.render(out)
	var result bool
	ch := make(chan keys.Key)
	defer close(ch)
	go keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		switch key.Code {
		case keys.RuneKey:
			switch key.String() {
			case "y", "Y", "n", "N":
				ch <- key
				return true, nil

			default:
				ch <- key
				return false, nil

			}
		default:
			{
				ch <- key
				return false, nil
			}
		}
	})
outer:
	for key := range ch {
		switch key.Code {
		case keys.CtrlC, keys.CtrlD:
			out.ClearLine()
			log.Error("User cancelled")
			return false, ErrCanceledPrompt
		case keys.RuneKey:
			switch key.String() {
			case "y", "Y":
				result = true
				break outer

			case "n", "N":
				result = false
				break outer

			default:
				continue

			}
		default:
			continue

		}
	}

	out.ClearLine()
	out.MoveCursor(0, 0)

	return result, nil

}