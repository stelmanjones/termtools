package main

import (
	// "time"

	// "github.com/gookit/color"
	// "github.com/stelmanjones/microterm/spin"
	//"github.com/charmbracelet/log"
	//"github.com/charmbracelet/log"
	//"github.com/stelmanjones/termtools/kv"
	"fmt"

	"github.com/stelmanjones/termtools/prompt"
)

func main() {
	// s := spin.New(spin.GrowHorizontal, time.Millisecond*50, spin.WithPrefix("SPINNING "), spin.WithSuffix("AFTER"), spin.WithColor(color.FgGreen), spin.WithFinalMsg("BYE!"))
	// s.Start()
	// time.Sleep(time.Second * 3)
	// s.Stop()
	//

	//store := kv.New(kv.WithAuth("kekw1337"), kv.WithPort(6666))
	//if err := store.Serve(); err != nil {
	//	log.Error(err)
	//}
	pr := prompt.NewSelectionPrompt(prompt.WithChoice("Number 1"), prompt.WithChoice("Number 2"), prompt.WithChoice("Number 3"))
	pr.SetLabel("Select a number")
	pr.RemoveWhenDone()
	res := pr.Run()
	fmt.Println(res)

	intpr := prompt.NewSelectionPrompt(prompt.WithChoices(1, 2, 3, 4, 5))
	intpr.SetLabel("Select a number")
	intpr.RemoveWhenDone()
	res2 := intpr.Run()
	p := prompt.NewSelectionPrompt[string]()
	p.Run()
	fmt.Println(res2)

	// kv := easykv.New()
	// kv.Set("key", "value")
	// result ,_ := kv.Get("key")
	// log.Info(result.String())
}
