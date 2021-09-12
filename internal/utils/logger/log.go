package logger

import (
	"encoding/json"
	"fmt"
	"log"
)

func execute(level string, message string) {
	log.Print(fmt.Sprintf("\n[%s] %s", level, message))
}

func Error(format string, args ...interface{}) {
	execute("ERROR", fmt.Sprintf(format, args...))
}

func Warn(format string, args ...interface{}) {
	execute("WARN", fmt.Sprintf(format, args...))
}

func Info(format string, args ...interface{}) {
	execute("INFO", fmt.Sprintf(format, args...))
}

func InfoObject(message string, object interface{}) {
	execute("INFO", message)
	parsed, _ := json.Marshal(object)
	execute("INFO", string(parsed))
}

func Debug(format string, args ...interface{}) {
	execute("DEBUG", fmt.Sprintf(format, args...))
}
