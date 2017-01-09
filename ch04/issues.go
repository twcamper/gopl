package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/twcamper/gopl/ch04/github"
)

var serviceUrl string

func init() {
	flag.StringVar(&serviceUrl, "u", "", "alternate url")
	flag.StringVar(&serviceUrl, "url", "", "alternate url")
}

func main() {
	flag.Parse()
	result, err := github.SearchIssues(url(), os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}

func url() string {
	if serviceUrl == "" {
		return github.IssuesURL
	}
	return serviceUrl
}
