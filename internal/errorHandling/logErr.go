package errorHandling

import (
	"log"

	"github.com/hra42/Go-TextType/internal/logging"
)

func CheckError(err error) {
	if err != nil {
		if logging.ErrorLogger != nil {
			logging.ErrorLogger.Println("Error: ", err)
		} else {
			log.Fatal("Logger is empty! Error: ", err)
		}
	}
}
