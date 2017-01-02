package main

import (
	"fmt"

	"github.com/twcamper/gopl/ch05/htmlutil"
	"golang.org/x/net/html"
)

func main() {
	doc := htmlutil.Read()
	outline(nil, doc)
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}
