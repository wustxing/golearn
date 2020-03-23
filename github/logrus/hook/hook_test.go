package hook

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"testing"
)

func Test_Hook(t *testing.T) {
	Init("xulog", 10)

	logrus.WithField("hi", 11).Info("hello")
	logrus.WithField("hi", 11).Debug("hello")
	logrus.WithField("hi", 11).WithError(errors.New("i am a error")).Warn("hello")
	logrus.Error("hello")
	logrus.Panic("hello")
	var i int32 = 0
	c := 100 / i
	fmt.Println(c)
}
