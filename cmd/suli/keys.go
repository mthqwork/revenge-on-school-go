package main

import (
	"fmt"
	"strings"

	"tclinux-suli/internal/charset"
)

// Scan codes for extended keys (ReadKey returns 0 followed by these).
const (
	scUp    = 0x48
	scDown  = 0x50
	scLeft  = 0x4b
	scRight = 0x4d
	scPgUp  = 0x49
	scPgDn  = 0x51
	scHome  = 0x47
	scEnd   = 0x4f
	scIns   = 0x52
	scDel   = 0x53
)

// altScan maps letters to their scan codes for Alt+letter combinations.
var altScan = map[rune]byte{
	'q': 16, 'w': 17, 'e': 18, 'r': 19, 't': 20, 'y': 21, 'u': 22, 'i': 23, 'o': 24, 'p': 25,
	'a': 30, 's': 31, 'd': 32, 'f': 33, 'g': 34, 'h': 35, 'j': 36, 'k': 37, 'l': 38,
	'z': 44, 'x': 45, 'c': 46, 'v': 47, 'b': 48, 'n': 49, 'm': 50,
}

// namedKeys maps script key names to ReadKey byte sequences.
var namedKeys = map[string][]byte{
	"enter": {13}, "return": {13}, "esc": {27}, "escape": {27}, "space": {' '},
	"backspace": {8}, "tab": {9},
	"up": {0, scUp}, "down": {0, scDown}, "left": {0, scLeft}, "right": {0, scRight},
	"pgup": {0, scPgUp}, "pgdn": {0, scPgDn}, "home": {0, scHome}, "end": {0, scEnd},
	"ins": {0, scIns}, "del": {0, scDel},
}

// scriptItem is one step of a headless input script.
type scriptItem struct {
	keys []byte
	shot string
}

// parseScript parses a headless input script. Items are separated by
// spaces or commas:
//
//	Enter Esc Space Up Down PgUp PgDn Del ...   named keys
//	Alt-s                                       Alt + letter
//	t:text                                      type the characters of text
//	x3:Down                                     repeat an item
//	shot:file.png                               save the screen (when the game waits for input)
func parseScript(s string) ([]scriptItem, error) {
	var out []scriptItem
	for _, tok := range strings.FieldsFunc(s, func(r rune) bool { return r == ' ' || r == ',' || r == '\n' }) {
		rep := 1
		if strings.HasPrefix(tok, "x") && strings.Contains(tok, ":") {
			if _, err := fmt.Sscanf(tok, "x%d:", &rep); err == nil {
				tok = tok[strings.Index(tok, ":")+1:]
			}
		}
		var it scriptItem
		lt := strings.ToLower(tok)
		switch {
		case strings.HasPrefix(lt, "shot:"):
			it.shot = tok[5:]
		case strings.HasPrefix(lt, "t:"):
			it.keys = charset.Encode(tok[2:])
		case strings.HasPrefix(lt, "alt-") && len([]rune(lt)) == 5:
			sc, ok := altScan[[]rune(lt)[4]]
			if !ok {
				return nil, fmt.Errorf("unknown key %q", tok)
			}
			it.keys = []byte{0, sc}
		default:
			k, ok := namedKeys[lt]
			if !ok {
				if len([]rune(tok)) == 1 {
					k = charset.Encode(tok)
				} else {
					return nil, fmt.Errorf("unknown key %q", tok)
				}
			}
			it.keys = k
		}
		for ; rep > 0; rep-- {
			out = append(out, it)
		}
	}
	return out, nil
}
