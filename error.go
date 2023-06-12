package log

// Error level msg
func (l *Log) Error(arg ...any) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.Write(errorLevel, arg...)

}
