package systray

import (
	"testing"
	"unicode/utf16"
)

// Windows takes a balloon's text in fixed UTF-16 fields; what does not fit
// is cut, never in the middle of a character, and the field always ends in
// a NUL.
func TestFitUTF16CutsAtACharacterAndEndsInNUL(t *testing.T) {
	for _, tc := range []struct {
		in   string
		size int
		want string
	}{
		{"hello", 64, "hello"},
		{"", 8, ""},
		{"abcdefgh", 4, "abc"},
		{"ação", 64, "ação"},
		{"ab😀cd", 4, "ab"},  // the emoji takes two units, which would leave no room for the NUL
		{"ab😀cd", 5, "ab😀"}, // it fits exactly
	} {
		got := fitUTF16(tc.in, tc.size)
		if len(got) > tc.size || got[len(got)-1] != 0 {
			t.Fatalf("fitUTF16(%q, %d) = %v: longer than the field, or not NUL-terminated", tc.in, tc.size, got)
		}
		if s := string(utf16.Decode(got[:len(got)-1])); s != tc.want {
			t.Errorf("fitUTF16(%q, %d) = %q, want %q", tc.in, tc.size, s, tc.want)
		}
	}
}
