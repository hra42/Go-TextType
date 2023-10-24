package logging

import (
	"fmt"
	"log"
	"os"
)

var Logger *log.Logger

var LogFile *os.File

func SetupLogger() *log.Logger {
	// Open the log file
	var err error
	LogFile, err = os.OpenFile("TextType.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Failed to open log file:", err)
		return nil
	}

	// Create a new logger
	Logger = log.New(LogFile, "TextType", log.LstdFlags)
	if Logger == nil {
		fmt.Println("Failed to create new logger")
		return nil
	}

	return Logger
}
