package spin

import "time"

// Copyright (c) 2024 Oscar Nordmar
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

type CharSet = []string
type SpinnerVariant struct {
	CharSet
	Interval time.Duration
}

func NewSpinnerVariant(charSet CharSet, interval time.Duration) SpinnerVariant {
	return SpinnerVariant{CharSet: charSet, Interval: interval}
}

var (
	GrowVertical   = NewSpinnerVariant(CharSets[0], 80)
	Bounce         = NewSpinnerVariant(CharSets[1], 120)
	Dots1          = NewSpinnerVariant(CharSets[2], 80)
	Dots2          = NewSpinnerVariant(CharSets[3], 80)
	Dots3          = NewSpinnerVariant(CharSets[4], 80)
	Letters        = NewSpinnerVariant(CharSets[5], 120)
	GrowHorizontal = NewSpinnerVariant(CharSets[6], 80)
	Simple         = NewSpinnerVariant(CharSets[7], 120)
	GrowHV         = NewSpinnerVariant(CharSets[8], 80)
	Arc            = NewSpinnerVariant(CharSets[9], 80)
	BouncingBar    = NewSpinnerVariant(CharSets[10], 80)
	BouncingSimple = NewSpinnerVariant(CharSets[11], 80)
	MovingDots     = NewSpinnerVariant(CharSets[12], 80)
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
	10: {"( ●    )", "(  ●   )", "(   ●  )", "(    ● )", "(     ●)", "(    ● )", "(   ●  )", "(  ●   )", "( ●    )"},
	11: {".  ", ".. ", "...", " ..", "  .", "   "},
	12: {"∙∙∙", "●∙∙", "∙●∙", "∙∙●", "∙∙∙"},
}

var CharSetNames = []string{"growVertical", "bounce", "dots1", "dots2", "dots3", "letters", "growHorizontal", "simple", "growHV", "arc", "bouncingBar", "bouncingSimple", "movingDots"}
