package walrus

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/sirupsen/logrus"
)

// Force implementation check.
var _ watermill.LoggerAdapter = &Logger{}

type Logger struct {
	entry logrus.FieldLogger
}

func New() watermill.LoggerAdapter {
	return &Logger{
		entry: logrus.StandardLogger(),
	}
}

func NewWithLogger(l logrus.FieldLogger) watermill.LoggerAdapter {
	return &Logger{
		entry: l,
	}
}

func (l *Logger) Error(msg string, err error, fields watermill.LogFields) {
	l.entry.WithFields(logrus.Fields(fields)).Error(msg, err)
}

func (l *Logger) Info(msg string, fields watermill.LogFields) {
	l.entry.WithFields(logrus.Fields(fields)).Info(msg)
}

func (l *Logger) Debug(msg string, fields watermill.LogFields) {
	l.entry.WithFields(logrus.Fields(fields)).Debug(msg)
}

func (l *Logger) Trace(msg string, fields watermill.LogFields) {
	l.entry.WithFields(logrus.Fields(fields)).Trace(msg)
}

func (l *Logger) With(fields watermill.LogFields) watermill.LoggerAdapter {
	return &Logger{
		entry: l.entry.WithFields(logrus.Fields(fields)),
	}
}
