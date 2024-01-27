package hotKeyConfig

import (
	"encoding/gob"
	"github.com/hra42/Go-TextType/internal/fileStorage"
	"os"
)

func SaveLastUsedHK(hkNumber int) error {
	filePath := fileStorage.Path + "/hotkey.gob"
	file, err := os.Create(filePath)
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
