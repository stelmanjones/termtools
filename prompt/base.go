package prompt

import (
	"golang.org/x/exp/constraints"
)

type PromptValue interface {
	constraints.Ordered
}

// PromptBase is the base struct for all prompts
type PromptBase[T PromptValue] struct {
	// PromptType
	Value    T
	label    string
	selector string
	//theme *TermtoolsTheme
}

func (p *PromptBase[T]) WithSeparator(selector string) *PromptBase[T] {
	p.selector = selector
	return p
}
