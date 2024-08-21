package examples

import (
	"time"

	"github.com/gookit/color"
	"github.com/stelmanjones/termtools/spin"
)

func spinnerExample() {
	s := spin.New().
		WithPrefix("SPINNING ").
		WithSuffix("AFTER").
		WithColor(color.FgGreen).
		WithFinalMsg("BYE!").
		Build()

	s.Start()

	time.Sleep(time.Second * 3)
	s.Stop()

}
