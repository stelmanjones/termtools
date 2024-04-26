// Package spin provides a thread safe Spinner with various available character sets, prefix, suffix and color.
// The Spinner can be controlled using an Options pattern.
// Example:
//
//	s := spin.New(spin.BouncingBar, time.Millisecond * 10,
//	    spin.WithPrefix("SPINNING "),
//	    spin.WithSuffix("AFTER"),
//	    spin.WithColor(color.FgGreen),
//	    spin.WithFinalMsg("BYE!"))
//
//	s.Start()
//	time.Sleep(time.Second * 3)
//	s.Stop()
package spin

import (
	"atomicgo.dev/cursor"
	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"

	"fmt"
	"io"
	"math"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
	"unicode/utf8"

	"golang.org/x/term"

	"github.com/gookit/color"
)

var isWindows = runtime.GOOS == "windows"
var isWindowsTerminalOnWindows = len(os.Getenv("WT_SESSION")) > 0 && isWindows

// TODO: Use builder pattern for the Spinner struct.

// Spinner represents a thread-safe spinner with customizable options such as character sets, prefix, suffix, and color.
type Spinner struct {
	mu         *sync.RWMutex
	variant    SpinnerVariant
	running    bool
	stopChan   chan struct{}
	CancelKeys []keys.KeyCode
	Writer     io.Writer
	WriterFile *os.File
	Color      color.Color
	FinalMsg   string
	Prefix     string
	Suffix     string
	PreUpdate  func(s *Spinner)
	PostUpdate func(s *Spinner)
	lastOut    string
}

// New creates a new Spinner with the provided CharSet, delay, and options.
func New(variant SpinnerVariant, options ...Option) *Spinner {
	s := &Spinner{
		mu:         &sync.RWMutex{},
		variant:    variant,
		running:    false,
		CancelKeys: []keys.KeyCode{keys.CtrlC, keys.Escape},
		Writer:     os.Stdout,
		stopChan:   make(chan struct{}),
		//keyChan:    make(chan keys.Key),
		Color:      color.FgDefault,
		WriterFile: os.Stdout,
		Prefix:     "",
		Suffix:     "",
		FinalMsg:   "",
		PreUpdate:  nil,
		PostUpdate: nil,
	}

	for _, option := range options {
		option(s)
	}
	return s
}

// Option is a function that modifies a Spinner.
type Option func(s *Spinner)

// WithColor returns an Option function that sets the color of the spinner.
func WithColor(c color.Color) Option {
	return func(s *Spinner) {
		s.Color = c
	}
}

// WithPrefix returns an Option function that sets the Prefix field of a Spinner.
func WithPrefix(p string) Option {
	return func(s *Spinner) {
		s.Prefix = p
	}
}

// WithSuffix returns an Option function that sets the suffix of a Spinner.
func WithSuffix(sf string) Option {
	return func(s *Spinner) {
		s.Suffix = sf
	}
}

// WithFinalMsg returns an Option function that sets the final message of a Spinner.
func WithFinalMsg(fm string) Option {
	return func(s *Spinner) {
		s.FinalMsg = fm
	}
}

// WithWriter takes an io.Writer and sets the spinner output.
func WithWriter(w io.Writer) Option {
	return func(s *Spinner) {
		s.mu.Lock()
		s.Writer = w
		s.WriterFile = os.Stdout // emulate previous behavior for terminal check
		s.mu.Unlock()
	}
}

// WithCancelKeys returns an Option function that sets the cancelation keys for the Spinner.
func WithCancelKeys(keys []keys.KeyCode) Option {
	return func(s *Spinner) {
		s.CancelKeys = keys
	}
}

func isTerminal(s *Spinner) bool {
	return term.IsTerminal(int(s.WriterFile.Fd()))
}

// WithWriterFile adds the given writer to the spinner.
func WithWriterFile(f *os.File) Option {
	return func(s *Spinner) {
		s.mu.Lock()
		s.Writer = f     // io.Writer for actual writing
		s.WriterFile = f // file used only for terminal check
		s.mu.Unlock()
	}
}

// Start starts the spinner.
func (s *Spinner) Start() {
	s.running = true
	s.mu.Lock()
	fmt.Fprint(s.Writer, "\033[?25l")
	if !isTerminal(s) {
		s.ShowCursor()
		cursor.Show()
		s.mu.Unlock()
		return
	}
	s.mu.Unlock()

	go keyboard.Listen(func(key keys.Key) (stop bool, err error) {

		for _, c := range s.CancelKeys {

			if key.Code == c {
				s.stopChan <- struct{}{}
				return true, nil // Stop listener by returning true on Ctrl+C
			}
		}
		return false, nil
	})

	go func() {
		for {
			for i := 0; i < len(s.variant.chars)-1; i++ {
				select {
				case <-s.stopChan:
					s.mu.Lock()
					s.running = false
					s.mu.Unlock()
					fmt.Fprint(s.Writer, "\033[?25h")
					os.Exit(0)
				default:
					s.mu.RLock()
					out := fmt.Sprintf("\r%s%s%s", s.Prefix, s.Color.Sprintf("%s", s.variant.chars[i]), s.Suffix)
					fmt.Fprint(s.Writer, out)
					delay := s.variant.Interval
					s.mu.RUnlock()
					time.Sleep(time.Duration(delay) * time.Millisecond)
				}
			}
		}

	}()

}

// Stop stops the spinner, prints the final message if set, and signals the stop channel.
func (s *Spinner) Stop() {
	s.running = false
	s.mu.Lock()
	defer s.mu.Unlock()
	s.erase()
	fmt.Fprint(s.Writer, "\033[?25h")
	s.stopChan <- struct{}{}
	if s.FinalMsg != "" {
		s.erase()
		fmt.Fprintln(s.Writer, s.FinalMsg)
	}
}

// Restart stops the spinner and starts it again.
func (s *Spinner) Restart() {
	s.Stop()
	s.Start()
}

// ShowCursor shows the cursor.
func (s *Spinner) ShowCursor() {
	fmt.Fprint(s.Writer, "\x1b[?25h")
}

// HideCursor hides the cursor.
func (s *Spinner) HideCursor() {
	fmt.Fprintf(s.Writer, "\x1b[?25l")
}

func (s *Spinner) erase() {
	n := utf8.RuneCountInString(s.lastOut)
	if runtime.GOOS == "windows" && !isWindowsTerminalOnWindows {
		clearString := "\r" + strings.Repeat(" ", n) + "\r"
		fmt.Fprint(s.Writer, clearString)
		s.lastOut = ""
		return
	}

	numberOfLinesToErase := computeNumberOfLinesNeededToPrintString(s.lastOut)

	eraseCodeString := strings.Builder{}
	// current position is at the end of the last printed line. Start by erasing current line
	eraseCodeString.WriteString("\r\033[K") // start by erasing current line
	for i := 1; i < numberOfLinesToErase; i++ {
		// For each additional lines, go up one line and erase it.
		eraseCodeString.WriteString("\033[F\033[K")
	}
	fmt.Fprint(s.Writer, eraseCodeString.String())
	s.lastOut = ""
}

func computeNumberOfLinesNeededToPrintString(linePrinted string) int {
	terminalWidth := math.MaxInt // assume infinity by default to keep behaviour consistent with what we had before
	if term.IsTerminal(0) {
		if width, _, err := term.GetSize(0); err == nil {
			terminalWidth = width
		}
	}
	return computeNumberOfLinesNeededToPrintStringInternal(linePrinted, terminalWidth)
}

// isAnsiMarker returns if a rune denotes the start of an ANSI sequence
func isAnsiMarker(r rune) bool {
	return r == '\x1b'
}

// isAnsiTerminator returns if a rune denotes the end of an ANSI sequence
func isAnsiTerminator(r rune) bool {
	return (r >= 0x40 && r <= 0x5a) || (r == 0x5e) || (r >= 0x60 && r <= 0x7e)
}

// computeLineWidth returns the displayed width of a line
func computeLineWidth(line string) int {
	width := 0
	ansi := false

	for _, r := range line {
		// increase width only when outside of ANSI escape sequences
		if ansi || isAnsiMarker(r) {
			ansi = !isAnsiTerminator(r)
		} else {
			width += utf8.RuneLen(r)
		}
	}

	return width
}

func computeNumberOfLinesNeededToPrintStringInternal(linePrinted string, maxLineWidth int) int {
	lineCount := 0
	for _, line := range strings.Split(linePrinted, "\n") {
		lineCount++

		lineWidth := computeLineWidth(line)
		if lineWidth > maxLineWidth {
			lineCount += int(float64(lineWidth) / float64(maxLineWidth))
		}
	}

	return lineCount
}
