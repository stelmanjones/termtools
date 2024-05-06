package remap

import (
	"os"

	"github.com/charmbracelet/log"
	hook "github.com/robotn/gohook"
)

var logger = log.NewWithOptions(os.Stderr, log.Options{
	Level:           log.DebugLevel,
	TimeFormat:      "15:04:05",
	ReportTimestamp: true,
})

// Start starts the remapper. It will listen for keydown events and execute the corresponding action if it exists.
// NOTE: Does not support special keys like CTRL, ALT, SHIFT, etc.
func Start(actions map[string]func()) error {

	ev := hook.Start()

	for e := range ev {
		if e.Kind == hook.KeyDown {
			if actions[string(e.Keychar)] != nil {
				actions[string(e.Keychar)]()
			}
		}
	}
	return nil

}
