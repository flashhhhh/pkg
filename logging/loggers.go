package logging

var Loggers = make(map[string]*Logger)

func InitLogger(serviceName, filename string, maxSize, maxBackups, maxAge int) {
	logger := NewLogger(filename, maxSize, maxBackups, maxAge)
	Loggers[serviceName] = logger
}

func GetLogger(serviceName string) *Logger {
	if logger, exists := Loggers[serviceName]; exists {
		return logger
	}
	return nil
}

func LogMessage(serviceName, message, level string) {
	logger := GetLogger(serviceName)
	if logger != nil {
		logger.Log(message, level)
	}
}