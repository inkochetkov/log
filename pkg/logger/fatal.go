package logger

import (
	"fmt"
	"log"
	"os"
)

// Fatal level msg, defer exit app
func (l *Log) Fatal(mes string, args ...any) {

	for _, arg := range args {
		mesArgs = fmt.Sprintf("%s %v", mesArgs, arg)
	}
	message = mes + mesArgs

	infoLog := log.New(l.File, "FATAL\t", log.Ldate|log.Ltime)
	infoLog.Print(message)

	l.Close()
	os.Exit(0)
}
