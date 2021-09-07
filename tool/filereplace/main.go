package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var source = flag.String("src", "", "souce file")
var dst = flag.String("dst", "", "dst file")

type arrayFlags []string

// Value ...
func (i *arrayFlags) String() string {
	return fmt.Sprint(*i)
}

// Set 方法是flag.Value接口, 设置flag Value的方法.
// 通过多个flag指定的值， 所以我们追加到最终的数组上.
func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func main() {
	var replaces arrayFlags
	flag.Var(&replaces, "replace", "replace string old new")

	flag.Parse()
	data, err := ioutil.ReadFile(*source)
	if err != nil {
		log.Fatalln(err)
	}

	str := string(data)
	if replaces.String() != "" {
		for _, v := range replaces {
			oldnew := strings.Split(v, ":")
			if len(oldnew) != 2 {
				continue
			}

			str = strings.ReplaceAll(str, oldnew[0], oldnew[1])
		}
	}

	err = ioutil.WriteFile(*dst, []byte(str), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}
