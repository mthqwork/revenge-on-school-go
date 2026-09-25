# Revenge on School – natív port

A **Revenge on School** (Tomcat Software, 1994–95) DOS-os kalandjáték natív
újraírása Go-ban. Egyetlen futtatható fájl, minden adat bele van ágyazva.
Linuxon, Windowson és macOS-en fut, DOSBox és emulátor nélkül.

A port a Tomcat Software engedélyével készült, az eredeti `SULI.EXE`
visszafejtésével, és ugyanazt a képet adja, mint az eredeti program.

A port AI segítségével (Claude) készült.

## Futtatás

```sh
go build -o suli ./cmd/suli
./suli              # új játék
./suli loadgame     # mentett állás betöltése
```

Kapcsolók: `-scale N` (ablaknagyítás, alapból 2), `-save FÁJL` (mentés helye,
alapból `~/.config/tclinux-suli/savegame.tsi`, Windowson `%AppData%`).

## Irányítás

Nyilak, Enter, Esc, PgUp/PgDn, Del, Szóköz – mint az eredetiben. **Alt-S**
mentés, **F11** vagy **Alt+Enter** teljes képernyő.

## Fordítás más rendszerekre

Linuxról keresztfordítva:

```sh
# Windows (64 bit)
GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -trimpath -ldflags "-s -w" -o suli.exe ./cmd/suli

# macOS, Apple Silicon (M1/M2/...)
GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build -trimpath -ldflags "-s -w" -o suli-mac-arm64 ./cmd/suli

# macOS, Intel
GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -trimpath -ldflags "-s -w" -o suli-mac-amd64 ./cmd/suli
```

A macOS-bináris nincs aláírva, ezért első indításkor a Gatekeeper blokkolhatja.
Ilyenkor a Macen: `xattr -d com.apple.quarantine suli-mac-arm64 && chmod +x suli-mac-arm64`.

Linuxon natívan az Ebitengine cgo-t használ, ehhez kellenek az X11/OpenGL
fejlécfájlok (Arch: `libx11 libxrandr libxcursor libxinerama libxi mesa`).

## Felépítés

| Hely | Tartalom |
|---|---|
| `cmd/suli` | Belépési pont: ablak (Ebitengine), billentyűzet, headless mód (`-headless -keys ... -shot kep.png`) |
| `internal/bgi` | A Borland BGI szükséges részei szoftveresen: 640×400×256 framebuffer, VGA 8×8 font, paletta |
| `internal/dos` | DOS/Turbo Pascal futtatókörnyezet: billentyűpuffer, Delay, szövegfájlok, Halt |
| `internal/charset` | Az eredeti kódlap (CP437 + CWI-2 ékezetek) ↔ UTF-8 |
| `internal/game` | A játéklogika. A `*_gen.go` fájlok a gépi kódból fordított Go kód, a többi kézzel írt |
| `assets` | Az eredeti adatfájlok és a font (beágyazva) |

A játéklogika szándékosan követi az eredeti Turbo Pascal program szerkezetét,
a furcsaságaival együtt.
