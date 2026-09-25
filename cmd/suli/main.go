// Command suli is the native port of the DOS game "Revenge on School"
// (Tomcat Software, 1994-95).
//
//	suli [loadgame]            play in a window; "loadgame" loads the save
//	suli -headless -keys ...   run without a window (for tests), see keys.go
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"tclinux-suli/assets"
	"tclinux-suli/internal/bgi"
	"tclinux-suli/internal/dos"
	"tclinux-suli/internal/game"
)

func main() {
	headless := flag.Bool("headless", false, "run without a window, driven by -keys")
	keys := flag.String("keys", "", "scripted input for -headless (see keys.go)")
	shot := flag.String("shot", "", "headless: write a PNG of the final screen here")
	savePath := flag.String("save", defaultSavePath(), "save file location")
	scale := flag.Int("scale", 2, "window scale factor")
	seed := flag.Int64("seed", time.Now().UnixNano(), "random seed")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: %s [flags] [loadgame]\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	scr := bgi.New(assets.Font8x8)
	m := dos.New(scr)
	g := game.New(m, *seed)
	g.SavePath = *savePath
	g.Args = flag.Args()

	var err error
	if *headless {
		err = runHeadless(m, g, *keys, *shot)
	} else {
		err = runWindow(m, g, *scale)
	}
	if s := m.Console.String(); s != "" {
		fmt.Print(s)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func defaultSavePath() string {
	dir, err := os.UserConfigDir()
	if err != nil {
		dir = "."
	}
	return filepath.Join(dir, "tclinux-suli", "savegame.tsi")
}

func init() {
	// Accept the original style "LOADGAME" parameter in any case.
	for i, a := range os.Args {
		if i > 0 && strings.EqualFold(a, "loadgame") {
			os.Args[i] = "loadgame"
		}
	}
}
