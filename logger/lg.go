package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

func date() string {
	now := time.Now()
	y, m, d := now.Date()
	return fmt.Sprintf("%d-%d-%d.log", y, int(m), d)
}

type Logger struct {
	Type     string
	FileName string
	INFO     func(v ...interface{})
	WARNING  func(v ...interface{})
	ERROR    func(v ...interface{})
}

func New() *Logger {

	err := os.Mkdir("logs", os.ModeDir)

	if err != nil {
		println(err.Error())
	}

	filename := date()
	path := "logs/" + filename
	file, oerr := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if oerr != nil {
		panic(oerr.Error())
	}

	INFO := log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WARNING := log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ERROR := log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	return &Logger{
		FileName: filename,
		INFO:     INFO.Println,
		WARNING:  WARNING.Println,
		ERROR:    ERROR.Println,
	}

}
