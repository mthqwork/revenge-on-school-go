// Package bgi is a small software re-implementation of the parts of the
// Borland Graphics Interface that the game uses, on a 640x400, 256 colour
// palettised framebuffer (the original ran on the SVGA256 BGI driver).
//
// All drawing coordinates are relative to the current viewport, like in BGI.
// The screen is safe for one drawing goroutine plus concurrent readers that
// use Snapshot.
package bgi

import (
	"image"
	"image/color"
	"sync"

	"tclinux-suli/internal/charset"
)

const (
	Width  = 640
	Height = 400
)

// Screen is the emulated graphics device.
type Screen struct {
	mu   sync.Mutex
	pix  [Width * Height]byte
	pal  [256][3]byte // 6-bit VGA DAC values
	font []byte       // 256 glyphs, 8 bytes each

	color     byte
	fillColor byte
	vx1, vy1  int
	vx2, vy2  int
	clip      bool

	ops uint64 // number of drawing operations, see Ops
}

// Ops returns a counter that changes whenever something is drawn.
func (s *Screen) Ops() uint64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.ops
}

// New creates a screen that renders text with the given 8x8 font (2048 bytes).
func New(font []byte) *Screen {
	s := &Screen{font: font, color: 15, fillColor: 15}
	s.vx2, s.vy2, s.clip = Width-1, Height-1, true
	return s
}

func (s *Screen) GetMaxX() int { return Width - 1 }
func (s *Screen) GetMaxY() int { return Height - 1 }

func (s *Screen) SetColor(c int) {
	s.mu.Lock()
	s.color = byte(c)
	s.mu.Unlock()
}

// SetFillStyle only supports the solid pattern, which is all the game uses.
func (s *Screen) SetFillStyle(pattern, c int) {
	s.mu.Lock()
	s.fillColor = byte(c)
	s.mu.Unlock()
}

func (s *Screen) SetViewPort(x1, y1, x2, y2 int, clip bool) {
	s.mu.Lock()
	s.vx1, s.vy1, s.vx2, s.vy2, s.clip = x1, y1, x2, y2, clip
	s.mu.Unlock()
}

// ClearViewPort fills the current viewport with the background colour (0).
func (s *Screen) ClearViewPort() {
	s.mu.Lock()
	s.ops++
	defer s.mu.Unlock()
	s.fillAbs(s.vx1, s.vy1, s.vx2, s.vy2, 0)
}

// ClearDevice clears the whole screen and resets the viewport.
func (s *Screen) ClearDevice() {
	s.mu.Lock()
	s.ops++
	defer s.mu.Unlock()
	for i := range s.pix {
		s.pix[i] = 0
	}
	s.vx1, s.vy1, s.vx2, s.vy2, s.clip = 0, 0, Width-1, Height-1, true
}

// inside reports whether an absolute pixel is visible under the clip rules.
func (s *Screen) inside(x, y int) bool {
	if x < 0 || y < 0 || x >= Width || y >= Height {
		return false
	}
	if s.clip && (x < s.vx1 || x > s.vx2 || y < s.vy1 || y > s.vy2) {
		return false
	}
	return true
}

func (s *Screen) plot(x, y int, c byte) {
	if s.inside(x, y) {
		s.pix[y*Width+x] = c
	}
}

func (s *Screen) fillAbs(x1, y1, x2, y2 int, c byte) {
	if x1 > x2 {
		x1, x2 = x2, x1
	}
	if y1 > y2 {
		y1, y2 = y2, y1
	}
	for y := y1; y <= y2; y++ {
		for x := x1; x <= x2; x++ {
			s.plot(x, y, c)
		}
	}
}

// Bar fills a rectangle (both corners inclusive) with the fill colour.
func (s *Screen) Bar(x1, y1, x2, y2 int) {
	s.mu.Lock()
	s.ops++
	defer s.mu.Unlock()
	s.fillAbs(x1+s.vx1, y1+s.vy1, x2+s.vx1, y2+s.vy1, s.fillColor)
}

func (s *Screen) PutPixel(x, y, c int) {
	s.mu.Lock()
	s.ops++
	defer s.mu.Unlock()
	s.plot(x+s.vx1, y+s.vy1, byte(c))
}

// FloodFill fills the 4-connected area around x,y that is bounded by the
// border colour with the fill colour.
func (s *Screen) FloodFill(x, y, border int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.ops++
	b := byte(border)
	x, y = x+s.vx1, y+s.vy1
	if !s.inside(x, y) || s.pix[y*Width+x] == b {
		return
	}
	seen := make([]bool, Width*Height)
	stack := [][2]int{{x, y}}
	for len(stack) > 0 {
		p := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		px, py := p[0], p[1]
		if !s.inside(px, py) {
			continue
		}
		i := py*Width + px
		if seen[i] || s.pix[i] == b {
			continue
		}
		seen[i] = true
		s.pix[i] = s.fillColor
		stack = append(stack, [2]int{px + 1, py}, [2]int{px - 1, py}, [2]int{px, py + 1}, [2]int{px, py - 1})
	}
}

// Line draws a one pixel wide solid line including both end points.
func (s *Screen) Line(x1, y1, x2, y2 int) {
	s.mu.Lock()
	s.ops++
	defer s.mu.Unlock()
	x1, y1, x2, y2 = x1+s.vx1, y1+s.vy1, x2+s.vx1, y2+s.vy1
	dx, dy := abs(x2-x1), -abs(y2-y1)
	sx, sy := sign(x2-x1), sign(y2-y1)
	e := dx + dy
	for {
		s.plot(x1, y1, s.color)
		if x1 == x2 && y1 == y2 {
			return
		}
		e2 := 2 * e
		if e2 >= dy {
			e += dy
			x1 += sx
		}
		if e2 <= dx {
			e += dx
			y1 += sy
		}
	}
}

// OutTextXY draws s with the 8x8 default font, top-left aligned. Characters
// that do not fit completely in the viewport are skipped, as the bit-mapped
// BGI font does.
func (s *Screen) OutTextXY(x, y int, str string) {
	s.mu.Lock()
	s.ops++
	defer s.mu.Unlock()
	x, y = x+s.vx1, y+s.vy1
	for _, ch := range charset.Encode(str) {
		if s.inside(x, y) && s.inside(x+7, y+7) {
			g := s.font[int(ch)*8 : int(ch)*8+8]
			for r := 0; r < 8; r++ {
				bits := g[r]
				for c := 0; c < 8; c++ {
					if bits&(0x80>>c) != 0 {
						s.pix[(y+r)*Width+x+c] = s.color
					}
				}
			}
		}
		x += 8
	}
}

// TextWidth returns the pixel width of s in the default font.
func (s *Screen) TextWidth(str string) int { return 8 * charset.Len(str) }

// PutImage draws a BGI GetImage-format bitmap (uint16 width-1, uint16
// height-1, then one byte per pixel) at x,y with NormalPut.
func (s *Screen) PutImage(x, y int, img []byte) {
	s.mu.Lock()
	s.ops++
	defer s.mu.Unlock()
	if len(img) < 4 {
		return
	}
	w := int(img[0]) | int(img[1])<<8 + 1
	h := int(img[2]) | int(img[3])<<8 + 1
	data := img[4:]
	x, y = x+s.vx1, y+s.vy1
	for r := 0; r < h; r++ {
		for c := 0; c < w; c++ {
			i := r*w + c
			if i >= len(data) {
				return
			}
			if px, py := x+c, y+r; px >= 0 && py >= 0 && px < Width && py < Height {
				s.pix[py*Width+px] = data[i]
			}
		}
	}
}

// SetPalette sets all 256 DAC entries (6-bit components).
func (s *Screen) SetPalette(p *[256][3]byte) {
	s.mu.Lock()
	s.ops++
	s.pal = *p
	s.mu.Unlock()
}

// Palette returns a copy of the current DAC.
func (s *Screen) Palette() [256][3]byte {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.pal
}

// Snapshot renders the framebuffer through the palette into dst.
func (s *Screen) Snapshot(dst *image.RGBA) {
	s.mu.Lock()
	defer s.mu.Unlock()
	var lut [256]color.RGBA
	for i, c := range s.pal {
		lut[i] = color.RGBA{dac(c[0]), dac(c[1]), dac(c[2]), 255}
	}
	for i, p := range s.pix {
		c := lut[p]
		o := i * 4
		dst.Pix[o], dst.Pix[o+1], dst.Pix[o+2], dst.Pix[o+3] = c.R, c.G, c.B, 255
	}
}

// Image returns a freshly rendered RGBA copy of the screen.
func (s *Screen) Image() *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, Width, Height))
	s.Snapshot(img)
	return img
}

// dac expands a 6-bit VGA DAC value to 8 bits the way DOSBox does.
func dac(v byte) byte {
	v &= 63
	return v<<2 | v>>4
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func sign(v int) int {
	switch {
	case v < 0:
		return -1
	case v > 0:
		return 1
	}
	return 0
}
