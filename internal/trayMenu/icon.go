package trayMenu

import "embed"

//go:embed icon.ico
var Icon embed.FS

func readIcon() (data []byte, err error) {
	data, err = Icon.ReadFile("icon.ico")
	return
}
