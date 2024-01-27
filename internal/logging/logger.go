package logging

import (
	"fmt"
	"github.com/hra42/Go-TextType/internal/fileStorage"
	"log"
	"os"
)

var Logger *log.Logger
var ErrorLogger *log.Logger

var LogFile *os.File
var ErrorLogFile *os.File

func SetupLogger() *log.Logger {
	// Open the log file
	var err error
	filePath := fileStorage.Path + "/TextType.log"
	LogFile, err = os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Failed to open log file:", err)
		return nil
	}

	// Create a new logger
	Logger = log.New(LogFile, "TextType :", log.LstdFlags)
	if Logger == nil {
		fmt.Println("Failed to create new logger")
		return nil
	}

	return Logger
}

func SetupErrorLogger() *log.Logger {
	// Open the log file
	var err error
	filePath := fileStorage.Path + "/Error.log"
	ErrorLogFile, err = os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Failed to open log file:", err)
		return nil
	}

	// Create a new logger
	ErrorLogger = log.New(ErrorLogFile, "TextType Error:", log.LstdFlags)
	if ErrorLogger == nil {
		fmt.Println("Failed to create new logger")
		return nil
	}

	return ErrorLogger
}
