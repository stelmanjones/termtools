package progress




import (
	"errors"
)

var (
	ErrTotalReached = errors.New("total reached")
	ErrNotDone      = errors.New("not done")
	ErrInvalidWidth = errors.New("invalid width")
	ErrInvalidValue = errors.New("invalid value")
)

