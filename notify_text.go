package systray

import "unicode/utf16"

// fitUTF16 encodes s as UTF-16 for a fixed field of size units, NUL
// included, cutting whatever does not fit at a character boundary.
func fitUTF16(s string, size int) []uint16 {
	out := make([]uint16, 0, size)
	for _, r := range s {
		enc := utf16.Encode([]rune{r})
		if len(out)+len(enc) > size-1 {
			break
		}
		out = append(out, enc...)
	}
	return append(out, 0)
}
