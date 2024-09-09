package boxes

import (
	"io"
	"os"
	"strings"

	"github.com/stelmanjones/termtools/text"
)

var DefaultBox = Box{
	Border: LightBorder,
	Padding: Padding{
		2,
		2,
		2,
		2,
	},
	writer: os.Stdout,
	Width:  100,
}

var RoundedBox = Box{
	Border: RoundBorder,
	Padding: Padding{
		2,
		2,
		2,
		2,
	},
	writer: os.Stdout,
	Width:  100,
}

//revive:disable:exported
var (
	LightBorder = Border{
		Horizontal:  "─",
		Vertical:    "│",
		TopLeft:     "┌",
		TopRight:    "┐",
		BottomLeft:  "└",
		BottomRight: "┘",
	}

	HeavyBorder = Border{
		Horizontal:  "━",
		Vertical:    "┃",
		TopLeft:     "┏",
		TopRight:    "┓",
		BottomLeft:  "┗",
		BottomRight: "┛",
	}

	DoubleBorder = Border{
		Horizontal:  "═",
		Vertical:    "║",
		TopLeft:     "╔",
		TopRight:    "╗",
		BottomLeft:  "╚",
		BottomRight: "╝",
	}

	RoundBorder = Border{
		Horizontal:  "─",
		Vertical:    "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "╰",
		BottomRight: "╯",
	}
)

//revive:enable:exported

// Padding is the spaceing between the box content and the border.
type Padding struct {
	Top    int
	Left   int
	Right  int
	Bottom int
}

// Margin is the spaceing between the box border and the window border.
type Margin struct {
	Top    int
	Left   int
	Right  int
	bottom int
}

// Border is the border of the box.
type Border struct {
	TopLeft     string
	TopRight    string
	Vertical    string
	Horizontal  string
	BottomLeft  string
	BottomRight string
}

// Box is the box to be drawn.
type Box struct {
	writer  io.Writer
	Border  Border
	content []*text.Line
	Padding Padding
	Margin  Margin
	Width   int
}

// TODO: Clean this s*** up
func (b *Box) Sprint(s ...string) string {
	var sb strings.Builder
	lenWithoutCorners := b.Width - 2
	usableWidth := b.Width - (2 + b.Padding.Left + b.Padding.Right)
	emptyLine := b.Border.Vertical + strings.Repeat(" ", lenWithoutCorners) + b.Border.Vertical
	sb.WriteString("\n" + b.Border.TopLeft + strings.Repeat(b.Border.Horizontal, lenWithoutCorners) + b.Border.TopRight + "\n")
	for range b.Padding.Top {
		sb.WriteString(emptyLine + "\n")
	}

	constrainedText := strings.Join(text.Chunks(strings.Join(s, ""), usableWidth), "")

	lines := text.MapLines(constrainedText, func(l *text.Line) *text.Line {
		l.Set(text.CenterText(l.Value(), usableWidth))
		return l
	})

	for _, l := range lines {

		rightP := func() int {
			if len(l.Value())%2 != 0 {
				return b.Padding.Right + 1
			}
			return b.Padding.Right
		}()
		sb.WriteString(b.Border.Vertical + strings.Repeat(" ", b.Padding.Left) + l.Value() + strings.Repeat(" ", rightP) + b.Border.Vertical + "\n")
	}

	for range b.Padding.Bottom {
		sb.WriteString(emptyLine + "\n")
	}
	sb.WriteString(b.Border.BottomLeft + strings.Repeat(b.Border.Horizontal, lenWithoutCorners) + b.Border.BottomRight + "\n")

	return sb.String()
}

func (b *Box) Print(s ...string) {
	b.writer.Write([]byte(b.Sprint(s...)))
}
