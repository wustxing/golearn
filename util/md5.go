package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"sort"
	"sync"
)

type result struct {
	path string
	data []byte
	err  error
}

func Md5File(file string) (string, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", md5.Sum(data)), nil
}

func MD5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	cipherStr := h.Sum(nil)
	ret := hex.EncodeToString(cipherStr)
	return ret
}

func Md5Files(files ...string) (string, error) {
	length := len(files)
	if length == 0 {
		return "", errors.New("no file")
	}
	resultChan := make(chan result, length)
	wg := sync.WaitGroup{}
	wg.Add(length)
	for _, v := range files {
		file := v
		go func() {
			data, err := ioutil.ReadFile(file)
			//md5Value, err := Md5File(file)
			resultChan <- result{file, data, err}
			wg.Done()
		}()
	}
	wg.Wait()
	close(resultChan)

	md5Map := make(map[string][]byte)
	for result := range resultChan {
		if result.err != nil {
			return "", result.err
		}
		md5Map[result.path] = result.data
	}

	var paths []string
	for path := range md5Map {
		paths = append(paths, path)
	}
	sort.Strings(paths)

	var allData []byte

	for _, v := range paths {
		allData = append(allData, md5Map[v]...)
	}
	return fmt.Sprintf("%x", md5.Sum(allData)), nil
}
