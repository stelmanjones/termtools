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
	Value     T
	label     string
	separator string
	theme *Theme
}



func (p *PromptBase[T]) WithSeparator(separator string) *PromptBase[T] {
	p.separator = separator
	return p
}
