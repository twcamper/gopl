// Exercise 5.2: Write a function to populate a mapping from element names to the number of elements with that name in an HTML document tree.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	tagCounts := make(map[string]int)
	tagMap(tagCounts, read())
	for tag, count := range tagCounts {
		fmt.Printf("%s:\t%d\n", tag, count)
	}
}

func tagMap(tagCounts map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		tagCounts[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		tagMap(tagCounts, c)
	}
}

func read() *html.Node {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		os.Exit(1)
	}
	return doc
}
