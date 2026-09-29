# Revenge on School – natív port

A **Revenge on School** (Tomcat Software, 1994–95) DOS-os kalandjáték natív
újraírása Go-ban. Egyetlen futtatható fájl, minden adat bele van ágyazva.
Linuxon (PC-n és ARM-on, pl. Raspberry Pi-n), Windowson és macOS-en fut,
DOSBox és emulátor nélkül.

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
GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -trimpath -ldflags "-s -w -H windowsgui" -o suli.exe ./cmd/suli

# macOS, Apple Silicon (M1/M2/...)
GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build -trimpath -ldflags "-s -w" -o suli-mac-arm64 ./cmd/suli

# macOS, Intel
GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -trimpath -ldflags "-s -w" -o suli-mac-amd64 ./cmd/suli
```

### Linux ARM (pl. Raspberry Pi)

ARM-os Linuxra is keresztfordítható, cgo nélkül, így nem kell hozzá ARM-os
C-fordító vagy más eszközlánc, elég a Go:

```sh
# 64 bites ARM Linux (aarch64), pl. 64 bites Raspberry Pi OS
GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -trimpath -ldflags "-s -w" -o suli-linux-arm64 ./cmd/suli

# 32 bites ARM Linux, pl. 32 bites Raspberry Pi OS
# GOARM=6: minden Pi-n fut (a Zero-n és az 1-esen is); GOARM=7 csak Pi 2-től felfelé
GOOS=linux GOARCH=arm GOARM=6 CGO_ENABLED=0 go build -trimpath -ldflags "-s -w" -o suli-linux-armv6 ./cmd/suli
```

Hogy melyik kell: a célgépen a `uname -m` parancs `aarch64`-et ír 64 bites
rendszeren, `armv6l`/`armv7l`-t 32 bitesen. Magán a Pi-n is lefordítható, ha
van rajta Go: `CGO_ENABLED=0 go build -o suli ./cmd/suli`.

Futtatáshoz asztali környezet (X11 vagy Wayland) és működő OpenGL-driver kell;
a friss Raspberry Pi OS-en ez alapból megvan.

### Böngésző (WebAssembly)

A játék böngészőben is fut, telepítés nélkül. A `web` mappában van hozzá egy
oldal: a játékképernyő alatt érintős gombok (nyilak, Enter, Esc, Szóköz,
PgUp/PgDn, Del, Alt-S), így telefonon is játszható.

```sh
GOOS=js GOARCH=wasm go build -trimpath -ldflags "-s -w" -o web/suli.wasm ./cmd/suli
cp "$(go env GOROOT)/lib/wasm/wasm_exec.js" web/
python3 -m http.server -d web 8000   # majd: http://localhost:8000
```

A `web` mappa tartalma bármilyen statikus tárhelyre feltölthető (pl. GitHub
Pages). `file://`-ból megnyitva nem működik, webszerver kell hozzá. A
`suli.wasm` kb. 15 MB, tömörítve kb. 3,5 MB; az oldal gzippel tömörített
`suli.wasm`-ot is elfogad.

A mentés (Alt-S) böngészőben a böngésző saját tárolójába (localStorage) kerül,
a „Mentés betöltése” gomb onnan tölti vissza. Az asztali változatok továbbra
is fájlba mentenek.

### Megjegyzések

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
| `web` | Böngészős változat: az oldal, ami a WebAssembly-buildet betölti, érintős gombokkal |

A játéklogika szándékosan követi az eredeti Turbo Pascal program szerkezetét,
a furcsaságaival együtt.
