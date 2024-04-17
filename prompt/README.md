# Prompts

The Prompt module is a part of the `termtools` package and provides a set of
interactive prompts for command-line interfaces. It includes selection prompts,
question prompts, and confirmation prompts.

## Install
```go
import "github.com/stelmanjones/termtools/prompt"
```

## Features

- **Selection Prompt**: Allows users to select an option from a list of choices.
  The
  [`NewSelectionPrompt`](command:_github.copilot.openSymbolInFile?%5B%22prompt%2Fselection.go%22%2C%22NewSelectionPrompt%22%5D "prompt/selection.go")
  function is used to create a new selection prompt. Choices can be added using
  the `AddChoice` method.

```go
p := prompt.NewSelectionPrompt[int]()
p.AddChoice(1)
p.AddChoices(2,3,4,5,6,7)
p.RemoveWhenDone()
result, err := p.Run()
```

- **Question Prompt**: Asks users a question and waits for their input. The
  `NewQuestionPrompt` function is used to create a new question prompt.

```go
q := prompt.NewQuestionPrompt("What is your name?")
result, err := q.Run()
```

- **Confirmation Prompt**: Asks users a yes/no question. The
  [`NewConfirmationPrompt`](command:_github.copilot.openSymbolInFile?%5B%22prompt%2Fconfirm.go%22%2C%22NewConfirmationPrompt%22%5D "prompt/confirm.go")
  function is used to create a new confirmation prompt.

```go
c := prompt.NewConfirmationPrompt("Are you sure?")
result, err := c.Run()
```

