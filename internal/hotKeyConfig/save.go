package hotKeyConfig

import (
	"encoding/gob"
	"os"
)

func SaveLastUsedHK(hkNumber int) error {
	file, err := os.Create("hotkey.gob")
	if err != nil {
		return err
	}
	defer CloseFile(file)

	HotKeyConfiguration.HotkeyNumber = hkNumber
	enc := gob.NewEncoder(file)
	err = enc.Encode(HotKeyConfiguration)
	if err != nil {
		return err
	}

	return nil
}
