package realtime

import "errors"

var ErrNotTerminal = errors.New("provided writer is not a terminal.")
