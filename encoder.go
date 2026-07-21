package morse

import "strings"

func Encode(text string) (string, error) {
	var result []string

	for _, char := range strings.ToUpper(text) {
		if char == ' ' {
			result = append(result, "/")
			continue
		}

		code, exists := charToMorse[char]

		if !exists {
			return "", ErrUnsupportedCharacter
		}

		result = append(result, code)
	}
	return strings.Join(result, " "), nil
}
