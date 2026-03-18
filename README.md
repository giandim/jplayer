# jplayer

Play your music from the terminal.

A TUI music player written in Go. Browse your local music library, manage a queue, and control playback — all from the terminal.

## Dependencies

jplayer requires `mpv` to be installed on your system.

**macOS**
```bash
brew install mpv
```

**Ubuntu/Debian**
```bash
sudo apt install mpv
```

**Arch**
```bash
sudo pacman -S mpv
```

## Installation

```bash
TBD
```

## Usage

By default jplayer opens your home directory. Navigate to your music folder and start playing.

## Keybindings

### Navigation
| Key | Action |
|-----|--------|
| `j` / `↓` | Move cursor down |
| `k` / `↑` | Move cursor up |
| `Enter` | Open folder / play track |
| `Backspace` | Go up a directory |


### Playback (TBD)
| Key | Action |
|-----|--------|

### General
| Key | Action |
|-----|--------|
| `q` | Quit |

## Planned features

- [ ] Navidrome / Subsonic streaming
- [ ] Spotify integration 
- [ ] Audio metadata parsing (ID3, Vorbis)
- [ ] Queue management
- [ ] Responsive panel layout

## Built with

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) — TUI framework
- [Lipgloss](https://github.com/charmbracelet/lipgloss) — terminal styling
- [mpv](https://mpv.io) — media playback
