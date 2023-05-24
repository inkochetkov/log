package logger

import (
	"fmt"
	"log"
)

// Info level msg
func (l *Log) Info(mes string, args ...any) {

	for _, arg := range args {
		mesArgs = fmt.Sprintf("%s %v", mesArgs, arg)
	}
	message = mes + mesArgs

	infoLog := log.New(l.File, "INFO\t", log.Ldate|log.Ltime)
	infoLog.Print(message)

}
