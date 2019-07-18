package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type readerGetter interface {
	get(path string)
}

type fileReaderGetter struct {
	path string
}

func (f *fileReaderGetter) get() io.Reader {
	file, err := os.Open(f.path)
	check(err)
	return file
}

type urlReaderGetter struct {
	path string
}

func (u *urlReaderGetter) get() io.Reader {
	var err error

	client := &http.Client{}

	req, err := http.NewRequest("GET", u.path, nil)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("do client err->%v", err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalf("read resp err->%v", err)
	}

	reader := strings.NewReader(string(body))

	return reader
}
