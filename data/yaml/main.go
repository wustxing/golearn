package main

import (
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
	"fmt"
	"path/filepath"
	"os"
)

type conf struct{
	Enabled bool `yaml:"enabled"`
	Path string `yaml:"path"`
}

func(c *conf) getConf() *conf{
	yamlFile,err:=ioutil.ReadFile("conf.yaml")
	if err!=nil{
		log.Printf("yamlFile.Get err  #%v",err)
	}
	err = yaml.Unmarshal(yamlFile,c)

	if err!=nil{
		log.Fatalf("Unmarshal:%v",err)
	}
	return c
}
func GetCurrentDir() string{
	dir,err:=filepath.Abs(filepath.Dir(os.Args[0]))
	if err!=nil{
		log.Fatal(err)
	}
	return dir
}
func main(){
	fmt.Println(GetCurrentDir())
	var c conf
	c.getConf()
	fmt.Println(c)
}
