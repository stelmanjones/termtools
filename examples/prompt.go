package examples

import (
	"github.com/stelmanjones/termtools/prompt"
)

func test() {

	p := prompt.NewSelectionPrompt[string]()
	p.SetLabel("What is your favourite day of the week?")
	p.AddChoice("friday")
	p.AddChoice("saturday")
	p.RemoveWhenDone()

	q := prompt.NewQuestionPrompt("What is your favourite color?")
	q.RemoveWhenDone()
}
