package log

import (
	"github.com/sirupsen/logrus"
)

var l *logrus.Logger

func init() {
	l = logrus.New()
}

func Get() *logrus.Logger {
	return l
}
