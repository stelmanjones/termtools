//go:build darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris || aix || zos
// +build darwin dragonfly freebsd linux netbsd openbsd solaris aix zos

package tty

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/term"
)

type WindowSize struct {
	Width, Height int
}

func newSize(w, h int) *WindowSize {
	return &WindowSize{w, h}
}

func TermSize(fd uintptr) (*WindowSize, error) {
	w, h, err := term.GetSize(int(fd))
	if err != nil {
		return nil, err
	}
	return newSize(w, h), nil
}

// NotifyOnResize listens for window size changes and runs the handler every time it does.
func NotifyOnResize(ctx context.Context, done chan bool, handler func()) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGWINCH)

	defer func() {
		signal.Stop(sig)
		close(done)
	}()

	for {
		select {
		case <-ctx.Done():
			return
		case <-sig:
		}
		handler()
	}
}
