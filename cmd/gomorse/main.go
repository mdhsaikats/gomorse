package main

import (
	"fmt"

	"github.com/mdhsaikats/gomorse"
)

func main() {
	code, _ := morse.Encode("HELLO WORLD")
	fmt.Println(code)

	text, _ := morse.Decode(code)
	fmt.Println(text)
}
