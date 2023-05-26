package logger

import (
	"log"
)

// Error level msg
func (l *Log) Error(mes string, arg any) {

	infoLog := log.New(l.File, "ERROR\t", log.Ldate|log.Ltime)
	infoLog.Print(mes, arg)

}
