package main

import (
	"github.com/getlantern/systray"
	"github.com/hra42/Go-TextType/internal/hotKeyConfig"
	"github.com/hra42/Go-TextType/internal/logging"
	"github.com/hra42/Go-TextType/internal/textType"
	"github.com/hra42/Go-TextType/internal/trayMenu"
	"github.com/hra42/Go-TextType/internal/update"
)

var (
	// AppVersion BuildID Information
	AppVersion string
	BuildID    string
)

func main() {
	logging.Logger = logging.SetupLogger()
	logging.Logger.Println("The App Version is", AppVersion)
	logging.Logger.Println("The Build is", BuildID)
	update.CheckUpdate(AppVersion)
	hotKeyConfig.LoadHK()
	textType.Setup()
	systray.Run(trayMenu.OnReady, trayMenu.OnExit)
}
