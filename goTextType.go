package main

import (
	"fmt"
	"golang.design/x/clipboard"
	"golang.design/x/hotkey"
	"log"
	"os/exec"
	"time"
)

func main() {
	log.Println("App is running...")
	// register clipboard
	err := clipboard.Init()
	checkError(err)

	// register hotkey
	hk := hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModShift}, hotkey.KeyV)
	err = hk.Register()
	// when the program exits, the hotkey will unregister
	defer func(hk *hotkey.Hotkey) {
		err := hk.Unregister()
		checkError(err)
	}(hk)
	log.Println("hotkey registered")
	checkError(err)

	for {
		// listen for hotkey press
		<-hk.Keydown()
		log.Println("hotkey pressed")

		// get clipboard text
		clipBoardText := string(clipboard.Read(clipboard.FmtText))
		if clipBoardText == "" {
			log.Println("clipboard is empty")
			break
		}
		log.Println(clipBoardText)

		// wait for 500ms before executing
		time.Sleep(time.Millisecond * 500)

		// build PowerShell command
		cmdTemplate := "[void] [System.Reflection.Assembly]::LoadWithPartialName(\"System.Windows.Forms\"); " +
			"$s = '%s'; " +
			"$s = $s -replace '%%', '{%%}'; " +
			"[System.Windows.Forms.SendKeys]::SendWait($s)"
		cmdText := fmt.Sprintf(cmdTemplate, clipBoardText)

		// execute PowerShell command
		cmd := exec.Command("powershell", "-command", cmdText)
		cmdOut, err := cmd.CombinedOutput()
		if err != nil {
			log.Println(err)
			log.Println("Command output:", string(cmdOut))
		}

		log.Println("clipboard entered")
	}
}

func checkError(err error) {
	if err != nil {
		log.Println("Error: ", err)
	}
}
