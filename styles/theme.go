package styles

import (
	"github.com/muesli/termenv"
)

var s termenv.Style

var (
	red      = termenv.RGBColor("#ff8272")
	green    = termenv.RGBColor("#b4fa72")
	yellow   = termenv.RGBColor("#fefdc2")
	blue     = termenv.RGBColor("#a5d5fe")
	neonBlue = termenv.RGBColor("#1ec9ff")
	pink     = termenv.RGBColor("#ff8ffd")
	magenta  = termenv.RGBColor("#d0d1fe")
	white    = termenv.RGBColor("#f1f1f1")
	darkGray = termenv.RGBColor("#616161")
)

var (
	Title             = s.Foreground(white).Bold()
	Foreground        = s.Foreground(white)
	AccentRed         = s.Bold().Foreground(red)
	AccentBlue        = s.Bold().Foreground(neonBlue)
	AccentGreen       = s.Bold().Foreground(green)
	RoundSelector     = s.Foreground(red).Styled("●")
	Selector          = s.Foreground(red).Styled("❯")
	Dimmed            = s.Faint()
	Highlight         = s.Foreground(green)
	Accent            = s.Foreground(pink)
	SelectedOption    = s.Bold().Foreground(green)
	NonSelectedOption = s.Faint()
	Warning           = s.Bold().Foreground(yellow)
)

