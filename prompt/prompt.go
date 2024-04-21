package prompt

func Ask(label string, removeWhenDone bool) (string, error) {

	p := NewQuestionPrompt(label)
	if removeWhenDone {
		p.RemoveWhenDone()
	}
	return p.Run()
}

func Select[T PromptValue](label string, choices []T, removeWhenDone bool) (*T, error) {
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

func Confirm(label string) (bool, error) {
	return NewConfirmationPrompt(label).Run()
}
