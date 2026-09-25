package main

import (
	"image"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"tclinux-suli/internal/bgi"
	"tclinux-suli/internal/charset"
	"tclinux-suli/internal/dos"
	"tclinux-suli/internal/game"
)

type window struct {
	m      *dos.Machine
	done   chan struct{}
	frame  *image.RGBA
	img    *ebiten.Image
	held   map[ebiten.Key]time.Time
	chars  []rune
	pushed map[ebiten.Key]time.Time
}

// special keys that produce ReadKey codes without a printable character.
var specialKeys = map[ebiten.Key][]byte{
	ebiten.KeyEnter: {13}, ebiten.KeyNumpadEnter: {13}, ebiten.KeyEscape: {27},
	ebiten.KeyBackspace: {8}, ebiten.KeyTab: {9},
	ebiten.KeyArrowUp: {0, scUp}, ebiten.KeyArrowDown: {0, scDown},
	ebiten.KeyArrowLeft: {0, scLeft}, ebiten.KeyArrowRight: {0, scRight},
	ebiten.KeyPageUp: {0, scPgUp}, ebiten.KeyPageDown: {0, scPgDn},
	ebiten.KeyHome: {0, scHome}, ebiten.KeyEnd: {0, scEnd},
	ebiten.KeyInsert: {0, scIns}, ebiten.KeyDelete: {0, scDel},
}

const (
	repeatDelay = 400 * time.Millisecond
	repeatRate  = 50 * time.Millisecond
)

func (w *window) Update() error {
	select {
	case <-w.done:
		return ebiten.Termination
	default:
	}
	now := time.Now()
	alt := ebiten.IsKeyPressed(ebiten.KeyAlt)
	// F11 or Alt+Enter toggles full screen (not passed to the game).
	if inpututil.IsKeyJustPressed(ebiten.KeyF11) || alt && inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
		w.pushed[ebiten.KeyEnter] = now.Add(time.Hour) // no auto repeat either
		return nil
	}
	// Special keys with auto repeat.
	for k, code := range specialKeys {
		if inpututil.IsKeyJustPressed(k) {
			w.m.PushKey(code...)
			w.pushed[k] = now.Add(repeatDelay)
		} else if ebiten.IsKeyPressed(k) {
			if t, ok := w.pushed[k]; ok && now.After(t) {
				w.m.PushKey(code...)
				w.pushed[k] = now.Add(repeatRate)
			}
		} else {
			delete(w.pushed, k)
		}
	}
	// Alt + letter.
	if alt {
		for r, sc := range altScan {
			k := ebiten.Key(int(ebiten.KeyA) + int(r-'a'))
			if inpututil.IsKeyJustPressed(k) {
				w.m.PushKey(0, sc)
			}
		}
		return nil
	}
	// Printable characters (keyboard layout aware).
	w.chars = ebiten.AppendInputChars(w.chars[:0])
	for _, r := range w.chars {
		if b := charset.Encode(string(r)); len(b) == 1 && b[0] != '?' || r == '?' {
			w.m.PushKey(b...)
		}
	}
	return nil
}

func (w *window) Draw(screen *ebiten.Image) {
	w.m.Scr.Snapshot(w.frame)
	w.img.WritePixels(w.frame.Pix)
	screen.DrawImage(w.img, nil)
}

func (w *window) Layout(int, int) (int, int) { return bgi.Width, bgi.Height }

func runWindow(m *dos.Machine, g *game.Game, scale int) error {
	w := &window{
		m:      m,
		done:   make(chan struct{}),
		frame:  image.NewRGBA(image.Rect(0, 0, bgi.Width, bgi.Height)),
		img:    ebiten.NewImage(bgi.Width, bgi.Height),
		pushed: map[ebiten.Key]time.Time{},
	}
	go func() {
		defer close(w.done)
		g.Run()
	}()
	ebiten.SetWindowTitle("Revenge on School")
	ebiten.SetWindowSize(bgi.Width*scale, bgi.Height*scale)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	err := ebiten.RunGame(w)
	m.Quit()
	<-w.done
	return err
}
