package main

import (
	"fmt"
	"log"
	"time"

	"github.com/ogier/pflag"
	"github.com/twcamper/gopl/ch04/dateTime"
	"github.com/twcamper/gopl/ch04/github"
)

var (
	result     *github.IssuesSearchResult
	serviceUrl string
	fromIn     string
	toIn       string
	from       time.Time
	to         time.Time
)

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
		if item.CreatedAt.After(from) && item.CreatedAt.Before(to) {
			fmt.Printf("#%-5d %9.9s %s %.85s\n",
				item.Number,
				item.User.Login,
				item.CreatedAt.Format(dateTime.Format),
				item.Title)
		}
	}
}
func printHeader() {
	fmt.Printf("%d issues:\n", result.TotalCount)
}

func parseDates() {
	from = parseDate(fromIn)
	to = parseDate(toIn)
	if to.IsZero() {
		to = time.Now()
	}
}

func parseDate(s string) time.Time {
	t, err := dateTime.NewDateTime(s)
	if err != nil {
		log.Fatal(err)
	}
	return t
}
