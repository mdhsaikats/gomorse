package morse

import (
	"errors"
	"testing"
)

func TestEncode(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr error
	}{
		{
			name:    "Standard word",
			input:   "HELLO",
			want:    ".... . .-.. .-.. ---",
			wantErr: nil,
		},
		{
			name:    "Multiple words with spaces",
			input:   "HELLO WORLD",
			want:    ".... . .-.. .-.. --- / .-- --- .-. .-.. -..",
			wantErr: nil,
		},
		{
			name:    "Lowercase input",
			input:   "hello",
			want:    ".... . .-.. .-.. ---",
			wantErr: nil,
		},
		{
			name:    "Numbers and punctuation",
			input:   "SOS 100?",
			want:    "... --- ... / .---- ----- ----- ..--..",
			wantErr: nil,
		},
		{
			name:    "Unsupported character",
			input:   "HELLO#",
			want:    "",
			wantErr: ErrUnsupportedCharacter,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Encode(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Encode() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr error
	}{
		{
			name:    "Standard word",
			input:   ".... . .-.. .-.. ---",
			want:    "HELLO",
			wantErr: nil,
		},
		{
			name:    "Multiple words",
			input:   ".... . .-.. .-.. --- / .-- --- .-. .-.. -..",
			want:    "HELLO WORLD",
			wantErr: nil,
		},
		{
			name:    "Extra spaces and empty letters",
			input:   ".... .  .-.. .-.. ---", // two spaces between . and -
			want:    "HELLO",
			wantErr: nil,
		},
		{
			name:    "Invalid Morse code sequence",
			input:   "........",
			want:    "",
			wantErr: ErrInvalidMorseCode,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decode(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Decode() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestIsValid(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "Valid Morse",
			input: ".... . .-.. .-.. ---",
			want:  true,
		},
		{
			name:  "Invalid Morse",
			input: "........",
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValid(tt.input); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMustEncode(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("MustEncode panicked unexpectedly: %v", r)
			}
		}()
		got := MustEncode("HELLO")
		want := ".... . .-.. .-.. ---"
		if got != want {
			t.Errorf("MustEncode() = %q, want %q", got, want)
		}
	})

	t.Run("Panics on invalid", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("MustEncode did not panic on invalid input")
			}
		}()
		MustEncode("HELLO#")
	})
}

func TestMustDecode(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("MustDecode panicked unexpectedly: %v", r)
			}
		}()
		got := MustDecode(".... . .-.. .-.. ---")
		want := "HELLO"
		if got != want {
			t.Errorf("MustDecode() = %q, want %q", got, want)
		}
	})

	t.Run("Panics on invalid", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("MustDecode did not panic on invalid input")
			}
		}()
		MustDecode("........")
	})
}
