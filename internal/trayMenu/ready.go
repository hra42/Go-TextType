package trayMenu

import (
	"github.com/getlantern/systray"
	"github.com/hra42/Go-TextType/internal/errorHandling"
	"github.com/hra42/Go-TextType/internal/hotKeyConfig"
	"github.com/hra42/Go-TextType/internal/logging"
	"github.com/hra42/Go-TextType/internal/textType"
	"golang.design/x/hotkey"
)

func OnReady() {
	logging.Logger.Println("App is running...")
	systray.SetIcon(ReadIcon())
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

	if hotKeyConfig.HotKeyConfiguration.HotkeyNumber == 1 {
		mHK1.Check()
		mHK2.Uncheck()
		mHK3.Uncheck()
	} else if hotKeyConfig.HotKeyConfiguration.HotkeyNumber == 2 {
		mHK1.Uncheck()
		mHK2.Check()
		mHK3.Uncheck()
	} else if hotKeyConfig.HotKeyConfiguration.HotkeyNumber == 3 {
		mHK1.Uncheck()
		mHK2.Uncheck()
		mHK3.Check()
	} else {
		mHK1.Uncheck()
		mHK2.Uncheck()
		mHK3.Uncheck()
	}

	if hotKeyConfig.HotKeyConfiguration.EnterKey {
		selectEnterPressAfterPaste.Check()
	} else {
		selectEnterPressAfterPaste.Uncheck()
	}

	go func() {
		for {
			select {
			case <-mHK1.ClickedCh:
				if hotKeyConfig.HotKeyConfiguration.HotkeyNumber == 1 {
					logging.Logger.Println("Hotkey", hotKeyConfig.HotKeyConfiguration.HotkeyNumber, "is already selected")
				} else {
					// delete current hotkey
					hotKeyConfig.UnregisterHotkey(hotKeyConfig.HK)
					// set the hotkey
					HK := hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModShift}, hotkey.KeyV)
					hotKeyConfig.RegisterHotKey(HK)
					// save last used hotkey to disk
					err := hotKeyConfig.SaveLastUsedHK(1)
					errorHandling.CheckError(err)
					// Update Menu
					mHK1.Check()
					mHK2.Uncheck()
					mHK3.Uncheck()
				}
			case <-mHK2.ClickedCh:
				if hotKeyConfig.HotKeyConfiguration.HotkeyNumber == 2 {
					logging.Logger.Println("Hotkey", hotKeyConfig.HotKeyConfiguration.HotkeyNumber, "is already used")
				} else {
					// delete current hotkey
					hotKeyConfig.UnregisterHotkey(hotKeyConfig.HK)
					// set the selected hotkey
					hotKeyConfig.HK = hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModAlt}, hotkey.KeyS)
					hotKeyConfig.RegisterHotKey(hotKeyConfig.HK)
					// save the selected hotkey to disk
					err := hotKeyConfig.SaveLastUsedHK(2)
					errorHandling.CheckError(err)
					// set the menu
					mHK1.Uncheck()
					mHK2.Check()
					mHK3.Uncheck()
				}
			case <-mHK3.ClickedCh:
				if hotKeyConfig.HotKeyConfiguration.HotkeyNumber == 3 {
					logging.Logger.Println("Hotkey", hotKeyConfig.HotKeyConfiguration.HotkeyNumber, "is already used")
				} else {
					// delete current hotkey
					hotKeyConfig.UnregisterHotkey(hotKeyConfig.HK)
					// set the selected hotkey
					hotKeyConfig.HK = hotkey.New([]hotkey.Modifier{hotkey.ModCtrl}, hotkey.KeyQ)
					hotKeyConfig.RegisterHotKey(hotKeyConfig.HK)
					// save the selected hotkey to disk
					err := hotKeyConfig.SaveLastUsedHK(3)
					errorHandling.CheckError(err)
					// set the menu
					mHK1.Uncheck()
					mHK2.Uncheck()
					mHK3.Check()
				}
			case <-mStop.ClickedCh:
				OnExit()
			case <-hotKeyConfig.HK.Keydown():
				textType.TextType()
			case <-selectEnterPressAfterPaste.ClickedCh:
				if hotKeyConfig.HotKeyConfiguration.EnterKey == true {
					hotKeyConfig.HotKeyConfiguration.EnterKey = false
					selectEnterPressAfterPaste.Uncheck()
					err := hotKeyConfig.SaveLastUsedHK(hotKeyConfig.HotKeyConfiguration.HotkeyNumber)
					if err != nil {
						logging.Logger.Println(err)
					}
				} else {
					hotKeyConfig.HotKeyConfiguration.EnterKey = true
					err := hotKeyConfig.SaveLastUsedHK(hotKeyConfig.HotKeyConfiguration.HotkeyNumber)
					if err != nil {
						logging.Logger.Println(err)
					}
					selectEnterPressAfterPaste.Check()
				}
			}
		}
	}()
}
