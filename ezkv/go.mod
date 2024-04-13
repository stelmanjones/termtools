module github.com/stelmanjones/termtools/ezkv

replace github.com/stelmanjones/termtools => ../termtools

replace github.com/stelmanjones/termtools/ezkv/errors => ./errors

go 1.22.1

require (
	github.com/bitly/go-simplejson v0.5.1
	github.com/charmbracelet/log v0.4.0
	github.com/emirpasic/gods v1.18.1
	github.com/gorilla/mux v1.8.1
)

require github.com/xo/terminfo v0.0.0-20210125001918-ca9a967f8778 // indirect

require (
	github.com/aymanbagabas/go-osc52/v2 v2.0.1 // indirect
	github.com/charmbracelet/lipgloss v0.10.0 // indirect
	github.com/go-logfmt/logfmt v0.6.0 // indirect
	github.com/gookit/color v1.5.4
	github.com/lucasb-eyer/go-colorful v1.2.0 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/muesli/reflow v0.3.0 // indirect
	github.com/muesli/termenv v0.15.2 // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	golang.org/x/exp v0.0.0-20231006140011-7918f672742d // indirect
	golang.org/x/sys v0.15.0 // indirect
)
