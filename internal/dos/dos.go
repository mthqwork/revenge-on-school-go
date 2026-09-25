// Package dos provides the small slice of the DOS / Turbo Pascal runtime
// that the game logic relies on: a BIOS-like keyboard buffer (Crt.KeyPressed
// and Crt.ReadKey semantics), Delay, vertical retrace waits, text files and
// Halt.
//
// The game logic runs sequentially in its own goroutine exactly like the
// original program did; the front end (window or headless runner) feeds keys
// and displays the screen.
package dos

import (
	"bufio"
	"bytes"
	"strconv"
	"strings"
	"sync"
	"time"

	"tclinux-suli/internal/bgi"
	"tclinux-suli/internal/charset"
)

// HaltError is raised (via panic) by Halt and when the front end shuts down.
type HaltError struct{ Code int }

// Machine is the runtime environment shared by the game and the front end.
type Machine struct {
	Scr *bgi.Screen

	mu   sync.Mutex
	cond *sync.Cond
	keys []byte
	quit bool
	// Idle, if set, is called (without the lock held) whenever the game polls
	// or waits for a key and the buffer is empty. waiting is true when the
	// game is really waiting for input (ReadKey, or KeyPressed polled
	// repeatedly) rather than draining the buffer. The headless runner uses
	// it to feed scripted input. It returns true if it did something (added
	// a key or called Quit), false to let the machine sleep and poll again.
	Idle       func(waiting bool) bool
	emptyPolls int
	lastOps    uint64
	// Fast disables all real-time waiting (Delay, retrace) for tests.
	Fast bool

	// Console collects what the program writes to the DOS console after
	// leaving graphics mode.
	Console strings.Builder
}

func New(scr *bgi.Screen) *Machine {
	m := &Machine{Scr: scr}
	m.cond = sync.NewCond(&m.mu)
	return m
}

// PushKey appends a raw ReadKey byte to the keyboard buffer. Extended keys
// are pushed as 0 followed by the scan code.
func (m *Machine) PushKey(b ...byte) {
	m.mu.Lock()
	m.keys = append(m.keys, b...)
	m.mu.Unlock()
	m.cond.Broadcast()
}

// Quit makes the next keyboard or timing call in the game goroutine halt.
func (m *Machine) Quit() {
	m.mu.Lock()
	m.quit = true
	m.mu.Unlock()
	m.cond.Broadcast()
}

func (m *Machine) checkQuit() {
	if m.quit {
		m.mu.Unlock()
		panic(HaltError{})
	}
}

func (m *Machine) idle(waiting bool) {
	if m.Idle != nil && m.Idle(waiting) {
		return
	}
	time.Sleep(5 * time.Millisecond)
}

// KeyPressed reports whether a key is waiting (Crt.KeyPressed).
func (m *Machine) KeyPressed() bool {
	m.mu.Lock()
	m.checkQuit()
	n := len(m.keys)
	m.mu.Unlock()
	if n > 0 {
		return true
	}
	if ops := m.Scr.Ops(); ops != m.lastOps {
		m.lastOps = ops
		m.emptyPolls = 0
	}
	m.emptyPolls++
	if m.emptyPolls == 1 {
		// First empty poll: typically the end of a buffer drain loop.
		// Report without waiting so the drain finishes immediately.
		return false
	}
	m.idle(true)
	m.mu.Lock()
	m.checkQuit()
	n = len(m.keys)
	m.mu.Unlock()
	return n > 0
}

// ReadKey waits for and returns the next key byte (Crt.ReadKey).
func (m *Machine) ReadKey() byte {
	for {
		m.mu.Lock()
		m.checkQuit()
		if len(m.keys) > 0 {
			k := m.keys[0]
			m.keys = m.keys[1:]
			m.emptyPolls = 0
			m.mu.Unlock()
			return k
		}
		m.mu.Unlock()
		m.idle(true)
	}
}

// Delay sleeps ms milliseconds (Crt.Delay).
func (m *Machine) Delay(ms int) {
	m.sleep(time.Duration(ms) * time.Millisecond)
}

// WaitRetrace waits for n vertical retraces of the 70 Hz VGA mode.
func (m *Machine) WaitRetrace(n int) {
	m.sleep(time.Duration(n) * time.Second / 70)
}

func (m *Machine) sleep(d time.Duration) {
	m.mu.Lock()
	m.checkQuit()
	m.mu.Unlock()
	if !m.Fast {
		time.Sleep(d)
	}
}

// Halt terminates the game program.
func (m *Machine) Halt() { panic(HaltError{}) }

// WriteLn writes a line to the DOS console (used after CloseGraph).
func (m *Machine) WriteLn(s string) {
	m.Console.WriteString(s)
	m.Console.WriteByte('\n')
}

// TextFile is a Turbo Pascal text file variable.
type TextFile struct {
	Name  string        // assigned file name
	W     *bytes.Buffer // non-nil while open for writing
	lines []string
	pos   int
}

// SetContent replaces the readable contents and rewinds.
func (f *TextFile) SetContent(raw []byte) {
	*f = TextFile{Name: f.Name, lines: NewTextFile(raw).lines}
}

// SkipLine is ReadLn(f) without variables.
func (f *TextFile) SkipLine() {
	if f.pos < len(f.lines) {
		f.pos++
	}
}

// NewTextFile splits raw DOS text into lines, decoding the game charset.
func NewTextFile(raw []byte) *TextFile {
	f := &TextFile{}
	sc := bufio.NewScanner(bytes.NewReader(raw))
	sc.Buffer(make([]byte, 64*1024), 1<<20)
	for sc.Scan() {
		line := strings.TrimRight(sc.Text(), "\r")
		if i := strings.IndexByte(line, 0x1a); i >= 0 { // ^Z end of file
			line = line[:i]
		}
		f.lines = append(f.lines, charset.Decode([]byte(line)))
	}
	return f
}

// Reset rewinds the file.
func (f *TextFile) Reset() { f.pos = 0 }

// Eof reports whether all lines were read.
func (f *TextFile) Eof() bool { return f.pos >= len(f.lines) }

// ReadLn reads a line, truncated to max characters like a string[max]
// variable. At end of file it returns "".
func (f *TextFile) ReadLn(max int) string {
	if f.pos >= len(f.lines) {
		return ""
	}
	s := f.lines[f.pos]
	f.pos++
	return Trunc(s, max)
}

// ReadInt reads an integer from the start of the next line and skips the
// rest of the line (Read(f, int) + ReadLn(f)).
func (f *TextFile) ReadInt() int {
	s := strings.TrimSpace(f.ReadLn(255))
	if i := strings.IndexAny(s, " \t"); i >= 0 {
		s = s[:i]
	}
	v, _ := strconv.Atoi(s)
	return v
}

// Trunc limits s to max characters (Pascal string[max] assignment).
func Trunc(s string, max int) string {
	r := []rune(s)
	if len(r) > max {
		return string(r[:max])
	}
	return s
}

// Val converts a string to an integer like Pascal Val: on error the result
// is 0 and code is the 1-based position of the offending character.
func Val(s string) (v int, code int) {
	s = strings.TrimLeft(s, " ")
	n, err := strconv.Atoi(s)
	if err != nil {
		for i, c := range s {
			if (c < '0' || c > '9') && !(i == 0 && (c == '-' || c == '+')) {
				return 0, i + 1
			}
		}
		return 0, len(s) + 1
	}
	return n, 0
}
