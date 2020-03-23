package main

import (
	"errors"
	"fmt"
	"github.com/0990/golearn/github/logrus/hook"
	"github.com/sirupsen/logrus"
)

func main() {
	hook.Init("xulog", 10)

	logrus.WithField("hi", 11).Info("hello")
	logrus.WithField("hi", 11).Debug("hello")
	logrus.WithField("hi", 11).WithError(errors.New("i am a error")).Warn("hello")
	logrus.Error("hello")
	var i int32 = 0
	c := 100 / i
	fmt.Println(c)
}
