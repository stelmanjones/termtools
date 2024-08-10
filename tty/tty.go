package tty

import (
	"os"
	"strconv"
)

// Screen
const (
	SaveSeq           = "\x1b[?47h"
	RestoreSeq        = "\x1b[?47l"
	ClearSeq          = "\x1b[2J"
	EnterAltScreenSeq = "\x1b[?1049h"
	LeaveAltScreenSeq = "\x1b[?1049l"
)

// Cursor
const (
	CursorHomeSeq      = "\x1b[H"
	CursorUpSeq        = "\x1b[A"
	CursorDownSeq      = "\x1b[B"
	CursorRightSeq     = "\x1b[C"
	CursorLeftSeq      = "\x1b[D"
	CursorSaveSeq      = "\x1b[7"
	CursorRestoreSeq   = "\x1b[8"
	CursorHideSeq      = "\x1b[?25l"
	CursorShowSeq      = "\x1b[?25h"
	CursorEraseLineSeq = "\x1b[2K"
)

// Requires Parameters
const (
	CursorEraseNLines = "\x1b[2K"
	CursorMoveSeq     = "\x1b[y;xH"
	CursorColumnSeq   = "\x1b[xG"
)

// Output is the output writer.
var Output = os.Stdout

// SaveScreen saves the screen.
func SaveScreen() (err error) {
	if _, err = Output.WriteString(SaveSeq); err != nil {
		return err
	}
	return nil
}

// RestoreScreen restores the screen.
func RestoreScreen() (err error) {
	if _, err = Output.WriteString(RestoreSeq); err != nil {
		return err
	}
	return nil
}

// ClearScreen clears the screen.
func ClearScreen() (err error) {
	if _, err = Output.WriteString(ClearSeq); err != nil {
		return err
	}
	return nil
}

// EnterAltScreen enters the alternate screen.
func EnterAltScreen() (err error) {
	if _, err = Output.WriteString(EnterAltScreenSeq); err != nil {
		return err
	}
	return nil
}

// LeaveAltScreen leaves the alternate screen.
func LeaveAltScreen() (err error) {
	if _, err = Output.WriteString(LeaveAltScreenSeq); err != nil {
		return err
	}
	return nil
}

// CursorHome moves the cursor to the home position.
func CursorHome() (err error) {
	if _, err = Output.WriteString(CursorHomeSeq); err != nil {
		return err
	}
	return nil
}

// CursorUp moves the cursor up n rows.
func CursorUp(n int) (err error) {
	if _, err = Output.WriteString(CursorUpSeq + strconv.Itoa(n)); err != nil {
		return err
	}
	return nil
}

// CursorDown moves the cursor down n rows.
func CursorDown(n int) (err error) {
	if _, err = Output.WriteString(CursorDownSeq + strconv.Itoa(n)); err != nil {
		return err
	}
	return nil
}

// CursorRight moves the cursor right n columns.
func CursorRight(n int) (err error) {
	if _, err = Output.WriteString(CursorRightSeq + strconv.Itoa(n)); err != nil {
		return err
	}
	return nil
}

// CursorLeft moves the cursor left n columns.
func CursorLeft(n int) (err error) {
	if _, err = Output.WriteString(CursorLeftSeq + strconv.Itoa(n)); err != nil {
		return err
	}
	return nil
}

// CursorEraseLine erases the current line.
func CursorEraseLine() (err error) {
	if _, err = Output.WriteString(CursorEraseLineSeq); err != nil {
		return err
	}
	return nil
}

// CursorMove moves the cursor to the y-th row and x-th column.
func CursorMove(y, x int) (err error) {
	if _, err = Output.WriteString(CursorMoveSeq + strconv.Itoa(y) + ";" + strconv.Itoa(x) + "H"); err != nil {
		return err
	}
	return nil
}

// CursorColumn sets the cursor to the x-th column of the current line.
func CursorColumn(x int) (err error) {
	if _, err = Output.WriteString(CursorColumnSeq + strconv.Itoa(x)); err != nil {
		return err
	}
	return nil
}
