package morse

import (
	"testing"
)

func BenchmarkEncode(b *testing.B) {
	input := "THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Encode(input)
	}
}

func BenchmarkDecode(b *testing.B) {
	input := "- .... . / --.- ..- .. -.-. -.- / -... .-. --- .-- -. / ..-. --- -..- / .--- ..- -- .--. ... / --- ...- . .-. / - .... . / .-.. .- --.. -.-- / -.. --- --."
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Decode(input)
	}
}
