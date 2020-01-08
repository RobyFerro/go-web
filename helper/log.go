package helper

import (
	"fmt"
	"log"
	"os"
)

func Log(message string) {
	logFile, err := os.OpenFile("storage/log/error.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	logger := log.New(logFile, "", log.LstdFlags)
	logger.Println(fmt.Sprintf("# Log := %s", message))
}
