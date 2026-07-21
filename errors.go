package morse

import (
	"errors"
)

var (
	ErrUnsupportedCharacter = errors.New("unsupported character")
	ErrInvalidMorseCode     = errors.New("invalid morse code")
)
