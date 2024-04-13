package spin

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

var (
	GrowVertical   CharSet = CharSets[0]
	Bounce         CharSet = CharSets[1]
	Dots1          CharSet = CharSets[2]
	Dots2          CharSet = CharSets[3]
	Dots3          CharSet = CharSets[4]
	Letters        CharSet = CharSets[5]
	GrowHorizontal CharSet = CharSets[6]
	Simple         CharSet = CharSets[7]
	GrowHV         CharSet = CharSets[8]
	Arc            CharSet = CharSets[9]
	BouncingBar    CharSet = CharSets[10]
	BouncingSimple CharSet = CharSets[11]
	MovingDots     CharSet = CharSets[12]
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
