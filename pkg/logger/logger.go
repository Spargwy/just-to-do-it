package logger

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const stackLevel = 2

func New(level string) {
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		lvl = logrus.InfoLevel
	}
	logrus.SetLevel(lvl)

	logrus.SetFormatter(&logrus.
		TextFormatter{
		FullTimestamp:          true,
		DisableLevelTruncation: true,
		PadLevelText:           true,
		DisableQuote:           true,
	})
}

type stackTracer interface {
	StackTrace() errors.StackTrace
}

// Debug prints error with stacktrace.
func Debug(err error) {
	if stack, ok := err.(stackTracer); ok {
		st := stack.StackTrace()
		if len(st) > stackLevel {
			st = st[:stackLevel]
		}
		logrus.Debugf("%s%+v", err, st)
	} else {
		logrus.Debug(err)
	}
}

func Warning(warning string, args ...interface{}) {
	logrus.Warningf(warning, args...)
}

func Info(info string, args ...interface{}) {
	logrus.Infof(info, args...)
}
