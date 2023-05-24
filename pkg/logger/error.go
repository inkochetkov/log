package logger

import (
	"fmt"
	"log"
)

// Error level msg
func (l *Log) Error(mes string, args ...any) {
	for _, arg := range args {
		mesArgs = fmt.Sprintf("%s %v", mesArgs, arg)
	}
	message = mes + mesArgs
	infoLog := log.New(l.File, "ERROR\t", log.Ldate|log.Ltime)
	infoLog.Print(message)

}
