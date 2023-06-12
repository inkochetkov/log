package log

import (
	"os"
)

// Fatal level msg, defer exit app
func (l *Log) Fatal(arg ...any) {

	l.mu.Lock()
	defer l.mu.Unlock()

	l.Write(fatalLevel, arg...)

	l.Close()
	os.Exit(0)
}
