package logger

import (
	"io"
	"log"
	"os"

	"github.com/apsdehal/go-logger"
)

const log_file = "log.log"

var Logger *logger.Logger

func init() {
	logFile, err := os.OpenFile(log_file, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }
	writer := io.MultiWriter(os.Stdout, logFile)
	
	log, err := logger.New("main", 1, writer)
	if err != nil {
		log.Fatal(err.Error())
	}

	Logger = log
}
