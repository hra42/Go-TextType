package hotKeyConfig

import (
	"github.com/hra42/Go-TextType/internal/errorHandling"
	"github.com/hra42/Go-TextType/internal/logging"
	"golang.design/x/hotkey"
	"os"
)

var HotKeyConfiguration struct {
	HotkeyNumber int
	EnterKey     bool
}

// HK Hotkey setup
var HK *hotkey.Hotkey

func RegisterHotKey(hk *hotkey.Hotkey) {
	err := hk.Register()
	logging.Logger.Println("hotkey", HotKeyConfiguration.HotkeyNumber, "registered")
	errorHandling.CheckError(err)
}

func CloseFile(file *os.File) {
	err := file.Close()
	errorHandling.CheckError(err)
}

func UnregisterHotkey(hk *hotkey.Hotkey) {
	err := hk.Unregister()
	errorHandling.CheckError(err)
	logging.Logger.Println("hotkey", HotKeyConfiguration.HotkeyNumber, "unregistered")
}
