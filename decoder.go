package morse

import (
	"strings"
)

func Decode(code string) (string, error) {

	var result strings.Builder

	words := strings.Split(code, " / ")

	for wordIndex, word := range words {

		letters := strings.Split(word, " ")

		for _, letter := range letters {

			if letter == "" {
				continue
			}

			char, exists := morseToChar[letter]

			if !exists {
				return "", ErrInvalidMorseCode
			}

			result.WriteRune(char)
		}

		if wordIndex != len(words)-1 {
			result.WriteRune(' ')
		}
	}

	return result.String(), nil
}
