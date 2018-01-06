package goReminder

import (
	"bytes"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func IsBinanceReStart(targetWord string) bool {
	doc, err := goquery.NewDocument("https://www.binance.com/register.html")
	if err != nil {
		log.Fatal(err)
	}

	registrationWord := doc.Find("form h3 span").Text()
	return strings.Contains(registrationWord, targetWord) == false
}

func PostSlack(url string) error {
	jsonStr := `{"text":"` + "registration restarts!" + `"}`

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
