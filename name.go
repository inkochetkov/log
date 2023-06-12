package log

// Info level msg
func (l *Log) Name(thisName string) {
	l.NameLog = &thisName
}
