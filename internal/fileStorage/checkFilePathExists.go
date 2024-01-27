package fileStorage

import (
	"fmt"
	"os"
)

func checkFolderExist(path string) (err error) {
	// Check if the directory exists
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		fmt.Println("Directory does not exist:", err)
		errDir := os.MkdirAll(path, 0755)
		if errDir != nil {
			fmt.Println("Error creating directory:", err)
			return err
		}
	} else if err != nil {
		fmt.Println("Directory exists.")
		fmt.Println("Other error:", err)
		return err
	}
	return nil
}
