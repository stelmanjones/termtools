package spin

import "iter"

// CharSet is a type alias for a slice of strings.
type CharSet = []string

// SpinnerVariant represents a variant of a spinner with a specific character set and interval.
type SpinnerVariant struct {
	chars    CharSet
	Interval int
}

func (v *SpinnerVariant) All() iter.Seq[string] {
	return func(yield func(string) bool) {
		for _, c := range v.chars {
			if !yield(c) {
				return
			}
		}
	}
}

// NewSpinnerVariant creates a new SpinnerVariant with the given character set and interval.
func NewSpinnerVariant(charSet CharSet, interval int) SpinnerVariant {
	return SpinnerVariant{chars: charSet, Interval: interval}
}

var (
	// GrowVertical is a spinner variant that grows the spinner vertically.
	GrowVertical = NewSpinnerVariant(CharSets[0], 80)
	// Bounce is a spinner variant that bounces the spinner.
	Bounce = NewSpinnerVariant(CharSets[1], 120)
	// Dots1 is a spinner variant that shows three dots.
	Dots1 = NewSpinnerVariant(CharSets[2], 80)
	// Dots2 is a spinner variant that shows three dots.
	Dots2 = NewSpinnerVariant(CharSets[3], 80)
	// Dots3 is a spinner variant that shows three dots.
	Dots3 = NewSpinnerVariant(CharSets[4], 80)
	// Letters is a spinner variant that shows the letters a-z.
	Letters = NewSpinnerVariant(CharSets[5], 120)
	// GrowHorizontal is a spinner variant that grows the spinner horizontally.
	GrowHorizontal = NewSpinnerVariant(CharSets[6], 80)
	// Simple is a spinner variant that shows a simple spinner.
	Simple = NewSpinnerVariant(CharSets[7], 120)
	// GrowHV is a spinner variant that grows the spinner horizontally and vertically.
	GrowHV = NewSpinnerVariant(CharSets[8], 80)
	// Arc is a spinner variant that shows an arc.
	Arc = NewSpinnerVariant(CharSets[9], 80)
	// BouncingBar is a spinner variant that shows a bouncing bar.
	BouncingBar = NewSpinnerVariant(CharSets[10], 80)
	// BouncingSimple is a spinner variant that shows a bouncing simple spinner.
	BouncingSimple = NewSpinnerVariant(CharSets[11], 80)
	// MovingDots is a spinner variant that shows moving dots.
	MovingDots = NewSpinnerVariant(CharSets[12], 80)
)

// CharSets contains the available character sets
var CharSets = map[int][]string{
	0: {"▁", "▃", "▄", "▅", "▆", "▇", "█", "▇", "▆", "▅", "▄", "▃", "▁"},
	1: {"▖", "▘", "▝", "▗"},
	2: {"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"},
	3: {"⠁", "⠂", "⠄", "⡀", "⢀", "⠠", "⠐", "⠈"},
	4: {"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"},
	5: {"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"},
	6: {"▉", "▊", "▋", "▌", "▍", "▎", "▏", "▎", "▍", "▌", "▋", "▊", "▉"},

	7:  {".  ", ".. ", "..."},
	8:  {"▁", "▂", "▃", "▄", "▅", "▆", "▇", "█", "▉", "▊", "▋", "▌", "▍", "▎", "▏", "▏", "▎", "▍", "▌", "▋", "▊", "▉", "█", "▇", "▆", "▅", "▄", "▃", "▂", "▁"},
	9:  {"◜", "◝", "◞", "◟"},
	10: {"(●    )", "( ●    )", "(  ●   )", "(   ●  )", "(    ● )", "(     ●)", "(    ● )", "(   ●  )", "(  ●   )", "( ●    )"},
	11: {".  ", ".. ", "...", " ..", "  .", "   "},
	12: {"∙∙∙", "●∙∙", "∙●∙", "∙∙●", "∙∙∙"},
}
