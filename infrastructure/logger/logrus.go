package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

func init() {
	Logger = NewLogger("main")
}

var Logger *logrus.Entry

func NewLogger(name string) *logrus.Entry {
	log := logrus.New()
	log.SetReportCaller(true)
	log.SetOutput(os.Stdout)
	return log.WithFields(logrus.Fields{
		"package": name,
	})
}
