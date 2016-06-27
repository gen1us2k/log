package log

// NewLogger creates new logger instance
func NewLogger(name string) Logger {
	me.Lock()
	defer me.Unlock()

	log, ok := channels[name]
	if !ok {
		log = newChannel(name, defaultLevel)
		channels[name] = log
	}

	return log
}

// RemoveLogger removes logger
func RemoveLogger(name string) {
	me.Lock()
	defer me.Unlock()
	delete(channels, name)
}

// Panic is equivalent to Print() followed by a call to panic().
func Panic(format string, v ...interface{}) {
	logger.Panic(format, v...)
}

// Fatal is equivalent to Print() followed by a call to os.Exit(1).
func Fatal(format string, v ...interface{}) {
	logger.Fatal(format, v...)
}

// Error is equivalent to Print() followed by error log level
func Error(format string, v ...interface{}) {
	logger.Error(format, v...)
}

// Warning is equivalent to Print() followed by warning log level
func Warning(format string, v ...interface{}) {
	logger.Warning(format, v...)
}

// Info is equivalent to Print() followed by info log level
func Info(format string, v ...interface{}) {
	logger.Info(format, v...)
}

// Debug is equivalent to Print() followed by debug log level
func Debug(format string, v ...interface{}) {
	logger.Debug(format, v...)
}

// SetLevel sets level of logger
func SetLevel(level int) {
	me.Lock()
	defer me.Unlock()

	defaultLevel = level

	for _, log := range channels {
		log.SetLevel(level)
	}
}
