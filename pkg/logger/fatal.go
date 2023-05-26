package logger

import (
	"log"
	"os"
)

// Fatal level msg, defer exit app
func (l *Log) Fatal(mes string, arg any) {

	infoLog := log.New(l.File, "FATAL\t", log.Ldate|log.Ltime)
	infoLog.Print(mes, arg)

	l.Close()
	os.Exit(0)
}
