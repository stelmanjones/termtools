package main

import (
	// "time"

	// "github.com/gookit/color"
	// "github.com/stelmanjones/microterm/spin"
	//"github.com/charmbracelet/log"
	//"github.com/charmbracelet/log"
	//"github.com/stelmanjones/termtools/kv"
	//"time"

	// "github.com/charmbracelet/log"
	//	"github.com/gookit/color"
	"github.com/charmbracelet/log"
	"github.com/stelmanjones/termtools/kv"
	// "github.com/stelmanjones/termtools/prompt"
	// "github.com/stelmanjones/termtools/spin"
)

func main() {

	db := kv.New()
	log.Fatal(db.Serve(9999))
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
