/* Exercise 5.3: Write a function to print the contents of all text nodes in an HTML document tree.
   Do not descend into <script> or <style> elements.
*/
package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc := read()
	textNodes(nil, doc)
}

func read() *html.Node {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "textNodes: %v\n", err)
		os.Exit(1)
	}
	return doc
}

func textNodes(stack []string, n *html.Node) {
	if n.Type == html.TextNode {
		text := strings.TrimSpace(n.Data)
		if len(text) > 0 {
			fmt.Println(text)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if !(c.Type == html.ElementNode && (c.Data == "script" || c.Data == "style")) {
			textNodes(stack, c)
		}
	}
}
