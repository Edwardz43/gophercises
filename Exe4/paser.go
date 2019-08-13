package main

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

// Link ...
type Link struct {
	Href string `json:"href"`
	Text string `json:"text"`
}

func parseHTMLNode(n *html.Node) {
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
		parseHTMLNode(c)
	}
}
