/*
  Prints the links in an HTML document read from stdin
  Exercise 5.1: Change the findlinks program to traverse the n.FirstChild linked list using recursion calls to _visit_ instead of a loop.
*/
package main

import (
	"fmt"

	"github.com/twcamper/gopl/ch05/htmlutil"
	"golang.org/x/net/html"
)

func main() {
	doc := htmlutil.Read()
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

// appends to links each link found in n and returns the result
func visit(links []string, n *html.Node) []string {
	if n != nil {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					links = append(links, a.Val)
				}
			}
		}
		links = visit(links, n.FirstChild)
		links = visit(links, n.NextSibling)
	}
	return links
}
