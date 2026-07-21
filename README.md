# gomorse

A fast, lightweight, and easy-to-use Go package for encoding text to Morse code and decoding Morse code back to text.

## Features

- **Bidirectional Conversion**: Easily encode plain text to Morse code and decode Morse code back to text.
- **Strict & Safe API**: Offers basic methods returning errors for robust error handling (`Encode`, `Decode`), as well as helper functions that panic on errors (`MustEncode`, `MustDecode`).
- **Validation**: Quickly validate if a given string contains valid Morse code sequences using `IsValid`.
- **Zero Dependencies**: Uses only the Go standard library.
- **Wide Character Support**: Fully supports uppercase and lowercase English letters (A-Z), numbers (0-9), and common punctuation/symbols.

---

## Installation

To use `gomorse` in your Go project, install it using `go get`:

```bash
go get github.com/mdhsaikats/gomorse
```

---

## Usage

### Import the Package

```go
import "github.com/mdhsaikats/gomorse"
```

### 1. Encoding Text to Morse Code

The package automatically converts lowercase characters to uppercase before encoding. Words are separated by ` / ` (space, forward slash, space), and letters within a word are separated by a single space.

```go
package main

import (
	"fmt"
	"log"

	"github.com/mdhsaikats/gomorse"
)

func main() {
	// Encode text
	code, err := morse.Encode("Hello World")
	if err != nil {
		log.Fatalf("failed to encode: %v", err)
	}
	fmt.Println(code)
	// Output: .... . .-.. .-.. --- / .-- --- .-. .-.. -..
}
```

### 2. Decoding Morse Code to Text

The decoder interprets ` / ` as a space separator between words and a single space as a separator between individual Morse symbols.

```go
package main

import (
	"fmt"
	"log"

	"github.com/mdhsaikats/gomorse"
)

func main() {
	// Decode Morse code
	text, err := morse.Decode(".... . .-.. .-.. --- / .-- --- .-. .-.. -..")
	if err != nil {
		log.Fatalf("failed to decode: %v", err)
	}
	fmt.Println(text)
	// Output: HELLO WORLD
}
```

### 3. Verification & Panic Variants

If you know your inputs are safe or prefer to handle failures with standard panics, you can use the `Must` wrappers:

```go
// Panics if the input contains unsupported characters
code := morse.MustEncode("HELLO")

// Panics if the input contains invalid Morse code sequences
text := morse.MustDecode(".... . .-.. .-.. ---")
```

To check if a string is a valid Morse sequence without raising errors or panics:

```go
isValid := morse.IsValid(".... . .-.. .-.. ---") // true
isInvalid := morse.IsValid("........")            // false (contains invalid sequence)
```

---

## API Reference

### Functions

#### [Encode](file:///w:/GoLang/gomorse/encoder.go#L5)
```go
func Encode(text string) (string, error)
```
Encodes plain text to Morse code. Returns `ErrUnsupportedCharacter` if a character cannot be mapped.

#### [Decode](file:///w:/GoLang/gomorse/decoder.go#L7)
```go
func Decode(code string) (string, error)
```
Decodes Morse code back to text (in uppercase). Returns `ErrInvalidMorseCode` if a sequence is invalid.

#### [MustEncode](file:///w:/GoLang/gomorse/morse.go#L15)
```go
func MustEncode(text string) string
```
Like `Encode`, but panics if an error is encountered.

#### [MustDecode](file:///w:/GoLang/gomorse/morse.go#L28)
```go
func MustDecode(code string) string
```
Like `Decode`, but panics if an error is encountered.

#### [IsValid](file:///w:/GoLang/gomorse/morse.go#L6)
```go
func IsValid(code string) bool
```
Returns `true` if `code` is a valid Morse sequence, and `false` otherwise.

### Errors

- `morse.ErrUnsupportedCharacter`: Returned when an input string contains characters not supported by the conversion table.
- `morse.ErrInvalidMorseCode`: Returned when decoding a sequence that is not recognized as a valid Morse character.

---

## Supported Characters

| Category | Supported Characters |
| :--- | :--- |
| **Letters** | `A` through `Z` (case-insensitive during encoding) |
| **Numbers** | `0` through `9` |
| **Punctuation** | `.`, `,`, `?`, `!`, `/`, `(`, `)`, `&`, `:`, `;`, `=`, `+`, `-`, `_`, `"`, `$`, `@` |

---

## Running the Command-Line Demo

The package includes a basic CLI command demo under the `cmd/gomorse` directory. You can run it directly:

```bash
go run github.com/mdhsaikats/gomorse/cmd/gomorse
```

---

## License

This project is licensed under the MIT License - see the [LICENSE](file:///w:/GoLang/gomorse/LICENSE) file for details.
