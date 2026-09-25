// Package assets embeds the original game data and the VGA 8x8 font.
package assets

import (
	"embed"
	"io/fs"
	"strings"
)

//go:embed data
var data embed.FS

//go:embed font8x8.bin
var Font8x8 []byte

// ReadData returns a file from the original DATA directory. Names are case
// insensitive and may carry a DOS style "data\" prefix.
func ReadData(name string) ([]byte, error) {
	name = strings.ToLower(strings.ReplaceAll(name, "\\", "/"))
	name = strings.TrimPrefix(name, "data/")
	return fs.ReadFile(data, "data/"+name)
}
