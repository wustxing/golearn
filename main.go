package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func main() {
	var offset int64
	of, err := os.OpenFile("offset.txt", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer of.Close()

	ifdata, err := ioutil.ReadAll(of)
	if err != nil {
		log.Fatal(err)
	}

	if len(ifdata) > 0 {
		offset, err = strconv.ParseInt(string(ifdata), 10, 64)
		if err != nil {
			log.Fatal(err)
		}
	}

	f, err := os.OpenFile("user.txt", os.O_RDONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = f.Seek(offset, io.SeekStart)
	if err != nil {
		log.Fatal(err)
	}

	//var currOffset int64

	br := bufio.NewReader(f)
	var count int32
	for {
		count++
		if count == 3 {
			break
		}
		line, length, c := NextLine(br)

		//s, _, c := br.ReadLine()
		if c == io.EOF {
			fmt.Println("eof")
			return
		}

		//fmt.Println(string(s))
		fmt.Println(string(line), length)
		//fOffset, _ := f.Seek(0, io.SeekCurrent)
		//currOffset = fOffset - int64(br.Buffered())
		offset += int64(length)
		WriteOffset(of, offset)
	}
}

func WriteOffset(f *os.File, offset int64) error {
	err := f.Truncate(0)
	if err != nil {
		return err
	}
	_, err = f.Seek(0, os.SEEK_SET)
	if err != nil {
		return err
	}
	_, err = f.WriteString(strconv.FormatInt(offset, 10))
	if err != nil {
		return err
	}
	return nil
}

func NextLine(b *bufio.Reader) (line []byte, lenWithDelim int, err error) {
	data, err := b.ReadBytes('\n')

	if len(data) == 0 {
		return nil, 0, err
	}

	if data[len(data)-1] == '\n' {
		drop := 1
		if len(data) > 1 && data[len(data)-2] == '\r' {
			drop = 2
		}
		length := len(data)
		return data[:length-drop], length, nil
	}

	return data, len(data), nil
}
