package main

import (
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/log"
	"github.com/stelmanjones/termtools/realtime"
)

var logger = log.NewWithOptions(os.Stderr, log.Options{
	Level:           log.DebugLevel,
	TimeFormat:      "15:04:05",
	ReportTimestamp: true,
})

func main() {
	// done := make(chan struct{}, 1)
	// k := make(chan keys.Key, 1)
	// ctx := context.Background()
	//
	// shutdown := func() {
	// 	close(done)
	// 	close(k)
	// 	os.Exit(0)
	// 	logger.Info("Exiting 👋🏻")
	// }
	//
	// defer shutdown()
	// go internal.ListenForInput(k)
	// go tty.NotifyOnResize(ctx, done, func() {
	// 	size, err := tty.TermSize(os.Stdout.Fd())
	// 	if err != nil {
	// 		logger.Error("Error getting size")
	// 	}
	//
	// 	logger.Infof("Terminal size:\n    Width: %d\n    Height: %d", size.Width, size.Height)
	// })
	//
	// for key := range k {
	// 	switch key.Code {
	// 	case keys.CtrlC, keys.CtrlD, keys.Esc:
	// 		shutdown()
	// 	}
	// }
	//
	//
	// a := 1
	// b := 2
	// _ = usure.Assert(func() (success bool) { return a == b }, "waaa")
	// //	logger.Info("Running 🚀")
	// hotkeys.Start(remappedKeys)
	//

	out, err := realtime.New(os.Stdout)
	if err != nil {
		logger.Fatal(err)
	}

	out.Start()
	count := 0

	time.Sleep(time.Second * 2)
	out.WriteString(fmt.Sprintf("\n\nHello, world! %d", count))
	count++

	time.Sleep(time.Second * 2)
	out.WriteString(fmt.Sprintf("\n\nHello, world! %d", count))
	count++

	time.Sleep(time.Second * 2)
	out.WriteString(fmt.Sprintf("\n\nHello, world! %d", count))
	count++

	out.Stop()
}
