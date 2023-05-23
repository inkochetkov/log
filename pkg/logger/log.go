package logger

import (
	"log"
	"os"
)

const (
	ProdLog = "prod" // writing to file
	DevLog  = "dev"  // console output
)

// type of logging based on ProdLog or DevLog param
type Log struct {
	File *os.File
}

// logging levels
type Logs interface {
	Info(mes string)
	Warn(mes string)
	Error(mes string, err error)
	Fatal(mes string, err error)
}

// Creating a logger based on type (log to file or output to console), directories and name for file type
func New(typeLog, patch, pacthFileName string) *Log {

	filePatch, err := checkFile(patch, pacthFileName)
	if err != nil {
		log.Fatal("checkFile", err)
	}

	l := &Log{}

	switch typeLog {
	case ProdLog:
		f, err := os.OpenFile(filePatch, os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			log.Fatal(err)
		}
		l.File = f
	case DevLog:
		l.File = os.Stdout
	default:
		log.Fatal("logger type not specified")
	}

	return l
}

// Close file if type prod
func (l *Log) Close() {
	l.File.Close()
}
