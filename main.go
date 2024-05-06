package main

import (
	"C"
	"os"

	"github.com/charmbracelet/log"
	"github.com/stelmanjones/termtools/remap"
)

var logger = log.NewWithOptions(os.Stderr, log.Options{
	Level:           log.DebugLevel,
	TimeFormat:      "15:04:05",
	ReportTimestamp: true,
})

var remappedKeys = map[string]func(){
	"a": func() {
		remap.DragMouseLeft(20)
	},
	"d": func() {
		remap.DragMouseRight(20)
	},
	"w": func() {
		remap.DragMouseUp(20)
	},
	"s": func() {
		remap.DragMouseDown(20)
	},
}

func main() {
	remap.Start(remappedKeys)
}
