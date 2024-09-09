package main

import (
	"C"
	"os"

	"github.com/charmbracelet/log"
	"github.com/stelmanjones/termtools/hotkeys"
)

import (
	"image"
	"image/png"

	"github.com/stelmanjones/termtools/boxes"
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

	s := `Assure polite his really and others figure though. Day age advantages end sufficient eat expression travelling. Of on am father by agreed supply rather either. Own handsome delicate its property mistress her end appetite. Mean are sons too sold nor said. Son share three men power boy you. Now merits wonder effect garret own. Admiration stimulated cultivated reasonable be projection possession of. Real no near room ye bred sake if some. Is arranging furnished knowledge agreeable so. Fanny as smile up small. It vulgar chatty simple months turned oh at change of. Astonished set expression solicitude way admiration.
Preserved defective offending he daughters on or. Rejoiced prospect yet material servants out answered men admitted. Sportsmen certainty prevailed suspected am as. Add stairs admire all answer the nearer yet length. Advantages prosperous remarkably my inhabiting so reasonably be if. Too any appearance announcing impossible one. Out mrs means heart ham tears shall power every.
Promotion an ourselves up otherwise my. High what each snug rich far yet easy. In companions inhabiting mr principles at insensible do. Heard their sex hoped enjoy vexed child for. Prosperous so occasional assistance it discovered especially no. Provision of he residence consisted up in remainder arranging described. Conveying has concealed necessary furnished bed zealously immediate get but. Terminated as middletons or by instrument. Bred do four so your felt with. No shameless principle dependent household do.
`

	log.Info(s)
	log.Warn("-----------------------")
	log.Info(boxes.RoundedBox.Sprint(s))
}
