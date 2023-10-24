package errorHandling

import (
	"github.com/hra42/Go-TextType/internal/logging"
	"log"
)

func CheckError(err error) {
	if err != nil {
		if logging.Logger != nil {
			logging.Logger.Println("Error: ", err)
		} else {
			log.Fatal("Logger is empty!")
		}
	}
}
