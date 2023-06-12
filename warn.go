package log

// Warn level msg
func (l *Log) Warn(arg ...any) {

	l.mu.Lock()
	defer l.mu.Unlock()

	l.Write(warnLevel, arg...)

}
