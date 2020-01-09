package helper

import (
	"ikdev/go-web/exception"
	"log"
	"os"
)

func Log(message string) {
	logFile, err := os.OpenFile("storage/log/error.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		exception.ProcessError(err)
	}

	logger := log.New(logFile, "", log.LstdFlags)
	logger.Println(message)
}
