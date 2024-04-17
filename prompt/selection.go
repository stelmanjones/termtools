package prompt

import (
	"fmt"
	"strings"

	"atomicgo.dev/keyboard/keys"
	"github.com/muesli/termenv"
	"github.com/stelmanjones/termtools/styles"
)

var footer = "\n ↓/↑, tab/S-tab, j/k: down/up • enter: select\n"

type SelectionPrompt[T PromptValue] struct {
	PromptBase[T]
	Choices        []T
	index          int
	removeWhenDone bool
}

func NewSelectionPrompt[T PromptValue](choices ...T) *SelectionPrompt[T] {
	p := &SelectionPrompt[T]{

		PromptBase: PromptBase[T]{
			label:    "",
			selector: styles.Selector,
		},
		Choices: make([]T, 0),
		index:   0,
	}
	
	p.Choices = append(p.Choices, choices...)
	
	return p

}




func (p *SelectionPrompt[T])AddChoice(choice T)  {
	p.Choices = append(p.Choices, choice)
}

func (p *SelectionPrompt[T])AddChoices(choices ...T)  {
	p.Choices = append(p.Choices, choices...)
}

func (p *SelectionPrompt[T]) SetLabel(label string) {
	p.label = label
}

func (p *SelectionPrompt[T]) RemoveWhenDone() {
	p.removeWhenDone = true
}

func (p *SelectionPrompt[T]) increaseIndex() {
	if p.index == len(p.Choices)-1 {
		p.index = 0
	} else {
		p.index++
	}
}

func (p *SelectionPrompt[T]) decreaseIndex() {
	if p.index == 0 {
		p.index = len(p.Choices) - 1
	} else {
		p.index--
	}
}

func (p *SelectionPrompt[T]) render(out *termenv.Output) {
	out.ClearLines(len(p.Choices) + 4)
	var sb strings.Builder
	if p.label != "" {

		_, err := sb.WriteString(styles.Title.Styled(" " + p.label + " " + "\n\n"))
		if err != nil {
			fmt.Println(err)
		}
	}
	for i, option := range p.Choices {

		if i == p.index {
			_, err := sb.WriteString(p.selector + styles.SelectedOption.Styled(fmt.Sprintf("  %v\n", option)))
			if err != nil {
				fmt.Println(err)
			}
		} else {
			_, err := sb.WriteString(styles.NonSelectedOption.Styled(fmt.Sprintf("   %v\n", option)))
			if err != nil {
				fmt.Println(err)
			}
		}

	}

	_, err := sb.WriteString(styles.Dimmed.Styled(footer))
	if err != nil {
		fmt.Println(err)
	}
	out.WriteString(sb.String())

}

func (p *SelectionPrompt[T]) Run() (*T,error) {
	if len(p.Choices) == 0 {
		panic("No options provided")
	}
	out := termenv.DefaultOutput()
	out.HideCursor()
	defer out.ShowCursor()
	p.render(out)
	ch := make(chan keys.Key)
	go ListenForInput(ch)

	//fmt.Printf("opts: %d",len(p.Options))
outer:
	for key := range ch {
		switch key.Code {
		case keys.RuneKey:
			switch key.String() {
			case "j", "J":
				p.increaseIndex()

			case "k", "K":
				p.decreaseIndex()
			}

		case keys.Enter:
			break outer
		case keys.CtrlC, keys.CtrlD, keys.Esc:
			return new(T),ErrCanceledPrompt
		case keys.Down, keys.Tab:
			p.increaseIndex()
		case keys.Up, keys.ShiftTab:
			p.decreaseIndex()
		}
		p.render(out)
	}

	close(ch)
	if p.removeWhenDone {
		out.ClearLines(len(p.Choices) + 4)
	}
	return &p.Choices[p.index],nil

}
