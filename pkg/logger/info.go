package logger

import (
	"log"
)

// Info level msg
func (l *Log) Info(mes string, arg any) {

	infoLog := log.New(l.File, "INFO\t", log.Ldate|log.Ltime)
	infoLog.Print(mes, arg)

}
