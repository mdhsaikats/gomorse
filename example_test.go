package morse_test

import (
	"fmt"

	"github.com/mdhsaikats/gomorse"
)

func ExampleEncode() {
	code, err := morse.Encode("HELLO WORLD")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(code)
	// Output: .... . .-.. .-.. --- / .-- --- .-. .-.. -..
}

func ExampleDecode() {
	text, err := morse.Decode(".... . .-.. .-.. --- / .-- --- .-. .-.. -..")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(text)
	// Output: HELLO WORLD
}
