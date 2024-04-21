
package examples

import (
	"github.com/stelmanjones/termtools/prompt"
)

func test() {

	// prompt.Ask("What is your name?", true)
	// prompt.Select("What is your favourite day of the week?", []string{"friday", "saturday"}, true)
	// prompt.Confirm("Are you sure you want to continue?")

	p := prompt.NewSelectionPrompt[string]()
	p.SetLabel("What is your favourite day of the week?")
	p.AddChoice("friday")
	p.AddChoice("saturday")
	p.RemoveWhenDone()

	q := prompt.NewQuestionPrompt("What is your favourite color?")
	q.RemoveWhenDone()
}
