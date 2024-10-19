package realtime

import (
	"context"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/mattn/go-isatty"
	"github.com/stelmanjones/termtools/tty"
)

func fpsToDuration(fps int) time.Duration {
	return time.Duration(1000 / fps)
}

// Output is a terminal output that updates at a fixed rate and responds to terminal resize events.
type Output struct {
	Out         *os.File
	ticker      *time.Ticker
	done        chan bool
	mtx         *sync.Mutex
	size        *tty.WindowSize
	buf         strings.Builder
	RefreshRate int
}

// New creates a new Output.
func New(out *os.File) (*Output, error) {
	if !isatty.IsTerminal(out.Fd()) {
		return nil, ErrNotTerminal
	}

	return &Output{
		Out:         out,
		RefreshRate: 30,
		mtx:         &sync.Mutex{},
	}, nil
}

// Start starts the output update loop.
func (o *Output) Start() {
	tty.Output = o.Out
	o.buf.Reset()
	o.ticker = time.NewTicker(fpsToDuration(o.RefreshRate))
	o.done = make(chan bool)
	s, err := tty.TermSize(o.Out.Fd())
	if err != nil {
		o.size = &tty.WindowSize{
			Width:  0,
			Height: 0,
		}
	} else {
		o.size = s
	}
	go o.Listen()
}

var clear = fmt.Sprintf("%s[%dA%s[2K", "\x1b", 1, "\x1b")

// Flush writes the contents of the buffer to the terminal.
func (o *Output) Flush() error {
	o.mtx.Lock()
	defer o.mtx.Unlock()
	if o.buf.Len() == 0 {
		return nil
	}

	_, err := o.Out.WriteString(clear + o.buf.String())
	if err != nil {
		return err
	}

	o.buf.Reset()
	return nil
}

// Listen starts the event listener loop.
func (o *Output) Listen() {
	go tty.NotifyOnResize(context.Background(), o.done, func() {
		o.mtx.Lock()
		defer o.mtx.Unlock()
		s, err := tty.TermSize(o.Out.Fd())
		if err != nil {
			o.size = &tty.WindowSize{
				Width:  0,
				Height: 0,
			}
		} else {
			o.size = s
		}
	})

	for {
		select {
		case <-o.ticker.C:
			if o.ticker != nil {
				_ = o.Flush()
			}
		case <-o.done:
			o.mtx.Lock()
			o.ticker.Stop()
			o.ticker = nil
			o.mtx.Unlock()
			close(o.done)

		}
	}
}

// Stop stops the output update loop.
func (o *Output) Stop() {
	_ = o.Flush()
	o.done <- true
}

// Write appends the contents of buf to Output's buffer. Write always returns len(p), nil.
func (o *Output) Write(buf []byte) (n int, err error) {
	o.mtx.Lock()
	defer o.mtx.Unlock()
	return o.buf.Write(buf)
}

// WriteString appends the contents of s to Output's buffer. It returns the length of s and a nil error.
func (o *Output) WriteString(s string) (n int, err error) {
	o.mtx.Lock()
	defer o.mtx.Unlock()
	return o.buf.WriteString(s)
}

// Reset resets the outputs buffer.
func (o *Output) Reset() {
	o.mtx.Lock()
	defer o.mtx.Unlock()
	o.buf.Reset()
}

// Size returns the current size of the terminal.
func (o *Output) Size() *tty.WindowSize {
	o.mtx.Lock()
	defer o.mtx.Unlock()
	return o.size
}

// Len returns the number of characters in the buffer.
func (o *Output) Len() int {
	o.mtx.Lock()
	defer o.mtx.Unlock()
	return o.buf.Len()
}
