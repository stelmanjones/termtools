package main

import (
	"errors"

	"github.com/stelmanjones/termtools/prompt"
	"github.com/stelmanjones/termtools/usure"
)

func retErr() error {
	return errors.New("fn error")
}

type kek struct {
	name      string
	something int
}

func main() {
	first := kek{"kek", 1}
	second := kek{"kwk", 2}
	usure.ExpectEqual("sadkek", first, second)
	

	prompt.Ask("What is your name?", true)
	prompt.Select("What is your favourite day of the week?", []string{"friday", "saturday"}, true)
	prompt.Confirm("Are you sure you want to continue?")

	/*
		p := prompt.NewSelectionPrompt[string]()
		p.SetLabel("What is your favourite day of the week?")
		p.AddChoice("friday")
		p.AddChoice("saturday")
		p.RemoveWhenDone()
		_, err := p.Run()
		if err != nil {
			log.Error(err)
		}

		q := prompt.NewQuestionPrompt("What is your favourite color?")
		q.RemoveWhenDone()
		name, err := q.Run()
		if err != nil {
			log.Error(err)
		}
		log.Info(name)
	*/

	/*
	   s := spin.New(spin.Dots1, 200*time.Millisecond, spin.WithPrefix("Loading... "), spin.WithColor(color.FgGreen))
	   s.Start()
	   time.Sleep(5 * time.Second)
	   s.Stop()
	*/
}
