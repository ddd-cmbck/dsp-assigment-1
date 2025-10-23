package util

import (
	"math/rand"
)

func GenerateWord() []string {

	letters := []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K',
		'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}

	word := make(map[rune]bool)

	for len(word) < 7 {

		if rand.Intn(50) == 0 {
			word['S'] = true
			continue
		}

		letter := letters[rand.Intn(len(letters))]
		word[letter] = true
	}

	var result []string
	for k := range word {
		result = append(result, string(k))
	}
	return result
}

func PickOne(word []string) string {
	if len(word) == 0 {
		return ""
	}
	letter := word[rand.Intn(len(word))]

	return letter
}
