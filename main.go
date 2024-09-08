package main

import (
	"C"
	"os"

	"github.com/charmbracelet/log"
	"github.com/stelmanjones/termtools/hotkeys"
)

import (
	"fmt"
	"image"
	"image/png"

	"github.com/stelmanjones/termtools/screencap"
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

func save(img *image.RGBA, filePath string) {
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	err = png.Encode(file, img)
	if err != nil {
		panic(err)
	}
}

func capture() {
	n := screencap.NumActiveDisplays()
	if n == 0 {
		panic("Active display not found")
	}

	all := image.Rect(0, 0, 0, 0)
	for i := 0; i < n; i++ {
		bounds := screencap.GetDisplayBounds(i)
		all = bounds.Union(all)

		img, err := screencap.CaptureRect(bounds)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%v\n", all)
		img, err = screencap.Capture(all.Min.X, all.Min.Y, all.Dx(), all.Dy())
		if err != nil {
			panic(err)
		}
		save(img, "all.png")
	}
}

func main() {
	//	logger.Info("Running 🚀")
	// hotkeys.Start(remappedKeys)
	/*
			s := "Line1\nLine2\nLine3"
			for line := range text.Lines(s) {
				log.Info("Line ->", "data", line, "index", line.Index(), "value", line.Value(), "runes", line.Runes(), "bytes", line.Bytes())
				line.Set(fmt.Sprintf("%s%s", line.Value(), " changed"))
				log.Info("Line ->", "data", line, "index", line.Index(), "value", line.Value(), "runes", line.Runes(), "bytes", line.Bytes())

			}


		s := spin.New().
			WithPrefix("SPINNING ").
			WithColor(color.FgGreen).
			WithFinalMsg("BYE!").
			Build()

		s.Start()

		time.Sleep(time.Second * 3)
		s.Stop()
	*/

	capture()
}
