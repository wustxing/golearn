package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"encoding/json"
	"fmt"
	"time"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.SetOutput(os.Stdout)

	logrus.SetLevel(logrus.WarnLevel)

	logrus.AddHook(&DefaultFieldsHook{})
}

var toWrite chan *logrus.Entry

type Hook interface{
	Levels()[]logrus.Level
	Fire(entry *logrus.Entry) error
}

type DefaultFieldsHook struct{
}
func(df *DefaultFieldsHook)Fire(entry *logrus.Entry)error{
	entry.Data["appName"] = "xujialong"

	//newEntity:=*entry
	//toWrite<-&newEntity
	toWrite<-entry
	return nil
}

func(df *DefaultFieldsHook)Levels()[]logrus.Level{
	return logrus.AllLevels
}

func main() {

	toWrite = make(chan *logrus.Entry, 100)

	go func(){
		for entry:=range toWrite{
			data, err := json.Marshal(entry.Data)
			if err==nil{
				fmt.Println("towrite",string(data))
			}
		}
	}()

	time.Sleep(time.Second*2)

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

	time.Sleep(time.Second*10)

}
