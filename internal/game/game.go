// Package game is the Go port of the game logic of Revenge on School.
//
// The code mirrors the structure of the original Turbo Pascal program: one
// big set of global variables (fields of Game) and procedures operating on
// them. Much of it was translated semi-automatically from the disassembly
// (with a custom Turbo Pascal to Go translator), so the control flow deliberately stays close to the
// original, including its quirks.
package game

import (
	"math/rand"
	"strconv"

	"tclinux-suli/internal/bgi"
	"tclinux-suli/internal/charset"
	"tclinux-suli/internal/dos"
)

type character struct {
	name  string // string[25]
	hp    int    // damage taken; >= 5 means dead / knocked out
	state int    // mood / how annoyed the person is
}

type move struct {
	label string // string[26]
	room  int
}

// Game holds the complete state of the original program.
type Game struct {
	m   *dos.Machine
	scr *bgi.Screen
	rnd *rand.Rand

	// SavePath is where Alt-S writes and "loadgame" reads the save file.
	SavePath string
	// Args are the command line parameters (ParamStr(1..)).
	Args []string

	// Globals, named after their role; the DS offset is in brackets.
	key      int // [0864] last ReadKey
	v0866    int // [0866]
	tgtState int // [0868]
	tgtHP    int // [086a]
	target   int // [086c] index into chars, 0 = none
	j        int // [086e]
	k        int // [0870]
	energy   int // [0872]
	room     int // [0874]
	n        int // [0876]
	m2       int // [0878]
	i        int // [087c] global loop variable, shared by many procedures!

	txt     *dos.TextFile // [0910] the one text file variable
	tgtName string        // [0a12] string[25]
	s0a2c   string        // [0a2c]
	line    string        // [0a46]
	line2   string        // [0b46]

	roomPeople [11]string    // [0d2c] array[1..10] of string[25]
	chars      [31]character // [0e2c] array[1..30]
	// The following string[25] arrays lie back to back in the original
	// data segment. The program sometimes indexes one past the end (e.g.
	// inv[13] is taken[1]), so they are slices of one backing array with
	// 1-based indexing, and element 0 of each is the last element of the
	// previous array, exactly like in memory.
	slots     [200]string
	roomItems []string // [11b4] items to be found here, [1..5]
	itemMenu  []string // [1236] compacted inventory shown in the menu, [1..12]
	inv       []string // [136e] inventory, [1..12]
	taken     []string // [14a6] every item ever picked up (no second pickup), [1..120]
	localMenu []string // [20d6] room specific actions, [1..8]
	// moves is array[1..6] in the original, but EnterRoom clears 8
	// entries, overflowing into a2259. The port keeps them separate.
	moves [9]move // [21a3]
	// curse is array[1..150]; SaySomething also reads curse[151], which in the
	// original is never-written (zero) memory after the array, i.e. "".
	curse [152]string // [2259] words of the curse being composed, string[20]

	loaded2  int // [2ebc]
	flagD    int // [2ebd]
	flagE    int // [2ebe]
	flagF    int // [2ebf]
	flag0    int // [2ec0]
	loaded   int // [2ec1]
	DosError int // [36ea]

	Output *dos.TextFile // [37ec] stands for the DOS console

	targetPal [256][3]byte // [2ef2] palette unit: fade-in target
	ioResult  int
	save      []byte // pending save file contents while writing
}

// New creates the game bound to a machine.
func New(m *dos.Machine, seed int64) *Game {
	g := &Game{m: m, scr: m.Scr, rnd: rand.New(rand.NewSource(seed)), Output: &dos.TextFile{}}
	// slots[0] is the element "0" of roomItems (index 0 is never used).
	g.roomItems = g.slots[0:]
	g.itemMenu = g.slots[5:]
	g.inv = g.slots[17:]
	g.taken = g.slots[29:]
	g.localMenu = g.slots[149:]
	return g
}

// ---------------------------------------------------------------- runtime

func itoa(v int) string { return strconv.Itoa(v) }

func itoaw(v, w int) string {
	s := strconv.Itoa(v)
	for len(s) < w {
		s = " " + s
	}
	return s
}

func trunc(s string, max int) string { return dos.Trunc(s, max) }

// val is Pascal Val(s, v, code) returning v.
func val(s string, code *int) int {
	v, c := dos.Val(s)
	*code = c
	return v
}

func (g *Game) Halt()            { g.m.Halt() }
func (g *Game) KeyPressed() bool { return g.m.KeyPressed() }
func (g *Game) ReadKey() int     { return int(g.m.ReadKey()) }
func (g *Game) Delay(ms int)     { g.m.Delay(ms) }
func (g *Game) Random(n int) int {
	if n <= 0 {
		return 0
	}
	return g.rnd.Intn(n)
}
func (g *Game) Randomize() {}

// WaitKey (unit at 2714:0000): wait for a key, then drain the buffer.
func (g *Game) WaitKey() {
	for !g.KeyPressed() {
	}
	for g.KeyPressed() {
		g.ReadKey()
	}
}

// FlushKeys (2714:0022): drain the keyboard buffer.
func (g *Game) FlushKeys() {
	for g.KeyPressed() {
		g.ReadKey()
	}
}

// ParamStr returns the n-th command line argument like Turbo Pascal.
func (g *Game) ParamStr(n int) string {
	if n >= 1 && n <= len(g.Args) {
		return g.Args[n-1]
	}
	return ""
}

// ---------------------------------------------------------------- graphics

func (g *Game) OutTextXY(x, y int, s string)         { g.scr.OutTextXY(x, y, s) }
func (g *Game) SetColor(c int)                       { g.scr.SetColor(c) }
func (g *Game) SetFillStyle(p, c int)                { g.scr.SetFillStyle(p, c) }
func (g *Game) Bar(x1, y1, x2, y2 int)               { g.scr.Bar(x1, y1, x2, y2) }
func (g *Game) Line(x1, y1, x2, y2 int)              { g.scr.Line(x1, y1, x2, y2) }
func (g *Game) FloodFill(x, y, border int)           { g.scr.FloodFill(x, y, border) }
func (g *Game) SetViewPort(x1, y1, x2, y2, clip int) { g.scr.SetViewPort(x1, y1, x2, y2, clip != 0) }
func (g *Game) ClearViewPort()                       { g.scr.ClearViewPort() }
func (g *Game) ClearDevice()                         { g.scr.ClearDevice() }
func (g *Game) GetMaxX() int                         { return g.scr.GetMaxX() }
func (g *Game) GetMaxY() int                         { return g.scr.GetMaxY() }
func (g *Game) TextWidth(s string) int               { return g.scr.TextWidth(s) }

// CloseGraph leaves graphics mode: the front end switches to the console.
func (g *Game) CloseGraph() {}

// charAt returns the byte (game encoding) at 1-based position i of s, like
// s[i] in Pascal. Position 0 is the length byte.
func charAt(s string, i int) int {
	b := charset.Encode(s)
	if i == 0 {
		return len(b)
	}
	if i > len(b) {
		return 0
	}
	return int(b[i-1])
}

// setChar implements s[i] := c for 1 <= i <= Length(s).
func setChar(s string, i, c int) string {
	b := charset.Encode(s)
	if i >= 1 && i <= len(b) {
		b[i-1] = byte(c)
	}
	return charset.Decode(b)
}

func upCase(c int) int {
	if c >= 'a' && c <= 'z' {
		return c - 32
	}
	return c
}
