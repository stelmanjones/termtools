package boxes

import (
	"io"
	"os"
	"strings"

	"github.com/stelmanjones/termtools/text"
)

// DefaultBox is a standard box.
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

// RoundedBox is a standard rounded box.
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
	title   string
	content []*text.Line
	Padding Padding
	Width   int
}

func (b *Box) buildHeader() string {
	var sb strings.Builder
	oddPadding := ""
	if text.OddVisibleLength(b.title) {
		oddPadding = " "
	}
	titleWidth := text.VisibleLength(b.title)
	sb.WriteString("\n" + b.Border.TopLeft)
	sb.WriteString(strings.Repeat(b.Border.Horizontal, (b.Width-2-titleWidth)/2))
	sb.WriteString(b.title)
	sb.WriteString(strings.Repeat(b.Border.Horizontal, (b.Width-2-titleWidth)/2))
	sb.WriteString(oddPadding)
	sb.WriteString(b.Border.TopRight)
	return sb.String()
}

func (b *Box) buildFooter() string {
	var sb strings.Builder
	sb.WriteString(b.Border.BottomLeft)
	sb.WriteString(strings.Repeat(b.Border.Horizontal, b.Width-2))
	sb.WriteString(b.Border.BottomRight + "\n")
	return sb.String()
}

// Sprint returns the rendered box as a string.
func (b *Box) Sprint(s ...string) string {
	var sb strings.Builder
	lenWithoutCorners := b.Width - 2
	usableWidth := b.Width - (2 + b.Padding.Left + b.Padding.Right)
	emptyLine := b.Border.Vertical + strings.Repeat(" ", lenWithoutCorners) + b.Border.Vertical

	// Header
	sb.WriteString(b.buildHeader() + "\n")

	// Top Padding
	for range b.Padding.Top {
		sb.WriteString(emptyLine + "\n")
	}

	// Content
	constrainedText := strings.Join(text.Chunks(strings.Join(s, ""), usableWidth), "")

	lines := text.MapLines(constrainedText, func(l *text.Line) *text.Line {
		l.Set(text.CenterText(strings.TrimSpace(l.Value()), usableWidth))
		return l
	})

	for _, l := range lines {
		oddPadding := ""
		if text.OddVisibleLength(l.Value()) {
			oddPadding = " "
		}
		sb.WriteString(b.Border.Vertical + strings.Repeat(" ", b.Padding.Left) + l.Value() + oddPadding + strings.Repeat(" ", b.Padding.Right) + b.Border.Vertical + "\n")
	}

	// Bottom Padding
	for range b.Padding.Bottom {
		sb.WriteString(emptyLine + "\n")
	}

	// Bottom Border
	sb.WriteString(b.buildFooter() + "\n")

	return sb.String()
}

// WithTitle sets the title of the box.
func (b *Box) WithTitle(title string) *Box {
	b.title = title
	return b
}

// Print prints the box to it's internal writer.
func (b *Box) Print(s ...string) {
	b.writer.Write([]byte(b.Sprint(s...)))
}
