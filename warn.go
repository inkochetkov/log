package log

import (
	"log"
)

// Warn level msg
func (l *Log) Warn(mes string, arg any) {

	infoLog := log.New(l.File, "WARN\t", log.Ldate|log.Ltime)
	infoLog.Print(mes, arg)

}
