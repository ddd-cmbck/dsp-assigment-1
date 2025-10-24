package service

import (
	"errors"
	"flag"
	"strings"
)

var BONUS = flag.Int("bonus", 7, "bonus which applied in case user used all 7 letters")

func EvalScore(word string, letters []string) (int32, error) {
	word = strings.ToLower(word)

	if len(word) < 4 {
		return 0, errors.New("word size too small")
	}

	if len(letters) != 7 {
		return 0, errors.New("letters size must be 7")
	}

	if len(word) == 4 {
		return 1, nil
	}

	score := int32(len(word))

	pangramCount := countFullPangramSets(word, letters)

	if pangramCount > 0 {
		score += int32(*BONUS) * int32(pangramCount)
	}

	return score, nil
}

func countFullPangramSets(word string, letters []string) int {

	freq := make(map[rune]int)
	for _, r := range word {
		for _, l := range letters {
			if r == []rune(strings.ToLower(l))[0] {
				freq[r]++
				break
			}
		}
	}

	minCount := -1
	for _, l := range letters {
		r := []rune(strings.ToLower(l))[0]
		count := freq[r]
		if count == 0 {
			return 0
		}
		if minCount == -1 || count < minCount {
			minCount = count
		}
	}

	return minCount
}
