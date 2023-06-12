package log

// Info level msg
func (l *Log) Info(arg ...any) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.Write(infoLevel, arg...)

}
