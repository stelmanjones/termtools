package styles

import (
	"math"
	"math/rand"
	"strings"
)

var glitchChars = []rune{'@', '#', '$', '%', '!', '&', '*', '(', ')', '[', ']', '{', '}', '<', '>', '?'}

func Glitch(str string) string {
	inputLength := len(str)
	chunkSize := int(math.Max(1, math.Round(float64(inputLength)*0.02)))
	chunks := []string{}

	for i := 0; i < inputLength; i++ {
		currentChar := str[i]
		skip := int(math.Round(math.Max(0, (rand.Float64()-0.8)*float64(chunkSize))))
		chunk := strings.Map(func(r rune) rune {
			if r != '\r' && r != '\n' {
				return ' '
			}
			return r
		}, str[i:i+skip])
		chunks = append(chunks, chunk)
		i += skip
		if currentChar != 0 && rand.Float64() > 0.900 {
			chunks = append(chunks, string(glitchChars[rand.Intn(len(glitchChars))]))
		} else if rand.Float64() > 0.005 {
			chunks = append(chunks, string(currentChar))
		}
	}

	result := strings.Join(chunks, "")
	if rand.Float64() > 0.99 {
		result = strings.ToUpper(result)
	} else if rand.Float64() < 0.01 {
		result = strings.ToLower(result)
	}

	return result
}
