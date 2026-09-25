// Package charset converts between the game's 8-bit text encoding and UTF-8.
//
// The original data and program use code page 437 with a few CWI-2
// (Hungarian) positions: 0x8F = Á, 0x93 = ő, 0x96 = ű. Text is rendered by
// byte value through the VGA ROM font, so strings are kept as UTF-8 in Go and
// encoded back to bytes only for drawing and for writing the save file.
package charset

import "strings"

var high = [128]rune{
	'Ç', 'ü', 'é', 'â', 'ä', 'à', 'å', 'ç', 'ê', 'ë', 'è', 'ï', 'î', 'ì', 'Ä', 'Á',
	'É', 'æ', 'Æ', 'ő', 'ö', 'ò', 'ű', 'ù', 'ÿ', 'Ö', 'Ü', '¢', '£', '¥', '₧', 'ƒ',
	'á', 'í', 'ó', 'ú', 'ñ', 'Ñ', 'ª', 'º', '¿', '⌐', '¬', '½', '¼', '¡', '«', '»',
	'░', '▒', '▓', '│', '┤', '╡', '╢', '╖', '╕', '╣', '║', '╗', '╝', '╜', '╛', '┐',
	'└', '┴', '┬', '├', '─', '┼', '╞', '╟', '╚', '╔', '╩', '╦', '╠', '═', '╬', '╧',
	'╨', '╤', '╥', '╙', '╘', '╒', '╓', '╫', '╪', '┘', '┌', '█', '▄', '▌', '▐', '▀',
	'α', 'ß', 'Γ', 'π', 'Σ', 'σ', 'µ', 'τ', 'Φ', 'Θ', 'Ω', 'δ', '∞', 'φ', 'ε', '∩',
	'≡', '±', '≥', '≤', '⌠', '⌡', '÷', '≈', '°', '∙', '·', '√', 'ⁿ', '²', '■', ' ',
}

var rev = func() map[rune]byte {
	m := make(map[rune]byte, 128)
	for i, r := range high {
		m[r] = byte(0x80 + i)
	}
	return m
}()

// Decode converts game bytes to a UTF-8 string.
func Decode(b []byte) string {
	var sb strings.Builder
	sb.Grow(len(b))
	for _, c := range b {
		if c < 0x80 {
			sb.WriteByte(c)
		} else {
			sb.WriteRune(high[c-0x80])
		}
	}
	return sb.String()
}

// Encode converts a UTF-8 string to game bytes. Unknown runes become '?'.
func Encode(s string) []byte {
	out := make([]byte, 0, len(s))
	for _, r := range s {
		switch {
		case r < 0x80:
			out = append(out, byte(r))
		default:
			if c, ok := rev[r]; ok {
				out = append(out, c)
			} else {
				out = append(out, '?')
			}
		}
	}
	return out
}

// Len returns the length of s in game bytes (Pascal Length()).
func Len(s string) int { return len([]rune(s)) }
