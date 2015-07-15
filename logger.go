package services

import (
	"fmt"

	f "github.com/gopherlabs/gopher-framework"
	log "github.com/gopherlabs/gopher-providers-logrus"
)

type LogProvider struct {
	log log.Logger
}

func (l *LogProvider) Register(c *f.Container, config interface{}) interface{} {
	conf := config.(f.ConfigLogger)
	l.log = *log.New()
	l.log.Level = log.Level(conf.LogLevel)
	l.log.Formatter = &log.TextFormatter{
		ForceColors:   true,
		FullTimestamp: conf.FullTimestamp,
	}
	//	l.log.Info("|   > FullTimestamp: %t ", conf.FullTimestamp)
	//	l.log.Info("|   > LogLevel: %s ", log.Level(conf.LogLevel).String())
	return l
}

func (l *LogProvider) GetKey() string {
	return f.LOGGER
}

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
