package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func main(){
	viper.SetConfigType("json")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		log.Panicln(err)
	}

	fmt.Println(err)
}
