package logs

import (
	"github.com/sirupsen/logrus"
	"os"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetReportCaller(true)
	logrus.SetOutput(os.Stdout)
}
