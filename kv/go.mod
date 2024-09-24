module github.com/stelmanjones/termtools/kv

replace github.com/stelmanjones/termtools => ../termtools

replace github.com/stelmanjones/termtools/kv/errors => ./errors

go 1.22.1

require (
	github.com/bitly/go-simplejson v0.5.1
	github.com/charmbracelet/log v0.4.0
	github.com/emirpasic/gods v1.18.1
	github.com/gorilla/mux v1.8.1
)

require github.com/xo/terminfo v0.0.0-20220910002029-abceb7e1c41e // indirect

require (
	github.com/aymanbagabas/go-osc52/v2 v2.0.1 // indirect
	github.com/charmbracelet/lipgloss v0.10.0 // indirect
	github.com/go-logfmt/logfmt v0.6.0 // indirect
	github.com/gookit/color v1.5.4
	github.com/lucasb-eyer/go-colorful v1.2.0 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.16 // indirect
	github.com/muesli/reflow v0.3.0 // indirect
	github.com/muesli/termenv v0.15.2 // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/wI2L/jettison v0.7.4
	golang.org/x/exp v0.0.0-20240416160154-fe59bbe5cc7f // indirect
	golang.org/x/sys v0.19.0 // indirect
)
