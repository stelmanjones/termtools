package main

import (
	// "time"

	// "github.com/gookit/color"
	// "github.com/stelmanjones/microterm/spin"
	//"github.com/charmbracelet/log"
	"github.com/charmbracelet/log"
	"github.com/stelmanjones/termtools/ezkv"
)

func main() {
	// s := spin.New(spin.GrowHorizontal, time.Millisecond*50, spin.WithPrefix("SPINNING "), spin.WithSuffix("AFTER"), spin.WithColor(color.FgGreen), spin.WithFinalMsg("BYE!"))
	// s.Start()
	// time.Sleep(time.Second * 3)
	// s.Stop()
	//

	r := ezkv.New(ezkv.WithAuth("kekw1337"), ezkv.WithPort(6666))
	if err := r.Run(); err != nil {
		log.Error(err)
	}

	// kv := easykv.New()
	// kv.Set("key", "value")
	// result ,_ := kv.Get("key")
	// log.Info(result.String())
}
