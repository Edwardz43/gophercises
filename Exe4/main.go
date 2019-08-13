package main

import (
	"bufio"
	"log"

	"golang.org/x/net/html"
)

func check(e error) {
	if e != nil {
		log.Panic(e)
	}
}

func main() {
	reader := fileReaderGetter{path: "ex5.html"}

	// reader := urlReaderGetter{path: "https://yami.io/golang-interface/"}

	n, err := html.Parse(bufio.NewReader(reader.get()))
	check(err)

	parseHTMLNode(n)
}
