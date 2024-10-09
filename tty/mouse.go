package tty

import (
	"regexp"
	"strconv"
)

var (
	unknownCSIRe  = regexp.MustCompile(`^\x1b\[[\x30-\x3f]*[\x20-\x2f]*[\x40-\x7e]`)
	mouseSGRRegex = regexp.MustCompile(`(\d+);(\d+);(\d+)([Mm])`)
)

type (
	MouseAction int
	MouseButton int
)

const (
	MouseActionPress MouseAction = iota
	MouseActionRelease
	MouseActionMotion
)

const (
	MouseButtonNone MouseButton = iota
	MouseButtonLeft
	MouseButtonMiddle
	MouseButtonRight
	MouseButtonWheelUp
	MouseButtonWheelDown
	MouseButtonWheelLeft
	MouseButtonWheelRight
	MouseButtonBackward
	MouseButtonForward
	MouseButton10
	MouseButton11
)

type MouseEvent struct {
	X      int
	Y      int
	Shift  bool
	Alt    bool
	Ctrl   bool
	Action MouseAction
	Button MouseButton
}

func (m MouseEvent) IsWheel() bool {
	return m.Button == MouseButtonWheelUp || m.Button == MouseButtonWheelDown ||
		m.Button == MouseButtonWheelLeft || m.Button == MouseButtonWheelRight
}

// HandleMessage parses a single mouse event.
func HandleMessage(b []byte) (int, MouseEvent) {
	const mouseEventX10Len = 6
	if len(b) >= mouseEventX10Len && b[0] == '\x1b' && b[1] == '[' {
		switch b[2] {
		case 'M':
			return mouseEventX10Len, parseX10MouseEvent(b)
		case '<':
			if matchIndices := mouseSGRRegex.FindSubmatchIndex(b[3:]); matchIndices != nil {
				// SGR mouse events length is the length of the match plus the length of the escape sequence
				mouseEventSGRLen := matchIndices[1] + 3 //nolint:gomnd
				return mouseEventSGRLen, parseSGRMouseEvent(b)
			}
		}
	}
	return 0, MouseEvent{}
}

func parseSGRMouseEvent(buf []byte) MouseEvent {
	str := string(buf[3:])
	matches := mouseSGRRegex.FindStringSubmatch(str)
	if len(matches) != 5 { //nolint:gomnd
		// Unreachable, we already checked the regex in `detectOneMsg`.
		panic("invalid mouse event")
	}

	b, _ := strconv.Atoi(matches[1])
	px := matches[2]
	py := matches[3]
	release := matches[4] == "m"
	m := parseMouseButton(b, true)

	// Wheel buttons don't have release events
	// Motion can be reported as a release event in some terminals (Windows Terminal)
	if m.Action != MouseActionMotion && !m.IsWheel() && release {
		m.Action = MouseActionRelease
	}

	x, _ := strconv.Atoi(px)
	y, _ := strconv.Atoi(py)

	// (1,1) is the upper left. We subtract 1 to normalize it to (0,0).
	m.X = x - 1
	m.Y = y - 1

	return m
}

const x10MouseByteOffset = 32

// Parse X10-encoded mouse events; the simplest kind. The last release of X10
// was December 1986, by the way. The original X10 mouse protocol limits the Cx
// and Cy coordinates to 223 (=255-032).
//
// X10 mouse events look like:
//
//	ESC [M Cb Cx Cy
//
// See: http://www.xfree86.org/current/ctlseqs.html#Mouse%20Tracking
func parseX10MouseEvent(buf []byte) MouseEvent {
	v := buf[3:6]
	m := parseMouseButton(int(v[0]), false)

	// (1,1) is the upper left. We subtract 1 to normalize it to (0,0).
	m.X = int(v[1]) - x10MouseByteOffset - 1
	m.Y = int(v[2]) - x10MouseByteOffset - 1

	return m
}

func parseMouseButton(b int, isSGR bool) MouseEvent {
	var m MouseEvent
	e := b
	if !isSGR {
		e -= 32
	}

	const (
		bitShift  = 0b0000_0100
		bitAlt    = 0b0000_1000
		bitCtrl   = 0b0001_0000
		bitMotion = 0b0010_0000
		bitWheel  = 0b0100_0000
		bitAdd    = 0b1000_0000 // additional buttons 8-11

		bitsMask = 0b0000_0011
	)

	if e&bitAdd != 0 {
		m.Button = MouseButtonBackward + MouseButton(e&bitsMask)
	} else if e&bitWheel != 0 {
		m.Button = MouseButtonWheelUp + MouseButton(e&bitsMask)
	} else {
		m.Button = MouseButtonLeft + MouseButton(e&bitsMask)
		// X10 reports a button release as 0b0000_0011 (3)
		if e&bitsMask == bitsMask {
			m.Action = MouseActionRelease
			m.Button = MouseButtonNone
		}
	}

	// Motion bit doesn't get reported for wheel events.
	if e&bitMotion != 0 && !m.IsWheel() {
		m.Action = MouseActionMotion
	}

	// Modifiers
	m.Alt = e&bitAlt != 0
	m.Ctrl = e&bitCtrl != 0
	m.Shift = e&bitShift != 0
	return m
}
