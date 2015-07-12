package services

import (
	"fmt"

	"github.com/gopherlabs/gopher-framework"
	log "github.com/gopherlabs/gopher-providers-logrus"
)

type LogProvider struct {
	log log.Logger
}

func (l *LogProvider) Register(config map[string]interface{}) interface{} {
	l.log = *log.New()
	l.log.Level = log.DebugLevel
	l.log.Formatter = &log.TextFormatter{
		ForceColors:   true,
		FullTimestamp: config["FullTimestamp"].(bool),
	}
	return l
}

func (l *LogProvider) GetKey() string {
	return framework.LOGGER
}

//func (l *LogProvider) NewLog() framework.Loggable {
//	return l
//}

func (l *LogProvider) Info(msg string, args ...interface{}) {
	l.log.Info(sprintf(msg, args...))
}

func (l *LogProvider) Debug(msg string, args ...interface{}) {
	l.log.Debug(sprintf(msg, args...))
}

func (l *LogProvider) Warn(msg string, args ...interface{}) {
	l.log.Warn(sprintf(msg, args...))
}

func (l *LogProvider) Error(msg string, args ...interface{}) {
	l.log.Error(sprintf(msg, args...))
}

// Calls os.Exit(1) after logging
func (l *LogProvider) Fatal(msg string, args ...interface{}) {
	l.log.Fatal(sprintf(msg, args...))
}

// Calls panic() after logging
func (l *LogProvider) Panic(msg string, args ...interface{}) {
	l.log.Panic(sprintf(msg, args...))
}

func sprintf(msg string, args ...interface{}) string {
	if args == nil {
		return msg
	}
	return fmt.Sprintf(msg, args...)
}
