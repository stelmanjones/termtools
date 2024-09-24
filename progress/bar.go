package progress

import (
	"bytes"
	"fmt"
	"os"
	"sync"
	"time"

	"golang.org/x/term"
)

var (
	Head      byte = '>'
	Filler    byte = '='
	Empty     byte = '-'
	Delimiter byte = '|'
)

type Bar struct {
	TimeStarted time.Time
	mtx         *sync.RWMutex
	MarginLeft  int
	current     int
	elapsed     time.Duration
	Width       int
	total       int
	MarginRight int
	showPercent bool
	Bar         byte
	Empty       byte
	Delimiter   byte
	Head        byte
	showElapsed bool
	showSteps   bool
}

func NewBar(total int) *Bar {
	return &Bar{
		Head: Head,
		Bar:  Filler,

		Empty:       Empty,
		Delimiter:   Delimiter,
		MarginLeft:  1,
		MarginRight: 1,
		Width: func() int {
			w, _, _ := term.GetSize(int(os.Stdout.Fd()))
			return w
		}(),

		total:   total,
		current: 0,
		elapsed: 0,
		mtx:     &sync.RWMutex{},
	}
}

func (p *Bar) Increment() error {
	p.mtx.Lock()
	defer p.mtx.Unlock()

	if p.current >= p.total {
		return ErrTotalReached
	}

	p.current++
	return nil
}

func (p *Bar) Done() error {
	p.mtx.Lock()
	defer p.mtx.Unlock()

	if p.current < p.total {
		return ErrNotDone
	}

	return nil
}

func (b *Bar) Set(current int) error {
	b.mtx.Lock()
	defer b.mtx.Unlock()

	if current < 0 || current > b.total {
		return ErrInvalidValue
	}

	b.current = current
	return nil
}

func (b *Bar) Incr() bool {
	b.mtx.Lock()
	defer b.mtx.Unlock()

	if b.current+1 >= b.total {
		return false
	}

	var t time.Time
	if b.TimeStarted.IsZero() || b.TimeStarted.Equal(t) {
		b.TimeStarted = time.Now()
	}

	b.elapsed = time.Since(b.TimeStarted)

	b.current++
	return true
}

// Current returns the current progress of the bar
func (b *Bar) Current() int {
	b.mtx.RLock()
	defer b.mtx.RUnlock()
	return b.current
}

// CompletedPercent return the percent completed
func (b *Bar) CompletedPercent() float64 {
	return (float64(b.Current()) / float64(b.total)) * 100.00
}

// CompletedPercentString returns the formatted string representation of the completed percent
func (b *Bar) CompletedPercentString() string {
	return fmt.Sprintf("%3.f%%", b.CompletedPercent())
}

// TimeElapsed returns the time elapsed
func (b *Bar) TimeElapsed() time.Duration {
	b.mtx.RLock()
	defer b.mtx.RUnlock()
	return b.elapsed
}

// Bytes returns the byte presentation of the progress bar
func (b *Bar) Bytes() []byte {
	completedWidth := int(float64(b.Width) * (b.CompletedPercent() / 100.00))

	// add fill and empty bits
	var buf bytes.Buffer
	for i := 0; i < completedWidth; i++ {
		buf.WriteByte(b.Bar)
	}
	for i := 0; i < b.Width-completedWidth; i++ {
		buf.WriteByte(b.Empty)
	}

	// set head bit
	pb := buf.Bytes()
	if completedWidth > 0 && completedWidth < b.Width {
		pb[completedWidth-1] = b.Head
	}

	// set left and right ends bits
	pb[0], pb[len(pb)-1] = b.Delimiter, b.Delimiter

	return pb
}

func (b *Bar) String() string {
	return string(b.Bytes())
}
