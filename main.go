package main

import (
	"embed"
	"encoding/gob"
	"github.com/getlantern/systray"
	"github.com/go-vgo/robotgo"
	"golang.design/x/clipboard"
	"golang.design/x/hotkey"
	"log"
	"os"
	"time"
)

var (
	// AppVersion BuildID Information
	AppVersion string
	BuildID    string
)

var HotKeyConfig struct {
	HotkeyNumber int
}

//go:embed icon.ico
var Icon embed.FS

// Logger setup
var Logger = setupLogger()

// LogFile setup
var LogFile *os.File

// HK Hotkey setup
var HK *hotkey.Hotkey

func main() {
	Logger.Println("The App Version is: ", AppVersion)
	Logger.Println("The Build ID is: ", BuildID)
	loadHK()
	setup()
	systray.Run(onReady, onExit)
}

func onReady() {
	Logger.Println("App is running...")
	systray.SetIcon(readIcon())
	systray.SetTitle("Text Type")
	systray.SetTooltip("Control Text Type")

	mHk := systray.AddMenuItem("Modify Hotkey", "Modify the hotkey")
	mHK1 := mHk.AddSubMenuItemCheckbox("Ctrl + Shift + V", "Ctrl + Shift + V", true)
	mHK2 := mHk.AddSubMenuItemCheckbox("Ctrl + Shift + S", "Ctrl + Shift + S", false)
	systray.AddSeparator()
	mStop := systray.AddMenuItem("Stop Text Type", "Stop the program")

	if HotKeyConfig.HotkeyNumber == 1 {
		mHK1.Check()
		mHK2.Uncheck()
	} else if HotKeyConfig.HotkeyNumber == 2 {
		mHK1.Uncheck()
		mHK2.Check()
	} else {
		mHK1.Uncheck()
		mHK2.Uncheck()
	}

	go func() {
		for {
			select {
			case <-mHK1.ClickedCh:
				if HotKeyConfig.HotkeyNumber == 1 {
					continue
				} else {
					// delete current hotkey
					unregisterHotkey(HK)
					// set the hotkey
					HK = hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModShift}, hotkey.KeyV)
					registerHotKey(HK)
					// save last used hotkey to disk
					err := saveLastUsedHK(1)
					checkError(Logger, err)
					// Update Menu
					mHK2.Uncheck()
					mHK1.Check()
				}
			case <-mHK2.ClickedCh:
				if HotKeyConfig.HotkeyNumber == 2 {
					continue
				} else {
					// delete current hotkey
					unregisterHotkey(HK)
					// set the selected hotkey
					HK = hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModShift}, hotkey.KeyS)
					registerHotKey(HK)
					// save the selected hotkey to disk
					err := saveLastUsedHK(2)
					checkError(Logger, err)
					// set the menu
					mHK1.Uncheck()
					mHK2.Check()
				}
			case <-mStop.ClickedCh:
				onExit()
			case <-HK.Keydown():
				textType()
			}
		}
	}()
}

func onExit() {
	Logger.Println("Exit the program...")
	err := LogFile.Close()
	checkError(Logger, err)
	// when the program exits, the hotkey will unregister
	unregisterHotkey(HK)
	systray.Quit()
	os.Exit(0)
}

func setup() {
	// register clipboard
	err := clipboard.Init()
	checkError(Logger, err)
}

func registerHotKey(hk *hotkey.Hotkey) {
	// hk = hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModShift}, hotkey.KeyV)
	err := hk.Register()
	Logger.Println("hotkey registered")
	checkError(Logger, err)
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
	Logger.Println("hotkey unregistered")
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
	logger = log.New(LogFile, "TextType ", log.LstdFlags)
	return
}

func closeFile(file *os.File) {
	err := file.Close()
	checkError(Logger, err)
}

func saveLastUsedHK(hkNumber int) error {
	file, err := os.Create("hotkey.gob")
	if err != nil {
		return err
	}
	defer closeFile(file)

	HotKeyConfig.HotkeyNumber = hkNumber
	enc := gob.NewEncoder(file)
	err = enc.Encode(HotKeyConfig)
	if err != nil {
		return err
	}

	return nil
}

func loadHK() {
	file, err := os.Open("hotkey.gob")
	if err != nil {
		if os.IsNotExist(err) {
			Logger.Println("No hotkey.gob file found")
			HK = hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModShift}, hotkey.KeyV)
			registerHotKey(HK)
			err = saveLastUsedHK(1)
			checkError(Logger, err)
		}
		Logger.Println("Error: ", err)
		checkError(Logger, err)
	}
	defer closeFile(file)

	dec := gob.NewDecoder(file)
	err = dec.Decode(&HotKeyConfig)
	checkError(Logger, err)

	switch HotKeyConfig.HotkeyNumber {
	case 1:
		HK = hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModShift}, hotkey.KeyV)
		registerHotKey(HK)
	case 2:
		HK = hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModShift}, hotkey.KeyS)
		registerHotKey(HK)
	}
}
