package prompt

import "golang.org/x/exp/constraints"

// Value is a type constraint that represents any type that is ordered.
type Value interface {
	constraints.Ordered
}

// Base is the base struct for all prompts
type Base[T Value] struct {
	// PromptType
	label    string
	selector string
	// theme *TermtoolsTheme
}

// SetSelector sets the selector for the prompt.
func (p *Base[T]) SetSelector(selector string) *Base[T] {
	p.selector = selector
	return p
}

// Ask is a convenience function that creates a new question prompt and runs it.
func Ask(label string, defaultValue string, removeWhenDone bool) (string, error) {
	p := NewQuestionPrompt(label)
	if removeWhenDone {
		p.RemoveWhenDone()
	}
	if defaultValue != "" {
		p.SetDefault(defaultValue)
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
