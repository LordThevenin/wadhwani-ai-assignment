package utils

import (
	"github.com/sirupsen/logrus"
	"os"
)

var logger = logrus.New()

func InitLogger() {
	logger.SetLevel(logrus.DebugLevel)
	logger.SetOutput(os.Stdout)
}

func Logger() *logrus.Logger {
	return logger
}
