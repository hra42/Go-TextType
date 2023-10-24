package trayMenu

import "embed"

//go:embed icon.ico
var Icon embed.FS

func ReadIcon() (data []byte) {
	data, _ = Icon.ReadFile("icon.ico")
	return
}
