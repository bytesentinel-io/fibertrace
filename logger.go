// Part of ByteSentinel.io - https://bytesentinel.io

package fibertrace

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"time"
)

type LogEntry struct {
	Timestamp   time.Time `json:"timestamp"`
	Level       string    `json:"level"`
	Application string    `json:"application"`
	Module      string    `json:"module"`
	Message     string    `json:"message"`
	User        string    `json:"user"`
}

type Logger struct {
	logger      *log.Logger
	file        *os.File
	application string
	jsonFormat  bool
}

func NewLogger(logFilePath, application string, jsonFormat bool) (*Logger, error) {
	logger := log.New(os.Stdout, "", 0)
	var file *os.File
	var err error

	if logFilePath != "" {
		file, err = openLogFile(logFilePath)
		if err != nil {
			return nil, err
		}
		logger.SetOutput(io.MultiWriter(logger.Writer(), file))
	}

	return &Logger{
		logger:      logger,
		file:        file,
		application: application,
		jsonFormat:  jsonFormat,
	}, nil
}

func (l *Logger) Info(message string) {
	l.log("INFO", message)
}

func (l *Logger) Error(message string) {
	l.log("ERROR", message)
}

func (l *Logger) Debug(message string) {
	l.log("DEBUG", message)
}

func (l *Logger) Errorf(message string) {
	l.log("ERROR", message)
	os.Exit(1)
}

func (l *Logger) log(level, message string) {
	_, file, _, _ := runtime.Caller(2)
	module := filepath.Base(file)

	user, err := userLookup()
	if err != nil {
		l.logger.Println("Failed to determine the current user:", err)
		return
	}

	logEntry := LogEntry{
		Timestamp:   time.Now().UTC(),
		Level:       level,
		Application: l.application,
		Module:      module,
		Message:     message,
		User:        user,
	}

	if l.jsonFormat {
		l.logJSON(logEntry)
	} else {
		l.logText(logEntry)
	}
	l.console(logEntry)
}

func (l *Logger) logJSON(logEntry LogEntry) {
	logJSON, err := json.Marshal(logEntry)
	if err != nil {
		l.logger.Println("Failed to marshal log entry:", err)
		return
	}

	l.writeToFile(logJSON)
}

func (l *Logger) logText(logEntry LogEntry) {
	timestamp := logEntry.Timestamp.Format("2006/01/02 15:04:05")
	logText := fmt.Sprintf("%s %s [%s] (%s:%s): %s", timestamp, logEntry.User, logEntry.Level, l.application, logEntry.Module, logEntry.Message)

	l.writeToFile([]byte(logText))
}

func (l *Logger) console(logEntry LogEntry) {
	timestamp := logEntry.Timestamp.Format("2006/01/02 15:04:05")
	logText := fmt.Sprintf("%s %s [%s] (%s:%s): %s", timestamp, logEntry.User, logEntry.Level, l.application, logEntry.Module, logEntry.Message)

	fmt.Println(logText)
}

func (l *Logger) writeToFile(logData []byte) {
	if l.file != nil {
		_, err := l.file.Write(logData)
		if err != nil {
			l.logger.Println("Failed to write log entry to file:", err)
		}
		_, err = l.file.WriteString("\n")
		if err != nil {
			l.logger.Println("Failed to write log entry to file:", err)
		}
		l.file.Sync()
	}
}

func openLogFile(logFilePath string) (*os.File, error) {
	absPath, err := filepath.Abs(logFilePath)
	if err != nil {
		return nil, err
	}
	return os.OpenFile(absPath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
}

func userLookup() (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", err
	}
	return currentUser.Username, nil
}
