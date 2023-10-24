package trayMenu

import (
	"github.com/getlantern/systray"
	"github.com/hra42/Go-TextType/internal/errorHandling"
	"github.com/hra42/Go-TextType/internal/hotKeyConfig"
	"github.com/hra42/Go-TextType/internal/logging"
	"os"
)

func OnExit() {
	logging.Logger.Println("Exit the program...")
	err := logging.LogFile.Close()
	errorHandling.CheckError(err)
	// when the program exits, the hotkey will unregister
	hotKeyConfig.UnregisterHotkey(hotKeyConfig.HK)
	systray.Quit()
	os.Exit(0)
}
