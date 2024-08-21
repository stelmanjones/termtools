package spin

import (
	"io"
	"os"
	"sync"

	"atomicgo.dev/keyboard/keys"
	"github.com/gookit/color"
)

// SpinnerBuilder is a builder for the spinner.
type SpinnerBuilder struct {
	writer     io.Writer
	writerFile *os.File
	preUpdate  func(s *Spinner)
	postUpdate func(s *Spinner)
	finalMsg   string
	prefix     string
	suffix     string
	cancelKeys []keys.KeyCode
	variant    SpinnerVariant
	color      color.Color
}

// New returns a new spinner builder.
func New() *SpinnerBuilder {
	return &SpinnerBuilder{
		writer:     os.Stdout,
		writerFile: os.Stdout,
		preUpdate:  nil,
		postUpdate: nil,
		finalMsg:   "",
		prefix:     "",
		suffix:     "",
		cancelKeys: []keys.KeyCode{keys.CtrlC, keys.Escape},
		variant:    Dots1,
		color:      color.FgDefault,
	}
}

// WithWriter sets the writer of the spinner.
func (s *SpinnerBuilder) WithWriter(writer io.Writer) *SpinnerBuilder {
	s.writer = writer
	return s
}

// WithWriterFile sets the writer file of the spinner.
func (s *SpinnerBuilder) WithWriterFile(writerFile *os.File) *SpinnerBuilder {
	s.writerFile = writerFile
	return s
}

// WithPreUpdate sets the pre update function of the spinner.
func (s *SpinnerBuilder) WithPreUpdate(preUpdate func(s *Spinner)) *SpinnerBuilder {
	s.preUpdate = preUpdate
	return s
}

// WithPostUpdate sets the post update function of the spinner.
func (s *SpinnerBuilder) WithPostUpdate(postUpdate func(s *Spinner)) *SpinnerBuilder {
	s.postUpdate = postUpdate
	return s
}

// WithFinalMsg sets the final message of the spinner.
func (s *SpinnerBuilder) WithFinalMsg(finalMsg string) *SpinnerBuilder {
	s.finalMsg = finalMsg
	return s
}

// WithPrefix sets the prefix of the spinner.
func (s *SpinnerBuilder) WithPrefix(prefix string) *SpinnerBuilder {
	s.prefix = prefix
	return s
}

// WithSuffix sets the suffix of the spinner.
func (s *SpinnerBuilder) WithSuffix(suffix string) *SpinnerBuilder {
	s.suffix = suffix
	return s
}

// WithCancelKeys sets the keys that can cancel the spinner.
func (s *SpinnerBuilder) WithCancelKeys(cancelKeys []keys.KeyCode) *SpinnerBuilder {
	s.cancelKeys = cancelKeys
	return s
}

// WithVariant sets the variant of the spinner.
func (s *SpinnerBuilder) WithVariant(variant SpinnerVariant) *SpinnerBuilder {
	s.variant = variant
	return s
}

// WithColor sets the color of the spinner.
func (s *SpinnerBuilder) WithColor(color color.Color) *SpinnerBuilder {
	s.color = color
	return s
}

// Build builds a new spinner with the given options.
func (s *SpinnerBuilder) Build() *Spinner {
	return &Spinner{
		Writer:     s.writer,
		WriterFile: s.writerFile,
		stopChan: make(chan struct{}),
		mu: &sync.RWMutex{},
		PreUpdate:  s.preUpdate,
		PostUpdate: s.postUpdate,
		FinalMsg:   s.finalMsg,
		Prefix:     s.prefix,
		Suffix:     s.suffix,
		lastOut: "",
		CancelKeys: s.cancelKeys,
		variant:    s.variant,
		Color:      s.color,
	}
}
