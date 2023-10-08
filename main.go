package main

import (
	"embed"
	"fmt"
	"github.com/getlantern/systray"
	"golang.design/x/clipboard"
	"golang.design/x/hotkey"
	"log"
	"os"
	"os/exec"
	"time"
)

//go:embed icon.ico
var Icon embed.FS

// Logger setup
var Logger = setupLogger()

// LogFile setup
var LogFile *os.File

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	Logger.Println("App is running...")
	systray.SetIcon(readIcon())
	systray.SetTitle("Text Type")
	systray.SetTooltip("Control Text Type")

	mStart := systray.AddMenuItem("Start Text Type", "Start the program")
	mStop := systray.AddMenuItem("Stop Text Type", "Stop the program")

	go func() {
		for {
			select {
			case <-mStart.ClickedCh:
				textType()
			case <-mStop.ClickedCh:
				log.Fatal("Stop the program")
			}
		}
	}()
}

func onExit() {
	Logger.Println("Exit the program...")
	LogFile.Close()
	systray.Quit()
	os.Exit(0)
}

func textType() {
	// register clipboard
	err := clipboard.Init()
	checkError(Logger, err)

	// register hotkey
	hk := hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModShift}, hotkey.KeyV)
	err = hk.Register()
	// when the program exits, the hotkey will unregister
	defer unregisterHotkey(hk)
	Logger.Println("hotkey registered")
	checkError(Logger, err)

	for {
		// listen for hotkey press
		<-hk.Keydown()
		Logger.Println("hotkey pressed")

		// get clipboard text
		clipBoardText := string(clipboard.Read(clipboard.FmtText))
		if clipBoardText == "" {
			Logger.Println("clipboard is empty")
			break
		}
		Logger.Println(clipBoardText)

		// wait for 500ms before executing
		time.Sleep(time.Millisecond * 500)

		// build PowerShell command
		/*
			TODO: The needs to be replaced with a better solution
			When the program runs without CMD Prompt open windows
			looses the focus of the main application the text is pasted into.
		*/
		cmdTemplate := "[void] [System.Reflection.Assembly]::LoadWithPartialName(\"System.Windows.Forms\"); " +
			"$s = '%s'; " +
			"$s = $s -replace '%%', '{%%}'; " +
			"[System.Windows.Forms.SendKeys]::SendWait($s)"
		cmdText := fmt.Sprintf(cmdTemplate, clipBoardText)

		// execute PowerShell command
		cmd := exec.Command("powershell", "-command", cmdText)
		cmdOut, err := cmd.CombinedOutput()
		if err != nil {
			Logger.Println(err)
			Logger.Println("Command output:", string(cmdOut))
		}

		Logger.Println("clipboard entered")
	}
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
