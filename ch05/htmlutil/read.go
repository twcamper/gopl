package htmlutil

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func Read() *html.Node {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "html.Read: %v\n", err)
		os.Exit(1)
	}
	return doc
}
