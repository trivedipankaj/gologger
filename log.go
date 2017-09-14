package logger

import (
	"os"
	"time"

	uuid "github.com/satori/go.uuid"
)

const (
	OFF = iota
	FATAL
	ERROR
	WARN
	INFO
	DEBUG
	ALL
)

type Log struct {
	UID       string
	Type      string
	Level     string
	Host      string
	Timestamp string
	Data      map[string]interface{}
}

func (L *Log) SetLevel(level int64) *Log {
	var ll string
	switch level {
	case ALL:
		ll = "ALL"
	case DEBUG:
		ll = "DEBUG"
	case INFO:
		ll = "INFO"
	case WARN:
		ll = "WARN"
	case ERROR:
		ll = "ERROR"
	case FATAL:
		ll = "FATAL"
	default:
	}

	L.Level = ll
	return L
}

func NewLog(kind string, data map[string]interface{}) *Log {
	host, _ := os.Hostname()
	return &Log{UID: uuid.NewV4().String(),
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
		Type:      kind,
		Level:     "INFO",
		Host:      host,
		Data:      data}
}
