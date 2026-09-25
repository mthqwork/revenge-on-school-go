package game

import "tclinux-suli/internal/dos"

// Run is the main program body (1010:c3ed). It returns when the game halts.
func (g *Game) Run() (err error) {
	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(dos.HaltError); ok {
				return
			}
			panic(r)
		}
	}()
	// The original checks for the SVGA driver and all data files here
	// (CheckFiles); the data is embedded in the port.
	g.resetViewPort()
	g.FindFirst("data\\savegame.tsi")
	loadGame := false
	if g.DosError == 0 {
		for n := 1; n <= 3; n++ {
			if g.ParamStr(n) == "loadgame" {
				loadGame = true
			}
		}
	}
	if !loadGame {
		g.fadeOutLoadPal(0, "data\\rgb.col")
		g.showTitle()
		g.fadeIn(10)
		g.FlushKeys()
		g.WaitKey()
	}
	g.fadeOutLoadPal(10, "data\\rgb.col")
	g.drawFrame()
	g.gameInit()
	for {
		g.mainLoopStep()
	}
}
