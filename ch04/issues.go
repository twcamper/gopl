package main

import (
	"fmt"
	"log"

	"github.com/ogier/pflag"
	"github.com/twcamper/gopl/ch04/dateTime"
	"github.com/twcamper/gopl/ch04/github"
)

var (
	result     *github.IssuesSearchResult
	serviceUrl string
	fromIn     string
	toIn       string
	from       *dateTime.DateTime
	to         *dateTime.DateTime
)

const dateTimeFormat string = "2006-01-02 15:04:05"

func init() {
	pflag.StringVarP(&serviceUrl, "url", "u", "", "alternate url")
	pflag.StringVarP(&fromIn, "from", "f", "", "min date 'from'")
	pflag.StringVarP(&toIn, "to", "t", "", "max date 'to'")
}

func main() {
	pflag.Parse()
	parseDates()
	get()
	printHeader()
	printRows()
}

func get() {
	var err error
	result, err = github.SearchIssues(url(), pflag.Args())
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
			item.CreatedAt.Format(dateTimeFormat),
			item.Title)
	}
}
func printHeader() {
	fmt.Printf("%d issues:\n", result.TotalCount)
}

func parseDates() {
	parse(&from, fromIn)
	parse(&to, toIn)
	fmt.Printf("From: %v\n", from)
	fmt.Printf("To: %s\n", to)
}

func parse(dt **dateTime.DateTime, s string) {
	var err error
	if len(s) > 0 {
		*dt, err = dateTime.NewDateTime(s)
		if err != nil {
			log.Fatal(err)
		}
	}
}
