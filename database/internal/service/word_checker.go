package service

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
	"sync"
)

type WordChecker struct {
	words map[string]int
}

var (
	instance *WordChecker
	once     sync.Once
)

func Load(filePath string) (*WordChecker, error) {
	var err error
	once.Do(func() {
		data, e := os.ReadFile(filePath)
		if e != nil {
			err = e
			return
		}

		var raw map[string]int
		if e = json.Unmarshal(data, &raw); e != nil {
			err = e
			return
		}

		instance = &WordChecker{words: raw}
	})

	if err != nil {
		return nil, err
	}

	return instance, nil
}

func GetInstance() (*WordChecker, error) {

	if instance == nil {
		return nil, errors.New("dictionary not initialised")
	}

	return instance, nil
}

func (s *WordChecker) CheckWord(word string) (bool, error) {

	if len(word) == 0 {
		return false, errors.New("word is too small")
	}

	_, exists := s.words[strings.ToLower(word)]

	if exists {
		return true, nil
	}

	return false, nil

}
