package util

import (
	"fmt"
	"strings"
)

func RenderWord(word []string, center string) ([]string, error) {
	length := len(word)
	if length%2 == 0 {
		return nil, fmt.Errorf("word length must be odd")
	}

	centerIndex := length / 2

	letterIndex := -1
	for i, ch := range word {
		if strings.EqualFold(ch, center) {
			letterIndex = i
			break
		}
	}

	// Swap the letter into the center if necessary
	if letterIndex != -1 && letterIndex != centerIndex {
		word[letterIndex], word[centerIndex] = word[centerIndex], word[letterIndex]
	}

	// Build the formatted string
	result := make([]string, length)
	for i, ch := range word {
		upper := strings.ToUpper(ch)
		if i == centerIndex {
			result[i] = fmt.Sprintf("[%s]", upper)
		} else {
			result[i] = upper
		}
	}

	return result, nil
}
