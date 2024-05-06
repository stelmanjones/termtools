
package prompt

// Ask is a convenience function that creates a new question prompt and runs it.
func Ask(label string, removeWhenDone bool) (string, error) {

	p := NewQuestionPrompt(label)
	if removeWhenDone {
		p.RemoveWhenDone()
	}
	return p.Run()
}

// Select is a convenience function that creates a new selection prompt and runs it.
func Select[T Value](label string, choices []T, removeWhenDone bool) (*T, error) {
	p := NewSelectionPrompt[T]()
	p.SetLabel(label)
	for _, choice := range choices {
		p.AddChoice(choice)
	}
	if removeWhenDone {
		p.RemoveWhenDone()
	}

	return p.Run()
}

// Confirm is a convenience function that creates a new confirmation prompt and runs it.
func Confirm(label string) (bool, error) {
	return NewConfirmationPrompt(label).Run()
}
