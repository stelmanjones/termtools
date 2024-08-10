package hotkeys

import "github.com/go-vgo/robotgo"

const (
	// HookEnabled honk enable status
	HookEnabled  = 1 // iota
	HookDisabled = 2

	KeyDown = 3
	KeyHold = 4
	KeyUp   = 5

	MouseUp    = 6
	MouseHold  = 7
	MouseDown  = 8
	MouseMove  = 9
	MouseDrag  = 10
	MouseWheel = 11

	FakeEvent = 12

	// Keychar could be v
	CharUndefined = 0xFFFF
	WheelUp       = -1
	WheelDown     = 1
)

// MoveCursorLeft moves the cursor left by the given amount
func MoveCursorLeft(amount int) {
	robotgo.MoveRelative(-amount, 0)
}

// MoveCursorRight moves the cursor right by the given amount
func MoveCursorRight(amount int) {
	robotgo.MoveRelative(amount, 0)
}

// MoveCursorUp moves the cursor up by the given amount
func MoveCursorUp(amount int) {
	robotgo.MoveRelative(0, -amount)
}

// MoveCursorDown moves the cursor down by the given amount
func MoveCursorDown(amount int) {
	robotgo.MoveRelative(0, amount)
}

// DragMouseLeft moves the cursor left by the given amount
func DragMouseLeft(amount int) {
	robotgo.Toggle("left")
	robotgo.MilliSleep(50)
	robotgo.MoveRelative(-amount, 0)
	robotgo.Toggle("left", "up")
}

// DragMouseRight moves the cursor right by the given amount
func DragMouseRight(amount int) {
	robotgo.Toggle("right")
	robotgo.MilliSleep(50)
	robotgo.MoveRelative(amount, 0)
	robotgo.Toggle("right", "up")
}

// DragMouseUp moves the cursor up by the given amount
func DragMouseUp(amount int) {
	robotgo.Toggle("up")
	robotgo.MilliSleep(50)
	robotgo.MoveRelative(0, -amount)
	robotgo.Toggle("up", "up")
}

// DragMouseDown moves the cursor down by the given amount
func DragMouseDown(amount int) {
	robotgo.Toggle("down")
	robotgo.MilliSleep(50)
	robotgo.MoveRelative(0, amount)
	robotgo.Toggle("down", "up")
}
