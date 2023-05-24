package logger

import (
	"fmt"
	"log"
)

// Warn level msg
func (l *Log) Warn(mes string, args ...any) {

	for _, arg := range args {
		mesArgs = fmt.Sprintf("%s %v", mesArgs, arg)
	}
	message = mes + mesArgs
	infoLog := log.New(l.File, "WARN\t", log.Ldate|log.Ltime)
	infoLog.Print(message)

}
