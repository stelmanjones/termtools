package progress

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	"github.com/gosuri/uilive"
	"golang.org/x/term"
)

var (
	Stdout          io.Writer     = os.Stdout
	Stderr          io.Writer     = os.Stderr
	RefreshRate     time.Duration = 10 * time.Millisecond
	defaultProgress               = New()
)

type Progress struct {
	Out         io.Writer
	mtx         *sync.RWMutex
	done        chan bool
	live        *uilive.Writer
	Bars        []*Bar
	Width       int
	RefreshRate time.Duration
	current     int
	total       int
	elapsed     time.Duration
	showSteps   bool
	showPercent bool
	showElapsed bool
}

func New() *Progress {
	live := uilive.New()
	live.Out = Stdout
	return &Progress{
		Bars:        make([]*Bar, 0),
		RefreshRate: RefreshRate,
		Width: func() int {
			w, _, _ := term.GetSize(int(os.Stdout.Fd()))
			return w
		}(),
		live: live,
		done: make(chan bool),
		mtx:  &sync.RWMutex{},
	}
}

// Start starts the rendering the progress of progress bars using the DefaultProgress. It listens for updates using `bar.Set(n)` and new bars when added using `AddBar`
func Start() {
	defaultProgress.Start()
}

// Stop stops listening
func Stop() {
	defaultProgress.Stop()
}

// Listen listens for updates and renders the progress bars
func Listen() {
	defaultProgress.Listen()
}

// Add adds a new bar to the default progress bar
func Add(total int) *Bar {
	return defaultProgress.Add(total)
}

// Add adds a new bar to the progress bar
func (p *Progress) Add(total int) *Bar {
	bar := NewBar(total)
	p.mtx.Lock()
	defaultProgress.Bars = append(defaultProgress.Bars, bar)
	defaultProgress.mtx.Unlock()
	return bar
}

// SetOut sets the output writer for the progress bar
func (p *Progress) SetOut(o io.Writer) {
	p.mtx.Lock()
	defer p.mtx.Unlock()

	p.Out = o
	p.live.Out = o
}

// SetRefreshRate sets the refresh rate for the progress bar
func (p *Progress) SetRefreshRate(interval time.Duration) {
	p.mtx.Lock()
	defer p.mtx.Unlock()
	p.RefreshRate = interval
}

func (p *Progress) print() {
	p.mtx.Lock()
	defer p.mtx.Unlock()
	for _, bar := range p.Bars {
		fmt.Fprintln(p.live, bar.String())
	}
	p.live.Flush()
}

// Listen listens for updates and renders the progress bars
func (p *Progress) Listen() {
	for {

		p.mtx.Lock()
		interval := p.RefreshRate
		p.mtx.Unlock()

		select {
		case <-time.After(interval):
			p.print()
		case <-p.done:
			p.print()
			close(p.done)
			return
		}
	}
}

// Start starts the rendering the progress of progress bars. It listens for updates using `bar.Set(n)` and new bars when added using `AddBar`
func (p *Progress) Start() {
	go p.Listen()
}

// Stop stops listening
func (p *Progress) Stop() {
	p.done <- true
	<-p.done
}

// Bypass returns a writer which allows non-buffered data to be written to the underlying output
func (p *Progress) Bypass() io.Writer {
	return p.live.Bypass()
}
