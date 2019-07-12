package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func check(e error) {
	if e != nil {
		log.Panic(e)
	}
}

// Link ...
type Link struct {
	Href string `json:"href"`
	Text string `json:"text"`
}

var f func(*html.Node)

func main() {
	file, err := os.Open("ex3.html")
	defer file.Close()
	check(err)

	n, err := html.Parse(bufio.NewReader(file))
	check(err)

	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					// fmt.Println(a.Val)
					s := strings.TrimPrefix(n.FirstChild.Data, "\n")
					s = strings.TrimSpace(n.FirstChild.Data)
					l := Link{
						Href: a.Val,
						Text: s,
					}
					fmt.Printf("Link{\n  Href : %s\n  Text : %s\n}\n", l.Href, l.Text)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(n)
}

// func preProcess(s string) string {

// }
