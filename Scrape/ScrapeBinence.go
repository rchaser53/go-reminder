package goReminder

import (
	"bytes"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ScrapeBinanceRegistratoinForm(targetWord string) {
	doc, err := goquery.NewDocument("https://www.binance.com/register.html")
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("form").Each(func(i int, s *goquery.Selection) {
		header := s.Find("h3 span").Text()

		if strings.Contains(header, targetWord) {
			println(28)
		}
	})
}

func PostSlack(url string) error {
	jsonStr := `{"text":"` + "nyan" + `"}`

	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer([]byte(jsonStr)),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return err
}
