// Package morse provides functions for encoding
// and decoding International Morse code.
package morse

// IsValid checks if a Morse sequence is valid.
func IsValid(code string) bool {

	_, err := Decode(code)

	return err == nil
}

// MustEncode converts text to Morse code.
// It panics if the text contains unsupported characters.
func MustEncode(text string) string {

	result, err := Encode(text)

	if err != nil {
		panic(err)
	}

	return result
}

// MustDecode converts Morse code to text.
// It panics if the Morse code is invalid.
func MustDecode(code string) string {

	result, err := Decode(code)

	if err != nil {
		panic(err)
	}

	return result
}
