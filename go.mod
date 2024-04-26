module github.com/stelmanjones/termtools

replace github.com/stelmanjones/termtools/kv => ./kv

replace github.com/stelmanjones/termtools/prompt => ./prompt

replace github.com/stelmanjones/termtools/styles => ./styles

replace github.com/stelmanjones/termtools/usure => ./usure

replace github.com/stelmanjones/termtools/spin => ./spin

go 1.22.1

require (
	github.com/charmbracelet/log v0.4.0
	github.com/stelmanjones/termtools/kv v0.0.0-20240421154834-24bb8b0366d8
	github.com/stelmanjones/termtools/prompt v0.0.0-20240421154834-24bb8b0366d8
	github.com/stelmanjones/termtools/spin v0.0.0-20240421154834-24bb8b0366d8
	github.com/stelmanjones/termtools/usure v0.0.0-20240421154834-24bb8b0366d8
)

require (
	atomicgo.dev/cursor v0.2.0 // indirect
	atomicgo.dev/keyboard v0.2.9 // indirect
	github.com/containerd/console v1.0.4 // indirect
	github.com/emirpasic/gods v1.18.1 // indirect
	github.com/stelmanjones/termtools/styles v0.0.0-20240421154834-24bb8b0366d8 // indirect
	github.com/wI2L/jettison v0.7.4 // indirect
	golang.org/x/term v0.19.0 // indirect
)

require (
	github.com/aymanbagabas/go-osc52/v2 v2.0.1 // indirect
	github.com/bitly/go-simplejson v0.5.1 // indirect
	github.com/charmbracelet/lipgloss v0.10.0 // indirect
	github.com/go-logfmt/logfmt v0.6.0 // indirect
	github.com/gookit/color v1.5.4
	github.com/gorilla/mux v1.8.1 // indirect
	github.com/lucasb-eyer/go-colorful v1.2.0 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/muesli/reflow v0.3.0 // indirect
	github.com/muesli/termenv v0.15.2 // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/xo/terminfo v0.0.0-20220910002029-abceb7e1c41e // indirect
	golang.org/x/exp v0.0.0-20240416160154-fe59bbe5cc7f // indirect
	golang.org/x/sys v0.19.0 // indirect
)
