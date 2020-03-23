package main

import (
	"github.com/natefinch/lumberjack"
	"log"
)

func main() {
	log.SetOutput(&lumberjack.Logger{
		Filename:   "lumberjack.log",
		MaxSize:    1,
		MaxAge:     1,
		MaxBackups: 10,
		LocalTime:  true,
		Compress:   false,
	})
	for i := 0; i < 300000; i++ {
		log.Println("hello")
	}

}
