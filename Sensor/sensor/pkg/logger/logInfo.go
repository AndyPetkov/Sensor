package logger

import (
	"errors"
	"os"

	"log"
)

type WarningLogger struct {
	WarningLogger *log.Logger
}
type ErrorLogger struct {
	ErrorLogger *log.Logger
}
type InfoLogger struct {
	InfoLogger *log.Logger
}

func newWarningLogger() (*WarningLogger, error) {
	file, err := os.OpenFile("sensor.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil, errors.New("cannot open file")
	}
	return &WarningLogger{
		WarningLogger: log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile),
	}, err
}
func newInfoLogger() (*InfoLogger, error) {
	file, err := os.OpenFile("sensor.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil, errors.New("cannot open file")
	}
	return &InfoLogger{
		InfoLogger: log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
	}, err
}
func newErrorLogger() (*ErrorLogger, error) {
	file, err := os.OpenFile("sensor.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil, errors.New("cannot open file")
	}
	return &ErrorLogger{
		ErrorLogger: log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}, err
}

var Errorloggers *ErrorLogger
var Infologgers *InfoLogger
var Warningloggers *WarningLogger
var err error

func init() {
	Errorloggers, err = newErrorLogger()
	Infologgers, err = newInfoLogger()
	Warningloggers, err = newWarningLogger()
	if err != nil {
		log.Fatal(err)
	}
}
