package game

import (
	"tclinux-suli/assets"
)

// Hand-written ports of the picture loading and palette fading routines.

// loadPic loads a BGI image file from DATA. The files are GetImage dumps
// padded to 128 byte records.
func (g *Game) loadPic(name string) []byte {
	b, err := assets.ReadData(name)
	if err != nil {
		g.DiskError()
	}
	return b
}

// showRoomPic (1e9b:6792) shows a 160x100 room picture at 20,20.
func (g *Game) showRoomPic(name string) {
	g.scr.PutImage(20, 20, g.loadPic(name))
}

// showTiled draws a 640x400 picture stored as eight 160x200 tiles.
func (g *Game) showTiled(prefix string) {
	for t := 0; t < 8; t++ {
		name := prefix + string(rune('1'+t)) + ".pic"
		g.scr.PutImage(t%4*160, t/4*200, g.loadPic(name))
	}
}

// showTitle (1e9b:68f0).
func (g *Game) showTitle() { g.showTiled("title0") }

// endSequence (1e9b:1593): the Tomcat Software closing screen.
func (g *Game) endSequence() {
	g.ClearViewPort()
	g.fadeOutLoadPal(0, "endseq.col")
	g.showTiled("endseq0")
	g.SetColor(5)
	g.OutTextXY(230, 270, "Tomcat Software Incorporated")
	g.OutTextXY(238, 290, "A DamnedSoft Hungary tagja")
	g.fadeIn(10)
	g.FlushKeys()
	g.WaitKey()
	g.fadeOutLoadPal(10, "endseq.col")
}

// ---------------------------------------------------------------- palette

// targetPal is the palette FadeIn fades to ([2ef2] in the palette unit).
func (g *Game) readPal(name string) [256][3]byte {
	var p [256][3]byte
	b, err := assets.ReadData(name)
	if err != nil {
		g.DiskError()
	}
	for i := 0; i < 256 && i*3+2 < len(b); i++ {
		p[i] = [3]byte{b[i*3], b[i*3+1], b[i*3+2]}
	}
	return p
}

// fadeWait waits speed/20 (at least one) retraces; speed 0 means no wait.
func (g *Game) fadeWait(speed int) {
	if speed == 0 {
		return
	}
	n := speed / 20
	if n == 0 {
		n = 1
	}
	g.m.WaitRetrace(n)
}

// fadeOutLoadPal (26c2:0070): sets the palette from a .COL file, fades it to
// black and remembers the file palette as the target of the next fadeIn.
// The original takes a set of colours to exclude; the game always passes
// the empty set.
func (g *Game) fadeOutLoadPal(speed int, file string) {
	cur := g.readPal(file)
	g.targetPal = cur
	for {
		done := true
		for c := range cur {
			for k := 0; k < 3; k++ {
				if cur[c][k] > 0 {
					cur[c][k]--
				}
				if cur[c][k] != 0 {
					done = false
				}
			}
		}
		g.fadeWait(speed)
		g.scr.SetPalette(&cur)
		if done {
			return
		}
	}
}

// fadeIn (26c2:0262) fades from black to the palette loaded last.
func (g *Game) fadeIn(speed int) {
	var cur [256][3]byte
	for {
		done := true
		for c := range cur {
			for k := 0; k < 3; k++ {
				if cur[c][k] != g.targetPal[c][k] {
					cur[c][k]++
				}
				if cur[c][k] != g.targetPal[c][k] {
					done = false
				}
			}
		}
		g.fadeWait(speed)
		g.scr.SetPalette(&cur)
		if done {
			return
		}
	}
}
