# leetui

A terminal UI for browsing and solving LeetCode problems.

## Install

### Arch Linux (AUR)

```sh
# stable, built from source
yay -S leetui

# stable, prebuilt binary
yay -S leetui-bin

# bleeding edge, built from main
yay -S leetui-git
```

### `go install`

```sh
go install github.com/lbarto12/leetui@latest
```

### From source

```sh
git clone https://github.com/lbarto12/leetui.git
cd leetui
go build -o leetui .
```

## Debug logging

Set `LEETUI_DEBUG=1` to write logs to `$XDG_STATE_HOME/leetui/debug.log`
(falls back to `~/.local/state/leetui/debug.log`).

## License

MIT — see [LICENSE](LICENSE).
