package textType

import (
	"github.com/go-vgo/robotgo"
	"github.com/hra42/Go-TextType/internal/errorHandling"
	"github.com/hra42/Go-TextType/internal/hotKeyConfig"
	"github.com/hra42/Go-TextType/internal/logging"
	"golang.design/x/clipboard"
	"strings"
	"time"
)

func TextType() {
	logging.Logger.Println("hotkey pressed")

	// get clipboard text
	clipBoardText := string(clipboard.Read(clipboard.FmtText))
	if clipBoardText == "" {
		logging.Logger.Println("clipboard is empty")
		return
	} else {
		logging.Logger.Println("clipboard has text")
	}

	// wait for 250ms before executing to make sure the clipboard has the text
	time.Sleep(time.Millisecond * 250)

	// trim whitespace from the clipboard text
	clipBoardText = strings.TrimSpace(clipBoardText)
	// use robotgo to type the clipboard text.
	robotgo.TypeStr(clipBoardText)
	// press enter key
	if hotKeyConfig.HotKeyConfiguration.EnterKey {
		robotgo.KeyTap("enter")
	}
	logging.Logger.Println("clipboard entered")
}

func Setup() {
	// register clipboard
	err := clipboard.Init()
	errorHandling.CheckError(err)
}
