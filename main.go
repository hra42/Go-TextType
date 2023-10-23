package main

import (
	"embed"
	"encoding/gob"
	"github.com/gen2brain/beeep"
	"github.com/getlantern/systray"
	"github.com/go-vgo/robotgo"
	"github.com/pkg/browser"
	"github.com/tcnksm/go-latest"
	"golang.design/x/clipboard"
	"golang.design/x/hotkey"
	"log"
	"os"
	"strings"
	"time"
)

var (
	// AppVersion BuildID Information
	AppVersion string
	BuildID    string
)

var HotKeyConfig struct {
	HotkeyNumber int
	EnterKey     bool
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
	checkUpdate()
	loadHK()
	setup()
	systray.Run(onReady, onExit)
}

func onReady() {
	Logger.Println("App is running...")
	systray.SetIcon(readIcon())
	systray.SetTitle("Text Type")
	systray.SetTooltip("Control Text Type")

	mHK := systray.AddMenuItem("Modify Hotkey", "Modify the hotkey")
	mHK1 := mHK.AddSubMenuItemCheckbox("Ctrl + Shift + V", "Ctrl + Shift + V", true)
	mHK2 := mHK.AddSubMenuItemCheckbox("Ctrl + Alt + S", "Ctrl + Alt + S", false)
	mHK3 := mHK.AddSubMenuItemCheckbox("Ctrl + Q", "Ctrl + Q", false)
	systray.AddSeparator()
	selectEnterPressAfterPaste := systray.AddMenuItemCheckbox(
		"Should the enter key be pressed after the Text is pasted?",
		"Select Enter Press",
		false,
	)
	systray.AddSeparator()
	mStop := systray.AddMenuItem("Stop Text Type", "Stop the program")

	if HotKeyConfig.HotkeyNumber == 1 {
		mHK1.Check()
		mHK2.Uncheck()
		mHK3.Uncheck()
	} else if HotKeyConfig.HotkeyNumber == 2 {
		mHK1.Uncheck()
		mHK2.Check()
		mHK3.Uncheck()
	} else if HotKeyConfig.HotkeyNumber == 3 {
		mHK1.Uncheck()
		mHK2.Uncheck()
		mHK3.Check()
	} else {
		mHK1.Uncheck()
		mHK2.Uncheck()
		mHK3.Uncheck()
	}

	if HotKeyConfig.EnterKey {
		selectEnterPressAfterPaste.Check()
	} else {
		selectEnterPressAfterPaste.Uncheck()
	}

	go func() {
		for {
			select {
			case <-mHK1.ClickedCh:
				if HotKeyConfig.HotkeyNumber == 1 {
					Logger.Println("Hotkey", HotKeyConfig.HotkeyNumber, "is already selected")
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
					mHK1.Check()
					mHK2.Uncheck()
					mHK3.Uncheck()
				}
			case <-mHK2.ClickedCh:
				if HotKeyConfig.HotkeyNumber == 2 {
					Logger.Println("Hotkey", HotKeyConfig.HotkeyNumber, "is already used")
				} else {
					// delete current hotkey
					unregisterHotkey(HK)
					// set the selected hotkey
					HK = hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModAlt}, hotkey.KeyS)
					registerHotKey(HK)
					// save the selected hotkey to disk
					err := saveLastUsedHK(2)
					checkError(Logger, err)
					// set the menu
					mHK1.Uncheck()
					mHK2.Check()
					mHK3.Uncheck()
				}
			case <-mHK3.ClickedCh:
				if HotKeyConfig.HotkeyNumber == 3 {
					Logger.Println("Hotkey", HotKeyConfig.HotkeyNumber, "is already used")
				} else {
					// delete current hotkey
					unregisterHotkey(HK)
					// set the selected hotkey
					HK = hotkey.New([]hotkey.Modifier{hotkey.ModCtrl}, hotkey.KeyQ)
					registerHotKey(HK)
					// save the selected hotkey to disk
					err := saveLastUsedHK(3)
					checkError(Logger, err)
					// set the menu
					mHK1.Uncheck()
					mHK2.Uncheck()
					mHK3.Check()
				}
			case <-mStop.ClickedCh:
				onExit()
			case <-HK.Keydown():
				textType()
			case <-selectEnterPressAfterPaste.ClickedCh:
				if HotKeyConfig.EnterKey == true {
					HotKeyConfig.EnterKey = false
					selectEnterPressAfterPaste.Uncheck()
					err := saveLastUsedHK(HotKeyConfig.HotkeyNumber)
					if err != nil {
						Logger.Println(err)
					}
				} else {
					HotKeyConfig.EnterKey = true
					err := saveLastUsedHK(HotKeyConfig.HotkeyNumber)
					if err != nil {
						Logger.Println(err)
					}
					selectEnterPressAfterPaste.Check()
				}
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
	err := hk.Register()
	Logger.Println("hotkey", HotKeyConfig.HotkeyNumber, "registered")
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

	// trim whitespace from the clipboard text
	clipBoardText = strings.TrimSpace(clipBoardText)
	// use robotgo to type the clipboard text.
	robotgo.TypeStr(clipBoardText)
	// press enter key
	if HotKeyConfig.EnterKey == true {
		robotgo.KeyTap("enter")
	}
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
	Logger.Println("hotkey", HotKeyConfig.HotkeyNumber, "unregistered")
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

func deleteFile(file *os.File) {
	err := os.Remove(file.Name())
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
			return
		}
		Logger.Println("Error: ", err)
		checkError(Logger, err)
	}
	defer closeFile(file)

	dec := gob.NewDecoder(file)
	err = dec.Decode(&HotKeyConfig)
	checkError(Logger, err)
	Logger.Println("hotkey.gob loaded and decoded")
	switch HotKeyConfig.HotkeyNumber {
	case 1:
		Logger.Println("Hotkey 1 is in use")
		HK = hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModShift}, hotkey.KeyV)
		registerHotKey(HK)
	case 2:
		Logger.Println("Hotkey 2 is in use")
		HK = hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModAlt}, hotkey.KeyS)
		registerHotKey(HK)
	case 3:
		Logger.Println("Hotkey 3 is in use")
		HK = hotkey.New([]hotkey.Modifier{hotkey.ModCtrl}, hotkey.KeyQ)
		registerHotKey(HK)
	}
}

func checkUpdate() {
	dumpfile, err := os.CreateTemp("", "icon.*.png")
	checkError(Logger, err)
	defer deleteFile(dumpfile)
	if _, err = dumpfile.Write(readIcon()); err != nil {
		checkError(Logger, err)
	}
	if err = dumpfile.Close(); err != nil {
		checkError(Logger, err)
	}
	githubTag := &latest.GithubTag{
		Owner:      "HRA42",
		Repository: "Go-TextType",
	}
	res, err := latest.Check(githubTag, AppVersion)
	checkError(Logger, err)
	if res.Outdated {
		err = beeep.Alert(
			"Update Available!",
			"A new version of Go-TextType is available!",
			dumpfile.Name(),
		)
		checkError(Logger, err)
		Logger.Println("Update available")
		err := browser.OpenURL("https://github.com/HRA42/Go-TextType/releases/latest")
		checkError(Logger, err)
	}
}
