package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/twcamper/gopl/ch04/github"
)

var serviceUrl string
var result *github.IssuesSearchResult

func init() {
	flag.StringVar(&serviceUrl, "u", "", "alternate url")
	flag.StringVar(&serviceUrl, "url", "", "alternate url")
}

func main() {
	flag.Parse()
	get()
	printHeader()
	printRows()
}

func get() {
	var err error
	result, err = github.SearchIssues(url(), os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
}

func url() string {
	if serviceUrl == "" {
		return github.IssuesURL
	}
	return serviceUrl
}

func printRows() {
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %s %.85s\n",
			item.Number,
			item.User.Login,
			item.CreatedAt.Format("2006-01-02 15:04:05"),
			item.Title)
	}
}
func printHeader() {
	fmt.Printf("%d issues:\n", result.TotalCount)
}
