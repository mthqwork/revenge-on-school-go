package main

import (
	"fmt"
	"image/png"
	"os"

	"tclinux-suli/internal/dos"
	"tclinux-suli/internal/game"
)

// runHeadless runs the game without a window. Input comes from the script;
// the game runs as fast as possible. When the script is exhausted the game
// is stopped and the final screen is written to shot (if given).
func runHeadless(m *dos.Machine, g *game.Game, script, shot string) error {
	items, err := parseScript(script)
	if err != nil {
		return err
	}
	m.Fast = true
	var ioErr error
	m.Idle = func(waiting bool) bool {
		if !waiting {
			return false
		}
		for len(items) > 0 && items[0].shot != "" {
			if err := savePNG(m, items[0].shot); err != nil && ioErr == nil {
				ioErr = err
			}
			items = items[1:]
		}
		if len(items) == 0 {
			m.Quit()
			return true
		}
		m.PushKey(items[0].keys...)
		items = items[1:]
		return true
	}
	g.Run()
	if shot != "" {
		if err := savePNG(m, shot); err != nil {
			return err
		}
	}
	return ioErr
}

func savePNG(m *dos.Machine, name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()
	if err := png.Encode(f, m.Scr.Image()); err != nil {
		return fmt.Errorf("writing %s: %w", name, err)
	}
	return nil
}
