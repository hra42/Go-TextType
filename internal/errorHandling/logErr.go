package errorHandling

import (
	"github.com/hra42/Go-TextType/internal/logging"
	"log"
)

func CheckError(err error) {
	if err != nil {
		if logging.ErrorLogger != nil {
			logging.ErrorLogger.Println("Error: ", err)
		} else {
			log.Fatal("Logger is empty!")
		}
	}
}
