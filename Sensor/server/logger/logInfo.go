package logger

import (
	"errors"
	"os"

	"log"
)

type Logger struct {
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
}

func GetInstance() *Logger {
	return loggers
}
func New() (*Logger, error) {
	file, err := os.OpenFile("sensor.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil, errors.New("cannot open file")
	}
	return &Logger{
		WarningLogger: log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile),
		InfoLogger:    log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		ErrorLogger:   log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}, err
}

var loggers *Logger
var err error

func init() {
	loggers, err = New()
	if err != nil {
		log.Fatal(err)
	}
}
