package hotKeyConfig

import (
	"encoding/gob"
	"github.com/hra42/Go-TextType/internal/errorHandling"
	"github.com/hra42/Go-TextType/internal/logging"
	"golang.design/x/hotkey"
	"os"
)

func LoadHK() {
	file, err := os.Open("hotkey.gob")
	if err != nil {
		if os.IsNotExist(err) {
			logging.Logger.Println("No hotkey.gob file found")
			HK = hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModShift}, hotkey.KeyV)
			RegisterHotKey(HK)
			err = SaveLastUsedHK(1)
			errorHandling.CheckError(err)
			return
		}
		logging.Logger.Println("Error: ", err)
		errorHandling.CheckError(err)
	}
	defer CloseFile(file)

	dec := gob.NewDecoder(file)
	err = dec.Decode(&HotKeyConfiguration)
	errorHandling.CheckError(err)
	logging.Logger.Println("hotkey.gob loaded and decoded")
	switch HotKeyConfiguration.HotkeyNumber {
	case 1:
		logging.Logger.Println("Hotkey 1 is in use")
		HK = hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModShift}, hotkey.KeyV)
		RegisterHotKey(HK)
	case 2:
		logging.Logger.Println("Hotkey 2 is in use")
		HK = hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModAlt}, hotkey.KeyS)
		RegisterHotKey(HK)
	case 3:
		logging.Logger.Println("Hotkey 3 is in use")
		HK = hotkey.New([]hotkey.Modifier{hotkey.ModCtrl}, hotkey.KeyQ)
		RegisterHotKey(HK)
	}
}
