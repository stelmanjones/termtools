package main

import (
	"errors"

	"github.com/charmbracelet/log"
	"github.com/stelmanjones/termtools/kv"
)

func retErr() error {
	return errors.New("fn error")
}

type kek struct {
	name      string
	something int
}

func main() {
	db := kv.New(kv.WithAuth("kekw1337"), kv.WithLimit(1))
	log.Fatal(db.Serve(6666))
	/*
		i := 0
		test := "Testing 123 string hello world"
		ticker := time.NewTicker(50 * time.Millisecond)
		out := termenv.DefaultOutput()

		for _ = range ticker.C {
			i++
			if i == 80 {
				ticker.Stop()
				out.ClearLine()
				out.MoveCursor(0, 0)
				return
			}
			out.MoveCursor(0, 0)
			out.ClearLine()
			out.WriteString(styles.AccentRed.Styled(styles.Glitch(test)))

		}
	*/

	// prompt.Ask("What is your name?", true)
	// prompt.Select("What is your favourite day of the week?", []string{"friday", "saturday"}, true)
	// prompt.Confirm("Are you sure you want to continue?")

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
