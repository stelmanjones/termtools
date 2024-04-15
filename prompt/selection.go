package prompt

import (
	"fmt"
	"strings"

	"atomicgo.dev/keyboard/keys"
	"github.com/gookit/color"
	"github.com/muesli/termenv"
)

var footer = "\nup:  ↑/shift+tab  down: ↓/tab, select: enter\n"
var s termenv.Style

type SelectionPrompt[T PromptValue] struct {
	PromptBase[T]
	Options        []T
	index          int
	removeWhenDone bool
}

func NewSelectionPrompt[T PromptValue](options ...SelectionPromptOption[T]) *SelectionPrompt[T] {
	p := &SelectionPrompt[T]{
		PromptBase: PromptBase[T]{
			theme:     ThemeCharm(),
			label:     "",
			separator: color.Hex("#FF69B4").Sprint(color.OpBold.Render(">")),
		},
		Options: make([]T, 0),
		index:   0,
	}
	for _, option := range options {
		option(p)
	}
	return p

}

type SelectionPromptOption[T PromptValue] func(*SelectionPrompt[T])

func WithChoices[T PromptValue](choices ...T) SelectionPromptOption[T] {
	return func(p *SelectionPrompt[T]) {
		p.Options = choices
	}
}

func WithChoice[T PromptValue](choice T) SelectionPromptOption[T] {
	return func(p *SelectionPrompt[T]) {
		p.Options = append(p.Options, choice)
	}
}

func (p *SelectionPrompt[T]) SetLabel(label string) {
	p.label = label
}

func (p *SelectionPrompt[T]) RemoveWhenDone() {
	p.removeWhenDone = true
}

func (p *SelectionPrompt[T]) increaseIndex() {
	if p.index == len(p.Options)-1 {
		p.index = 0
	} else {
		p.index++
	}
}

func (p *SelectionPrompt[T]) decreaseIndex() {
	if p.index == 0 {
		p.index = len(p.Options) - 1
	} else {
		p.index--
	}
}

func (p *SelectionPrompt[T]) render(out *termenv.Output) {
	out.ClearLines(len(p.Options) + 4)
	var sb strings.Builder
	_, err := sb.WriteString(p.theme.Focused.Title.Render(p.label + "\n\n"))
	if err != nil {
		fmt.Println(err)
	}
	for i, option := range p.Options {

		if i == p.index {
			_, err := sb.WriteString(p.theme.Focused.SelectSelector.Render()+p.theme.Focused.SelectedOption.Render(fmt.Sprintf("%v\n", option)))
			if err != nil {
				fmt.Println(err)
			}
		} else {
			_, err := sb.WriteString(p.theme.Focused.Option.Render(fmt.Sprintf("  %v\n", option)))
			if err != nil {
				fmt.Println(err)
			}
		}

	}

	_, err = sb.WriteString(p.theme.Blurred.Base.Render(footer))
	out.Write([]byte(sb.String()))
	if err != nil {
		fmt.Println(err)
	}

}

func (p *SelectionPrompt[T]) Run() T {
	if len(p.Options) == 0 {
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
		case keys.Enter:
			break outer
		case keys.Down, keys.Tab:
			p.increaseIndex()
		case keys.Up, keys.ShiftTab:
			p.decreaseIndex()
		}
		p.render(out)
	}

	close(ch)
	if p.removeWhenDone {
		out.ClearLines(len(p.Options) + 4)
	}
	return p.Options[p.index]

}
