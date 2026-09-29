//go:build js

package game

import (
	"encoding/base64"
	"errors"
	"syscall/js"
)

// saveKey is the localStorage key that holds the save game in the browser.
const saveKey = "revenge-on-school/savegame.tsi"

// loadSave reads the save game from the browser's localStorage.
func loadSave(string) (data []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			data, err = nil, errors.New("localStorage unavailable")
		}
	}()
	v := js.Global().Get("localStorage").Call("getItem", saveKey)
	if v.IsNull() || v.IsUndefined() {
		return nil, errors.New("no save game")
	}
	return base64.StdEncoding.DecodeString(v.String())
}

// storeSave writes the save game to the browser's localStorage.
func storeSave(_ string, data []byte) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("localStorage unavailable")
		}
	}()
	js.Global().Get("localStorage").Call("setItem", saveKey, base64.StdEncoding.EncodeToString(data))
	if cb := js.Global().Get("onSuliSaved"); cb.Type() == js.TypeFunction {
		cb.Invoke()
	}
	return nil
}
