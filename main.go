package main

import (
	// "time"

	// "github.com/gookit/color"
	// "github.com/stelmanjones/microterm/spin"
	//"github.com/charmbracelet/log"
	//"github.com/charmbracelet/log"
	//"github.com/stelmanjones/termtools/kv"
	"github.com/charmbracelet/log"
	"github.com/stelmanjones/termtools/prompt"
)

func main() {

	p := prompt.NewSelectionPrompt[string]()
	p.AddChoice("monday")
	p.AddChoice("tuesday")
	p.AddChoice("wednesday")
	p.AddChoice("thursday")
	p.RemoveWhenDone()
	_,err := p.Run()
	if err != nil {
		log.Error(err)
	}

	q := prompt.NewQuestionPrompt("What is your name?")
	q.RemoveWhenDone()
	name,err := q.Run()
	if err != nil {
		log.Error(err)
	}
	log.Info(name)

	// c := prompt.NewConfirmationPrompt("Are you sure?")

	

	// kv := easykv.New()
	// kv.Set("key", "value")
	// result ,_ := kv.Get("key")
	// log.Info(result.String())
}
