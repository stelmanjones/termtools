package prompt

import (
	"golang.org/x/exp/constraints"
)

// Value is a type constraint that represents any type that is ordered.
type Value interface {
	constraints.Ordered
}

// Base is the base struct for all prompts
type Base[T Value] struct {
	// PromptType
	label    string
	selector string
	//theme *TermtoolsTheme
}

// SetSelector sets the selector for the prompt.
func (p *Base[T]) SetSelector(selector string) *Base[T] {
	p.selector = selector
	return p
}
