package fileStorage

import (
	"fmt"
	"os"
)

var Path string

func Init() (path string) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Print("Error getting home directory: ", err)
	}
	path = homeDir + "\\.TextType"
	err = checkFolderExist(path)
	if err != nil {
		panic(err)
	}
	return path
}
