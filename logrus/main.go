package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.SetOutput(os.Stdout)

	logrus.SetLevel(logrus.WarnLevel)
}
func main() {
	logrus.WithFields(logrus.Fields{
		"animal": "dog",
		"size":   10,
	}).Info("A dog appers")

	logrus.WithFields(logrus.Fields{
		"omg":  true,
		"size": 10,
	}).Warn("A dog appers")

	contextLogger := logrus.WithFields(logrus.Fields{
		"common": "this is a common field",
		"other":  "hello ",
	})

	contextLogger.Info("sdfsf")
	contextLogger.Warn("me too")
}
