package log

import (
	"log"
)

func (l *Log) Write(level string, arg ...any) {

	infoLog := log.New(l.File, level, log.Ldate|log.Ltime)

	var args []any

	if l.NameLog != nil {
		args = append(args, *l.NameLog, tab)
	}

	for _, ar := range arg {
		args = append(args, ar, tab)
	}

	infoLog.Print(args...)

}
