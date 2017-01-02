package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

var contentOnly bool

func init() {
	flag.BoolVar(&contentOnly, "c", false, "do not report meta data: only response bodies, e.g. HTML")
	flag.BoolVar(&contentOnly, "content-only", false, "do not report meta data: only response bodies, e.g. HTML")
}
func main() {
	flag.Parse()
	howMany := len(flag.Args()) - 1
	for i, arg := range flag.Args() {
		which := i + 1
		url := toURL(arg)
		response := getResponse(url.String())
		if !contentOnly {
			printHeader(response, which, howMany)
		}
		size := reportContent(response)
		if !contentOnly {
			reportMeta(response, size)
			printFooter(response, which, howMany)
		}
	}
}

func toURL(s string) *url.URL {
	_url, err := url.Parse(s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch.toURL(): url unparsable '%s'\n%v\n", s, err)
		os.Exit(1)
	}
	if _url.Scheme == "" {
		_url.Scheme = "http"
	}
	return _url
}

func getResponse(url string) *http.Response {
	r, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch.getResponse(): %v\n", err)
		os.Exit(1)
	}
	return r
}

func reportMeta(r *http.Response, size int64) {
	fmt.Printf("\nRead %d bytes\n", size)
	fmt.Println()
	fmt.Printf("%s %s\n", r.Request.Method, r.Request.URL.String())
	fmt.Println(r.Status)
	fmt.Println("\nResponse Headers:")
	for header, value := range r.Header {
		fmt.Printf("\t%s: %s\n", header, value)
	}
	fmt.Println()
}

func reportContent(r *http.Response) int64 {
	size, err := io.Copy(os.Stdout, r.Body)
	r.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch.reportContent(): reading %s: %v\n", r.Request.URL.String(), err)
		os.Exit(1)
	}
	return size
}

func printHeader(r *http.Response, i int, total int) {
	fmt.Printf("******************** BEGIN %s %s (%d of %d) ********************\n", r.Request.Method, r.Request.URL.String(), i, total)
}

func printFooter(r *http.Response, i int, total int) {
	fmt.Printf("******************** END %s %s (%d of %d) *******************\n\n", r.Request.Method, r.Request.URL.String(), i, total)
}
