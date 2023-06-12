package log

import (
	"log"
	"os"
	"sync"
)

const (
	ProdLog = "prod" // writing to file
	DevLog  = "dev"  // console output
)

// type of logging based on ProdLog or DevLog param
type Log struct {
	mu      sync.Mutex
	File    *os.File
	NameLog *string
}

// logging levels
type Logs interface {
	Info(arg ...any)
	Warn(arg ...any)
	Error(arg ...any)
	Fatal(arg ...any)

	Name(name string)
	Write(level string, arg ...any)
}

// Creating a logger based on type (log to file or output to console), directories and name for file type
func New(typeLog, patch, pacthFileName string) *Log {

	filePatch, err := checkFile(patch, pacthFileName)
	if err != nil {
		log.Fatal("checkFile fail:", tab,
			patch, tab,
			pacthFileName, tab,
			err,
		)
	}

	l := &Log{}

	switch typeLog {
	case ProdLog:
		f, err := os.OpenFile(filePatch, os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			log.Fatal("typeLog", tab,
				typeLog, tab,
				err)
		}
		l.File = f
	case DevLog:
		l.File = os.Stdout
	default:
		log.Fatal("logger type not specified:", tab,
			typeLog)
	}

	return l
}

// Close file if type prod
func (l *Log) Close() {
	l.File.Close()
}
