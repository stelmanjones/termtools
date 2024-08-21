package main

import (
	"C"
	"os"

	"github.com/charmbracelet/log"
	"github.com/stelmanjones/termtools/hotkeys"
	"github.com/stelmanjones/termtools/text"
)
import "fmt"

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
	//	logger.Info("Running 🚀")
	// hotkeys.Start(remappedKeys)
	s := "Line1\nLine2\nLine3"
	for line := range text.Lines(s) {
		log.Info("Line ->", "data", line, "index", line.Index(), "value", line.Value(), "runes", line.Runes(), "bytes", line.Bytes())
		line.Set(fmt.Sprintf("%s%s", line.Value(), " changed"))
		log.Info("Line ->", "data", line, "index", line.Index(), "value", line.Value(), "runes", line.Runes(), "bytes", line.Bytes())

	}
}
