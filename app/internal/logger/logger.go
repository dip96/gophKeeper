package logger

import (
	"github.com/sirupsen/logrus"
	configServer "gophKeeper/internal/config"
)

type Logger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
}

type logger struct {
	log *logrus.Logger
}

func Init() (Logger, error) {
	l := logrus.New()
	l.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	cnf := configServer.MustLoad()
	level, err := logrus.ParseLevel(cnf.GetLoggerLevel())

	if err != nil {
		return nil, err
	}

	l.SetLevel(level)

	return &logger{log: l}, nil
}

func (l *logger) Debug(args ...interface{}) {
	l.log.Debug(args...)
}

func (l *logger) Debugf(format string, args ...interface{}) {
	l.log.Debugf(format, args...)
}

func (l *logger) Info(args ...interface{}) {
	l.log.Info(args...)
}

func (l *logger) Infof(format string, args ...interface{}) {
	l.log.Infof(format, args...)
}

func (l *logger) Error(args ...interface{}) {
	l.log.Error(args...)
}

func (l *logger) Errorf(format string, args ...interface{}) {
	l.log.Errorf(format, args...)
}
