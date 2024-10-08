package text

import (
	"fmt"
	"strings"

	"github.com/mattn/go-runewidth"

	"github.com/gookit/color"
)

// VisibleLength returns the length of the string as seen by a human.
// It removes all ANSI sequences from the string.
func VisibleLength(s ...interface{}) int {
	return runewidth.StringWidth(stripSequences(s...))
}

// stripSequences removes all ANSI sequences from the string.
func stripSequences(a ...interface{}) string {
	return color.ClearCode(fmt.Sprint(a...))
}

// mapLines runs function fn on every line of the string.
func mapLines(str string, fn func(string) string) string {
	return splitAndMap(str, "\n", fn)
}

// splitAndMap splits the string and runs the function fn on every line.
func splitAndMap(str string, split string, fn func(string) string) string {
	arr := strings.Split(str, split)
	for i := 0; i < len(arr); i++ {
		arr[i] = fn(arr[i])
	}
	return strings.Join(arr, split)
}

// AlignLeft aligns string to the left.
func AlignLeft(str string) string {
	return mapLines(str, func(line string) string {
		return strings.TrimLeft(line, " ")
	})
}

// AlignRight aligns string to the right.
func AlignRight(str string, width int) string {
	return mapLines(str, func(line string) string {
		line = strings.Trim(line, " ")
		return strings.Repeat(" ", width-VisibleLength(line)) + line
	})
}

// AlignTop aligns the text to the top.
func AlignTop(str string, width, height int) string {
	var result strings.Builder
	emptyLine := strings.Repeat(" ", width)
	lineCount := LineCount(str)
	lines := SplitLines(str)

	for _, s := range lines {
		result.WriteString(s + "\n")
	}
	for i := 0; i < height-lineCount; i++ {
		result.WriteString(emptyLine + "\n")
	}
	return result.String()
}

// AlignMiddle aligns the text to the middle.
func AlignMiddle(str string, width, height int) string {
	var result strings.Builder
	lineCount := LineCount(str)
	emptyLine := strings.Repeat(" ", width)
	lines := SplitLines(str)
	padding := (height - lineCount) / 2
	for i := 0; i < padding; i++ {
		result.WriteString(emptyLine + "\n")
	}
	for _, s := range lines {
		result.WriteString(s + "\n")
	}
	for i := 0; i < padding; i++ {
		result.WriteString(emptyLine + "\n")
	}
	return result.String()
}

// AlignBottom aligns the text to the bottom.
func AlignBottom(str string, width, height int) string {
	var result strings.Builder
	lineCount := LineCount(str)
	emptyLine := strings.Repeat(" ", width)
	lines := SplitLines(str)
	for i := 0; i < height-lineCount; i++ {
		result.WriteString(emptyLine + "\n")
	}
	for _, s := range lines {
		result.WriteString(s + "\n")
	}
	return result.String()
}

// AlignCenter centers str. It trims and then centers all the lines in the text with space.
func AlignCenter(str string, width int) string {
	return mapLines(str, func(line string) string {
		line = strings.Trim(line, " ")
		return CenterText(line, width)
	})
}

// CenterText centers the text by adding spaces to the left and right.
// It assumes the text is one line. For multiple lines use AlignCenter.
func CenterText(str string, width int) string {
	return strings.Repeat(" ", (width-VisibleLength(str))/2) + str + strings.Repeat(" ", (width-VisibleLength(str))/2)
}

// SplitLines splits the string by new line character ("\n")
func SplitLines(s ...interface{}) []string {
	return strings.Split(fmt.Sprint(s...), "\n")
}

// LineCount returns the number of lines in the string.
func LineCount(s ...interface{}) int {
	return len(SplitLines(s...))
}
