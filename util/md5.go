package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"sort"
	"sync"
	"time"
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

func MD5FileSync(files ...string) (string, error) {
	length := len(files)
	if length == 0 {
		return "", errors.New("no file")
	}

	md5Map := make(map[string][]byte)
	var paths []string
	for _, v := range files {
		file := v
		data, err := ioutil.ReadFile(file)
		if err != nil {
			return "", err
		}
		md5Map[file] = data
		paths = append(paths, file)
	}

	sort.Strings(paths)
	var allData []byte
	for _, v := range paths {
		allData = append(allData, md5Map[v]...)
	}
	return fmt.Sprintf("%x", md5.Sum(allData)), nil
}

func MD5FileSyncFast(files ...string) (string, error) {
	now := time.Now()
	defer func() {
		duration := time.Since(now)
		fmt.Println("async", duration)
	}()
	length := len(files)
	if length == 0 {
		return "", errors.New("no file")
	}

	md5Map := make(map[string][16]byte)
	var paths []string
	for _, v := range files {
		file := v
		data, err := ioutil.ReadFile(file)
		if err != nil {
			return "", err
		}
		md5Map[file] = md5.Sum(data)
		paths = append(paths, file)
	}

	sort.Strings(paths)
	var allData []byte
	for _, v := range paths {
		md5file := md5Map[v]
		allData = append(allData, md5file[:]...)
	}
	//fmt.Println("append after", duration)
	return fmt.Sprintf("%x", md5.Sum(allData)), nil
}

func MD5FileAsync(files ...string) (string, error) {
	now := time.Now()
	defer func() {
		duration := time.Since(now)
		fmt.Println("async", duration)
	}()
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
			now := time.Now()
			defer func() {
				duration := time.Since(now)
				fmt.Println(duration)
			}()
			data, err := ioutil.ReadFile(file)
			//md5Value, err := Md5File(file)
			resultChan <- result{file, data, err}
			wg.Done()
		}()
	}
	wg.Wait()
	duration := time.Since(now)
	fmt.Println("wait after", duration)
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
	duration = time.Since(now)
	fmt.Println("append after", duration)
	return fmt.Sprintf("%x", md5.Sum(allData)), nil
}

func MD5FileAsyncFast(files ...string) (string, error) {
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
			md5 := md5.Sum(data)
			resultChan <- result{file, md5[:], err}
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
