package game

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"

	"tclinux-suli/assets"
	"tclinux-suli/internal/charset"
	"tclinux-suli/internal/dos"
)

func isSaveName(name string) bool {
	return strings.EqualFold(filepath.Base(strings.ReplaceAll(name, "\\", "/")), "savegame.tsi")
}

// readFile resolves a DOS file name of the original program.
func (g *Game) readFile(name string) ([]byte, error) {
	if isSaveName(name) {
		return os.ReadFile(g.SavePath)
	}
	return assets.ReadData(name)
}

// AssignText binds a text file variable to a name (Assign).
func (g *Game) AssignText(name string) *dos.TextFile {
	return &dos.TextFile{Name: name}
}

// ResetText opens a text file for reading ({$I-} Reset; see IOResult).
func (g *Game) ResetText(f *dos.TextFile) {
	b, err := g.readFile(f.Name)
	if err != nil {
		g.ioResult = 2
		f.SetContent(nil)
		return
	}
	g.ioResult = 0
	f.SetContent(b)
}

// RewriteText opens a text file for writing.
func (g *Game) RewriteText(f *dos.TextFile) {
	f.W = &bytes.Buffer{}
	g.ioResult = 0
}

// CloseText closes a text file, writing it out if it was rewritten.
func (g *Game) CloseText(f *dos.TextFile) {
	if f.W == nil {
		return
	}
	data := f.W.Bytes()
	f.W = nil
	if isSaveName(f.Name) {
		if err := os.MkdirAll(filepath.Dir(g.SavePath), 0o755); err == nil {
			if err := os.WriteFile(g.SavePath, data, 0o644); err != nil {
				g.ioResult = 5
			}
		}
	}
}

// IOResult returns and clears the last I/O error code.
func (g *Game) IOResult() int {
	r := g.ioResult
	g.ioResult = 0
	return r
}

func (g *Game) write(f *dos.TextFile, s string) {
	if f == g.Output || f.W == nil {
		g.m.Console.WriteString(s)
		return
	}
	f.W.Write(charset.Encode(s))
}

func (g *Game) writeln(f *dos.TextFile) {
	if f == g.Output || f.W == nil {
		g.m.Console.WriteString("\n")
		return
	}
	f.W.WriteString("\r\n")
}

// FindFirst only checks for existence; DosError is 0 if the file exists.
func (g *Game) FindFirst(path string) {
	if _, err := g.readFile(path); err != nil {
		if strings.EqualFold(path, "suli.exe") {
			g.DosError = 0
			return
		}
		g.DosError = 18
		return
	}
	g.DosError = 0
}

// DiskError (1e9b:00fa) - fatal read error.
func (g *Game) DiskError() {
	g.CloseGraph()
	g.writeln(g.Output)
	g.write(g.Output, " Kurva nagy hiba lemezolvasás közben.")
	g.writeln(g.Output)
	g.writeln(g.Output)
	g.Halt()
}
