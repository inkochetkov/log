package logger

import "log"

// Warn level msg
func (l *Log) Warn(mes string) {
	infoLog := log.New(l.File, "WARN\t", log.Ldate|log.Ltime)
	infoLog.Print(mes)

}
