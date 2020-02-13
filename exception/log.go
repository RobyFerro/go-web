package exception

import (
	"fmt"
	"log"
	"os"
)

// Log an error into the main error.log file
func Log(message string) {
	logFile, err := os.OpenFile("storage/log/error.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err.Error())
	}

	logger := log.New(logFile, "", log.LstdFlags)
	logger.Println(message)
}
