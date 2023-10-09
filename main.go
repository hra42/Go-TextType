package main

import (
	"embed"
	"github.com/getlantern/systray"
	"github.com/go-vgo/robotgo"
	"golang.design/x/clipboard"
	"golang.design/x/hotkey"
	"log"
	"os"
	"time"
)

var (
	AppVersion string
	BuildID    string
)

//go:embed icon.ico
var Icon embed.FS

// Logger setup
var Logger = setupLogger()

// LogFile setup
var LogFile *os.File

// HK Hotkey setup
var HK = setup()

func main() {
	Logger.Println("The App Version is: ", AppVersion)
	Logger.Println("The Build ID is: ", BuildID)
	systray.Run(onReady, onExit)
}

func onReady() {
	Logger.Println("App is running...")
	systray.SetIcon(readIcon())
	systray.SetTitle("Text Type")
	systray.SetTooltip("Control Text Type")

	mStop := systray.AddMenuItem("Stop Text Type", "Stop the program")

	go func() {
		for {
			select {
			case <-mStop.ClickedCh:
				log.Fatal("Stop the program")
			case <-HK.Keydown():
				textType()
			}
		}
	}()
}

func onExit() {
	Logger.Println("Exit the program...")
	LogFile.Close()
	// when the program exits, the hotkey will unregister
	unregisterHotkey(HK)
	systray.Quit()
	os.Exit(0)
}

func setup() (hk *hotkey.Hotkey) {
	// register clipboard
	err := clipboard.Init()
	checkError(Logger, err)

	// register hotkey
	hk = hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModShift}, hotkey.KeyV)
	err = hk.Register()
	Logger.Println("hotkey registered")
	checkError(Logger, err)
	return
}

func textType() {
	Logger.Println("hotkey pressed")

	// get clipboard text
	clipBoardText := string(clipboard.Read(clipboard.FmtText))
	if clipBoardText == "" {
		Logger.Println("clipboard is empty")
		return
	} else {
		Logger.Println("clipboard has text")
	}

	// wait for 500ms before executing
	time.Sleep(time.Millisecond * 500)

	// use robotgo to type the clipboard text
	robotgo.TypeStr(clipBoardText)
	Logger.Println("clipboard entered")
}

func checkError(logger *log.Logger, err error) {
	if err != nil {
		logger.Println("Error: ", err)
	}
}

func unregisterHotkey(hk *hotkey.Hotkey) {
	err := hk.Unregister()
	checkError(Logger, err)
}

func readIcon() (data []byte) {
	data, _ = Icon.ReadFile("icon.ico")
	return
}

func setupLogger() (logger *log.Logger) {
	// Open the log file
	var err error
	LogFile, err = os.OpenFile("TextType.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// Create a new logger
	logger = log.New(LogFile, "app ", log.LstdFlags)
	return
}
