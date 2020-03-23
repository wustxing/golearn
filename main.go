package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	_ "net/http/pprof"
	"net/url"
	"os"
	"path"
	"strings"
)

func main() {
	bytes.Compare()
	fmt.Println("Abcccccc" > "Abcccc")
}

func downloadFile(rUrl string, exts []string) error {
	ret, err := url.Parse(rUrl)
	if err != nil {
		return err
	}
	s := strings.Split(ret.Path, "/")
	//name := ret.Path[strings.LastIndex(ret.Path, "/")+1:]
	name := s[len(s)-1]

	if len(exts) > 0 {
		if !isInStrings(path.Ext(name), exts) {
			return errors.New("ext not suppot")
		}
	}

	dir := s[len(s)-2]

	res, err := http.Get(rUrl)
	if err != nil {
		return err
	}
	createDirIfNoExist(dir)
	f, err := os.Create(dir + "/" + name)
	if err != nil {
		return err
	}
	defer f.Close()
	io.Copy(f, res.Body)
	return nil
}

func isInStrings(s string, ss []string) bool {
	for _, v := range ss {
		if s == v {
			return true
		}
	}
	return false
}

func createDirIfNoExist(path string) {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(path, os.ModePerm) //  Everyone can read write and execute
			return
		}
		return
	}
}
