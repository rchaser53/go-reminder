package main

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var TargetWord string = "Registrations Pause"

func ExampleScrape() {
	doc, err := goquery.NewDocument("https://www.binance.com/register.html")
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("form").Each(func(i int, s *goquery.Selection) {
		header := s.Find("h3 span").Text()

		if strings.Contains(header, TargetWord) {
			println(28)
		}
	})
}

func main() {
	ExampleScrape()
}
