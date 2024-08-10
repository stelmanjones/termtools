package main

import (
	"C"
	"os"

	"github.com/charmbracelet/log"
	"github.com/stelmanjones/termtools/hotkeys"
)

var logger = log.NewWithOptions(os.Stderr, log.Options{
	Level:           log.DebugLevel,
	TimeFormat:      "15:04:05",
	ReportTimestamp: true,
})


var remappedKeys = map[string]func(){
	"a": func() {
		hotkeys.DragMouseLeft(20)
	},
	"d": func() {
		hotkeys.DragMouseRight(20)
	},
	"w": func() {
		hotkeys.DragMouseUp(20)
	},
	"s": func() {
		hotkeys.DragMouseDown(20)
	},
}

func main() {
	hotkeys.Start(remappedKeys)
}
