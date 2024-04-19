package examples

import (
	"time"

	"github.com/gookit/color"
	"github.com/stelmanjones/termtools/spin"
)

func Spin() {
	s := spin.New(spin.BouncingBar, time.Millisecond*10,
		spin.WithPrefix("SPINNING "),
		spin.WithSuffix("AFTER"),
		spin.WithColor(color.FgGreen),
		spin.WithFinalMsg("BYE!"))

	s.Start()
	time.Sleep(time.Second * 3)
	s.Stop()

}
