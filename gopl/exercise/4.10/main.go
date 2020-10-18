package main

import (
	"flag"
	"fmt"
	"github.com/eanson023/golearning/gopl/ch4/github"
	"log"
	"os"
	"time"
)

var interval string

const (
	Month = "month"
	Year  = "year"
)

func init() {
	flag.StringVar(&interval, "i", "year", "month year")
}

func main() {
	flag.Parse()
	t := time.Now()
	switch interval {
	case Month:
		t = t.AddDate(0, -1, 0)
	case Year:
		t = t.AddDate(-1, 0, 0)
	default:
		log.Fatal("error input")
	}
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		if item.CreateAt.After(t) {
			fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
		}
	}
}
