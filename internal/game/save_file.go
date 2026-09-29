//go:build !js

package game

import (
	"os"
	"path/filepath"
)

// loadSave reads the save game from the file system.
func loadSave(path string) ([]byte, error) {
	return os.ReadFile(path)
}

// storeSave writes the save game to the file system.
func storeSave(path string, data []byte) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o644)
}
