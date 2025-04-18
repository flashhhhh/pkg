package logging

import (
	"log"
	"time"

	"github.com/natefinch/lumberjack"
)

type Logger struct {
	Writer *lumberjack.Logger
}

func NewLogger(filename string, maxSize, maxBackups, maxAge int) *Logger {
	writer := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,    // megabytes
		MaxBackups: maxBackups, // number of backups
		MaxAge:     maxAge,     // days
	}

	return &Logger{
		Writer: writer,
	}
}

func (l *Logger) Log(message string, level string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logMessage := "[" + timestamp + "] [" + level + "] " + message
	l.Writer.Write([]byte(logMessage + "\n"))

	if (level != "DEBUG") {
		log.Println(logMessage)
	}
}